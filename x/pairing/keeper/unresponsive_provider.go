package keeper

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/utils"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
	"github.com/lavanet/lava/x/pairing/types"
)

const (
	THRESHOLD_FACTOR = 4
	SOFT_JAILS       = 2
	SOFT_JAIL_TIME   = 1 * time.Hour / time.Second
	HARD_JAIL_TIME   = 24 * time.Hour / time.Second
)

// PunishUnresponsiveProviders punished unresponsive providers (current punishment: freeze)
func (k Keeper) PunishUnresponsiveProviders(ctx sdk.Context, epochsNumToCheckCUForUnresponsiveProvider, epochsNumToCheckCUForComplainers uint64) {
	// check the epochsNum consts
	if epochsNumToCheckCUForComplainers <= 0 || epochsNumToCheckCUForUnresponsiveProvider <= 0 {
		utils.LavaFormatError("epoch to check CU for unresponsive provider or for complainer is zero",
			fmt.Errorf("invalid unresponsive provider consts"),
			utils.Attribute{Key: "epochsNumToCheckCUForUnresponsiveProvider", Value: epochsNumToCheckCUForUnresponsiveProvider},
			utils.Attribute{Key: "epochsNumToCheckCUForComplainers", Value: epochsNumToCheckCUForComplainers},
		)
	}

	// Get current epoch
	currentEpoch := k.epochStorageKeeper.GetEpochStart(ctx)

	// Get recommendedEpochNumToCollectPayment
	recommendedEpochNumToCollectPayment := k.RecommendedEpochNumToCollectPayment(ctx)

	// check which of the consts is larger
	largerEpochsNumConst := epochsNumToCheckCUForComplainers
	if epochsNumToCheckCUForUnresponsiveProvider > epochsNumToCheckCUForComplainers {
		largerEpochsNumConst = epochsNumToCheckCUForUnresponsiveProvider
	}

	// To check for punishment, we have to go back:
	//   recommendedEpochNumToCollectPayment+
	//   max(epochsNumToCheckCUForComplainers,epochsNumToCheckCUForUnresponsiveProvider)
	// epochs from the current epoch.
	minHistoryBlock, err := k.getBlockEpochsAgo(ctx, currentEpoch, largerEpochsNumConst+recommendedEpochNumToCollectPayment)
	if err != nil {
		// not enough history, do nothing
		return
	}

	// Get the current stake storages (from all chains).
	// Stake storages contain a list of stake entries (each for a different chain).
	providerStakeStorageList := k.getCurrentProviderStakeStorageList(ctx)
	if len(providerStakeStorageList) == 0 {
		// no provider is staked -> no one to punish
		return
	}

	// Go back recommendedEpochNumToCollectPayment
	minPaymentBlock, err := k.getBlockEpochsAgo(ctx, currentEpoch, recommendedEpochNumToCollectPayment)
	if err != nil {
		// not enough history, do nothiing
		return
	}

	// find the minimum number of providers in all the plans
	minProviders := uint64(math.MaxUint64)
	planIndices := k.planKeeper.GetAllPlanIndices(ctx)
	for _, index := range planIndices {
		plan, found := k.planKeeper.FindPlan(ctx, index, uint64(ctx.BlockHeight()))
		if found && plan.PlanPolicy.MaxProvidersToPair < minProviders {
			minProviders = plan.PlanPolicy.MaxProvidersToPair
		}
	}

	ProviderChainID := func(provider, chainID string) string {
		return provider + " " + chainID
	}

	// check all supported providers from all geolocations prior to making decisions
	existingProviders := map[string]uint64{}
	stakeEntries := map[string]epochstoragetypes.StakeEntry{}
	for _, providerStakeStorage := range providerStakeStorageList {
		providerStakeEntriesForChain := providerStakeStorage.GetStakeEntries()
		// count providers per geolocation
		for _, providerStakeEntry := range providerStakeEntriesForChain {
			if !providerStakeEntry.IsFrozen() {
				existingProviders[providerStakeEntry.GetChain()]++
			}
			stakeEntries[ProviderChainID(providerStakeEntry.Address, providerStakeEntry.Chain)] = providerStakeEntry
		}
	}

	// Go over the staked provider entries (on all chains) that has complaints
	// build a map that has all the relevant details: provider address, chain, epoch and ProviderEpochCu object
	keys := []string{}
	pecsDetailed := k.GetAllProviderEpochComplainerCuStore(ctx)
	complainedProviders := map[string]map[uint64]types.ProviderEpochComplainerCu{} // map[provider chainID]map[epoch]ProviderEpochComplainerCu
	for _, pec := range pecsDetailed {
		entry, ok := stakeEntries[ProviderChainID(pec.Provider, pec.ChainId)]
		if ok {
			if minHistoryBlock < entry.StakeAppliedBlock && entry.Jails == 0 {
				// this staked provider has too short history (either since staking
				// or since it was last unfrozen) - do not consider for jailing
				continue
			}
		} else {
			continue
		}

		key := ProviderChainID(pec.Provider, pec.ChainId)

		if _, ok := complainedProviders[key]; !ok {
			complainedProviders[key] = map[uint64]types.ProviderEpochComplainerCu{pec.Epoch: pec.ProviderEpochComplainerCu}
			keys = append(keys, key)
		} else {
			if _, ok := complainedProviders[key][pec.Epoch]; !ok {
				complainedProviders[key][pec.Epoch] = pec.ProviderEpochComplainerCu
			} else {
				utils.LavaFormatError("duplicate ProviderEpochCu key", fmt.Errorf("did not aggregate complainers CU"),
					utils.LogAttr("key", types.ProviderEpochCuKey(pec.Epoch, pec.Provider, pec.ChainId)),
				)
				continue
			}
		}
	}

	// go over all the providers, count the complainers CU and punish providers
	for _, key := range keys {
		components := strings.Split(key, " ")
		provider := components[0]
		chainID := components[1]
		// update the CU count for this provider in providerCuCounterForUnreponsivenessMap
		epochs, complaintCU, servicedCU, err := k.countCuForUnresponsiveness(ctx, provider, chainID, minPaymentBlock, epochsNumToCheckCUForUnresponsiveProvider, epochsNumToCheckCUForComplainers, complainedProviders[key])
		if err != nil {
			utils.LavaFormatError("unstake unresponsive providers failed to count CU", err,
				utils.Attribute{Key: "provider", Value: provider},
			)
			continue
		}

		// providerPaymentStorageKeyList is not empty -> provider should be punished
		if len(epochs) != 0 && existingProviders[chainID] > minProviders {
			entry, ok := stakeEntries[key]
			if !ok {
				utils.LavaFormatError("Jail_cant_get_stake_entry", types.FreezeStakeEntryNotFoundError, []utils.Attribute{{Key: "chainID", Value: chainID}, {Key: "providerAddress", Value: provider}}...)
				continue
			}
			err = k.punishUnresponsiveProvider(ctx, epochs, entry, complaintCU, servicedCU, complainedProviders[key])
			existingProviders[chainID]--
			if err != nil {
				utils.LavaFormatError("unstake unresponsive providers failed to punish provider", err,
					utils.Attribute{Key: "provider", Value: provider},
				)
			}
		}
	}
}

// getBlockEpochsAgo returns the block numEpochs back from the given blockHeight
func (k Keeper) getBlockEpochsAgo(ctx sdk.Context, blockHeight, numEpochs uint64) (uint64, error) {
	for counter := 0; counter < int(numEpochs); counter++ {
		var err error
		blockHeight, err = k.epochStorageKeeper.GetPreviousEpochStartForBlock(ctx, blockHeight)
		if err != nil {
			// too early in the chain life: bail without an error
			return uint64(0), err
		}
	}
	return blockHeight, nil
}

// Function to count the CU serviced by the unresponsive provider and the CU of the complainers. The function returns the keys of the objects containing complainer CU
func (k Keeper) countCuForUnresponsiveness(ctx sdk.Context, provider, chainId string, epoch, epochsNumToCheckCUForUnresponsiveProvider, epochsNumToCheckCUForComplainers uint64, providerEpochCuMap map[uint64]types.ProviderEpochComplainerCu) (epochs []uint64, complainersCu uint64, servicedCu uint64, errRet error) {
	// check which of the epoch consts is larger
	max := epochsNumToCheckCUForComplainers
	if epochsNumToCheckCUForUnresponsiveProvider > epochsNumToCheckCUForComplainers {
		max = epochsNumToCheckCUForUnresponsiveProvider
	}

	// count the CU serviced by the unersponsive provider and used CU of the complainers
	for counter := uint64(0); counter < max; counter++ {
		pec, ok := providerEpochCuMap[epoch]
		if ok {
			// counter is smaller than epochsNumToCheckCUForComplainers -> count complainer CU
			if counter < epochsNumToCheckCUForComplainers {
				complainersCu += pec.ComplainersCu
				epochs = append(epochs, epoch)
			}

			// counter is smaller than epochsNumToCheckCUForUnresponsiveProvider -> count CU serviced by the provider in the epoch
			if counter < epochsNumToCheckCUForUnresponsiveProvider {
				pec, found := k.GetProviderEpochCu(ctx, epoch, provider, chainId)
				if found {
					servicedCu += pec.ServicedCu
				}
			}
		}

		// Get previous epoch (from epochTemp)
		previousEpoch, err := k.epochStorageKeeper.GetPreviousEpochStartForBlock(ctx, epoch)
		if err != nil {
			return nil, 0, 0, utils.LavaFormatWarning("couldn't get previous epoch", err,
				utils.Attribute{Key: "epoch", Value: epoch},
			)
		}
		// update epoch
		epoch = previousEpoch
	}

	// the complainers' CU is larger than the provider serviced CU -> should be punished
	// epochs list is returned so the complainers' CU can be reset
	if complainersCu > THRESHOLD_FACTOR*servicedCu {
		return epochs, complainersCu, servicedCu, nil
	}

	return nil, complainersCu, servicedCu, nil
}

// Function that return the current stake storage for all chains
func (k Keeper) getCurrentProviderStakeStorageList(ctx sdk.Context) []epochstoragetypes.StakeStorage {
	var stakeStorageList []epochstoragetypes.StakeStorage

	// get all chain IDs
	chainIdList := k.specKeeper.GetAllChainIDs(ctx)

	// go over all chain IDs and keep their stake storage. If there is no stake storage for a specific chain, continue to the next one
	for _, chainID := range chainIdList {
		stakeStorage, found := k.epochStorageKeeper.GetStakeStorageCurrent(ctx, chainID)
		if !found {
			continue
		}
		stakeStorageList = append(stakeStorageList, stakeStorage)
	}

	return stakeStorageList
}

// Function that punishes providers. Current punishment is freeze
func (k Keeper) punishUnresponsiveProvider(ctx sdk.Context, epochs []uint64, stakeEntry epochstoragetypes.StakeEntry, complaintCU uint64, servicedCU uint64, providerEpochCuMap map[uint64]types.ProviderEpochComplainerCu) error {
	// if last jail was more than 24H ago, reset the jails counter
	if !stakeEntry.IsJailed(ctx.BlockTime().UTC().Unix() - int64(HARD_JAIL_TIME)) {
		stakeEntry.Jails = 0
	}
	stakeEntry.Jails++

	if stakeEntry.Jails > SOFT_JAILS {
		stakeEntry.Freeze()
		stakeEntry.JailEndTime = ctx.BlockTime().UTC().Unix() + int64(HARD_JAIL_TIME)
	} else {
		stakeEntry.JailEndTime = ctx.BlockTime().UTC().Unix() + int64(SOFT_JAIL_TIME)
		epochduration := k.downtimeKeeper.GetParams(ctx).EpochDuration / time.Second
		epochblocks := k.epochStorageKeeper.EpochBlocksRaw(ctx)
		stakeEntry.StakeAppliedBlock = uint64(ctx.BlockHeight()) + uint64(SOFT_JAIL_TIME/epochduration)*epochblocks
	}
	k.epochStorageKeeper.ModifyStakeEntryCurrent(ctx, stakeEntry.Chain, stakeEntry)

	utils.LogLavaEvent(ctx, k.Logger(ctx), types.ProviderJailedEventName,
		map[string]string{
			"provider_address": stakeEntry.Address,
			"chain_id":         stakeEntry.Chain,
			"complaint_cu":     strconv.FormatUint(complaintCU, 10),
			"serviced_cu":      strconv.FormatUint(servicedCU, 10),
		},
		"Unresponsive provider was freezed due to unresponsiveness")

	// reset the provider's complainer CU (so he won't get punished for the same complaints twice)
	k.resetComplainersCU(ctx, epochs, stakeEntry.Address, stakeEntry.Chain, providerEpochCuMap)

	return nil
}

// resetComplainersCU resets the complainers CU for a specific provider and chain
func (k Keeper) resetComplainersCU(ctx sdk.Context, epochs []uint64, provider string, chainID string, providerEpochCuMap map[uint64]types.ProviderEpochComplainerCu) {
	for _, epoch := range epochs {
		k.RemoveProviderEpochComplainerCu(ctx, epoch, provider, chainID)
	}
}
