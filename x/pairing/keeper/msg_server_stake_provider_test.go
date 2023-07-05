package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testcommon "github.com/lavanet/lava/testutil/common"
	testkeeper "github.com/lavanet/lava/testutil/keeper"
	"github.com/lavanet/lava/utils/sigs"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
	"github.com/lavanet/lava/x/pairing/client/cli"
	"github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
	"github.com/stretchr/testify/require"
)

// Test that the optional moniker argument in StakeProvider doesn't break anything
func TestStakeProviderWithMoniker(t *testing.T) {
	// Create teststruct ts
	ts := &testStruct{
		providers: make([]*testcommon.Account, 0),
		clients:   make([]*testcommon.Account, 0),
	}
	ts.servers, ts.keepers, ts.ctx = testkeeper.InitAllKeepers(t)
	ts.keepers.Epochstorage.SetEpochDetails(sdk.UnwrapSDKContext(ts.ctx), *epochstoragetypes.DefaultGenesis().EpochDetails)
	// Create a mock spec
	ts.spec = testcommon.CreateMockSpec()
	ts.keepers.Spec.SetSpec(sdk.UnwrapSDKContext(ts.ctx), ts.spec)

	// define tests (valid indicates whether the test should succeed)
	tests := []struct {
		name         string
		moniker      string
		validStake   bool
		validMoniker bool
	}{
		{"NormalMoniker", "exampleMoniker", true, true},
		{"WeirdCharsMoniker", "ビッグファームへようこそ", true, true},
		{"OversizedMoniker", "aReallyReallyReallyReallyReallyReallyReallyLongMoniker", true, false}, // validMoniker = false because moniker should be < 50 characters -> the original moniker won't be equal to the assigned one
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Advance epoch
			ts.ctx = testkeeper.AdvanceEpoch(ts.ctx, ts.keepers)

			// Stake provider with moniker
			sk, address := sigs.GenerateFloatingKey()
			ts.providers = append(ts.providers, &testcommon.Account{SK: sk, Addr: address})
			err := ts.keepers.BankKeeper.SetBalance(sdk.UnwrapSDKContext(ts.ctx), address, sdk.NewCoins(sdk.NewCoin(epochstoragetypes.TokenDenom, sdk.NewInt(balance))))
			require.Nil(t, err)
			endpoints := []epochstoragetypes.Endpoint{}
			endpoints = append(endpoints, epochstoragetypes.Endpoint{IPPORT: "123", ApiInterfaces: []string{ts.spec.ApiCollections[0].CollectionData.ApiInterface}, Geolocation: 1})
			_, err = ts.servers.PairingServer.StakeProvider(ts.ctx, &types.MsgStakeProvider{Creator: address.String(), ChainID: ts.spec.Name, Amount: sdk.NewCoin(epochstoragetypes.TokenDenom, sdk.NewInt(stake)), Geolocation: 1, Endpoints: endpoints, Moniker: tt.moniker})
			require.Nil(t, err)

			// Advance epoch to apply the stake
			ts.ctx = testkeeper.AdvanceEpoch(ts.ctx, ts.keepers)

			// Get the stake entry and check the provider is staked
			stakeEntry, foundProvider, _ := ts.keepers.Epochstorage.GetStakeEntryByAddressCurrent(sdk.UnwrapSDKContext(ts.ctx), ts.spec.GetIndex(), address)
			require.Equal(t, tt.validStake, foundProvider)

			// Check the assigned moniker
			if tt.validMoniker {
				require.Equal(t, tt.moniker, stakeEntry.Moniker)
			} else {
				require.NotEqual(t, tt.moniker, stakeEntry.Moniker)
			}
		})
	}
}

func TestModifyStakeProviderWithMoniker(t *testing.T) {
	// Create teststruct ts
	ts := &testStruct{
		providers: make([]*testcommon.Account, 0),
		clients:   make([]*testcommon.Account, 0),
	}
	ts.servers, ts.keepers, ts.ctx = testkeeper.InitAllKeepers(t)
	ts.keepers.Epochstorage.SetEpochDetails(sdk.UnwrapSDKContext(ts.ctx), *epochstoragetypes.DefaultGenesis().EpochDetails)
	// Create a mock spec
	ts.spec = testcommon.CreateMockSpec()
	ts.keepers.Spec.SetSpec(sdk.UnwrapSDKContext(ts.ctx), ts.spec)

	// Advance epoch
	ts.ctx = testkeeper.AdvanceEpoch(ts.ctx, ts.keepers)

	moniker := "exampleMoniker"

	// Stake provider with moniker
	sk, address := sigs.GenerateFloatingKey()
	ts.providers = append(ts.providers, &testcommon.Account{SK: sk, Addr: address})
	err := ts.keepers.BankKeeper.SetBalance(sdk.UnwrapSDKContext(ts.ctx), address, sdk.NewCoins(sdk.NewCoin(epochstoragetypes.TokenDenom, sdk.NewInt(balance))))
	require.Nil(t, err)
	endpoints := []epochstoragetypes.Endpoint{}
	endpoints = append(endpoints, epochstoragetypes.Endpoint{IPPORT: "123", ApiInterfaces: []string{ts.spec.ApiCollections[0].CollectionData.ApiInterface}, Geolocation: 1})
	_, err = ts.servers.PairingServer.StakeProvider(ts.ctx, &types.MsgStakeProvider{Creator: address.String(), ChainID: ts.spec.Name, Amount: sdk.NewCoin(epochstoragetypes.TokenDenom, sdk.NewInt(stake/2)), Geolocation: 1, Endpoints: endpoints, Moniker: moniker})
	require.Nil(t, err)

	// Advance epoch to apply the stake
	ts.ctx = testkeeper.AdvanceEpoch(ts.ctx, ts.keepers)

	// Get the stake entry and check the provider is staked
	stakeEntry, foundProvider, _ := ts.keepers.Epochstorage.GetStakeEntryByAddressCurrent(sdk.UnwrapSDKContext(ts.ctx), ts.spec.GetIndex(), address)
	require.True(t, foundProvider)
	require.Equal(t, moniker, stakeEntry.Moniker)

	// modify moniker
	moniker = "anotherExampleMoniker"
	_, err = ts.servers.PairingServer.StakeProvider(ts.ctx, &types.MsgStakeProvider{Creator: address.String(), ChainID: ts.spec.Name, Amount: sdk.NewCoin(epochstoragetypes.TokenDenom, sdk.NewInt(stake)), Geolocation: 1, Endpoints: endpoints, Moniker: moniker})
	require.Nil(t, err)

	// Advance epoch to apply the stake
	ts.ctx = testkeeper.AdvanceEpoch(ts.ctx, ts.keepers)

	// Get the stake entry and check the provider is staked
	stakeEntry, foundProvider, _ = ts.keepers.Epochstorage.GetStakeEntryByAddressCurrent(sdk.UnwrapSDKContext(ts.ctx), ts.spec.GetIndex(), address)
	require.True(t, foundProvider)

	require.Equal(t, moniker, stakeEntry.Moniker)
}

func TestCmdStakeProviderGeoConfigAndEnum(t *testing.T) {
	buildEndpoint := func(geoloc string) []string {
		hostip := "127.0.0.1:3351"
		apiInterface := "jsonrpc"
		return []string{hostip + "," + geoloc + "," + apiInterface}
	}
	testCases := []struct {
		name        string
		endpoints   []string
		geolocation string
		valid       bool
	}{
		// single uint geolocation config tests
		{
			name:        "Single uint geolocation - happy flow",
			endpoints:   buildEndpoint("1"),
			geolocation: "1",
			valid:       true,
		},
		{
			name:        "Single uint geolocation - endpoint geo not equal to geo",
			endpoints:   buildEndpoint("2"),
			geolocation: "1",
			valid:       false,
		},
		{
			name:        "Single uint geolocation - endpoint geo not equal to geo (geo includes endpoint geo)",
			endpoints:   buildEndpoint("1"),
			geolocation: "3",
			valid:       false,
		},
		{
			name:        "Single uint geolocation - endpoint has geo of multiple regions",
			endpoints:   buildEndpoint("3"),
			geolocation: "3",
			valid:       false,
		},
		{
			name:        "Single uint geolocation - bad endpoint geo",
			endpoints:   buildEndpoint("20555"),
			geolocation: "1",
			valid:       false,
		},

		// single string geolocation config tests
		{
			name:        "Single string geolocation - happy flow",
			endpoints:   buildEndpoint("EU"),
			geolocation: "EU",
			valid:       true,
		},
		{
			name:        "Single string geolocation - endpoint geo not equal to geo",
			endpoints:   buildEndpoint("AS"),
			geolocation: "EU",
			valid:       false,
		},
		{
			name:        "Single string geolocation - endpoint geo not equal to geo (geo includes endpoint geo)",
			endpoints:   buildEndpoint("EU"),
			geolocation: "EU,USC",
			valid:       false,
		},
		{
			name:        "Single string geolocation - bad geo",
			endpoints:   buildEndpoint("EU"),
			geolocation: "BLABLA",
			valid:       false,
		},
		{
			name:        "Single string geolocation - bad geo",
			endpoints:   buildEndpoint("BLABLA"),
			geolocation: "EU",
			valid:       false,
		},

		// multiple uint geolocation config tests
		{
			name:        "Multiple uint geolocations - happy flow",
			endpoints:   append(buildEndpoint("1"), buildEndpoint("2")...),
			geolocation: "3",
			valid:       true,
		},
		{
			name:        "Multiple uint geolocations - endpoint geo not equal to geo",
			endpoints:   append(buildEndpoint("1"), buildEndpoint("4")...),
			geolocation: "2",
			valid:       false,
		},
		{
			name:        "Multiple uint geolocations - one endpoint has multi-region geo",
			endpoints:   append(buildEndpoint("1"), buildEndpoint("3")...),
			geolocation: "2",
			valid:       false,
		},

		// multiple string geolocation config tests
		{
			name:        "Multiple string geolocations - happy flow",
			endpoints:   append(buildEndpoint("AS"), buildEndpoint("EU")...),
			geolocation: "EU,AS",
			valid:       true,
		},
		{
			name:        "Multiple string geolocations - endpoint geo not equal to geo",
			endpoints:   append(buildEndpoint("EU"), buildEndpoint("USC")...),
			geolocation: "EU,AS",
			valid:       false,
		},

		// global config tests
		{
			name:        "Global uint geolocation - happy flow",
			endpoints:   buildEndpoint("65535"),
			geolocation: "65535",
			valid:       true,
		},
		{
			name:        "Global uint geolocation - happy flow 2 - global in one endpoint",
			endpoints:   append(buildEndpoint("2"), buildEndpoint("65535")...),
			geolocation: "65535",
			valid:       true,
		},
		{
			name:        "Global uint geolocation - endpoint geo not match geo",
			endpoints:   append(buildEndpoint("2"), buildEndpoint("65535")...),
			geolocation: "7",
			valid:       false,
		},
		{
			name:        "Global string geolocation - happy flow",
			endpoints:   buildEndpoint("GL"),
			geolocation: "GL",
			valid:       true,
		},
		{
			name:        "Global string geolocation - happy flow 2 - global in one endpoint",
			endpoints:   append(buildEndpoint("EU"), buildEndpoint("GL")...),
			geolocation: "GL",
			valid:       true,
		},
		{
			name:        "Global string geolocation - endpoint geo not match geo",
			endpoints:   append(buildEndpoint("EU"), buildEndpoint("GL")...),
			geolocation: "EU,AS,USC",
			valid:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _, err := cli.HandleEndpointsAndGeolocationArgs(tc.endpoints, tc.geolocation)
			if tc.valid {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
			}
		})
	}
}

func TestStakeEndpoints(t *testing.T) {
	ts := &testStruct{
		providers: make([]*testcommon.Account, 0),
		clients:   make([]*testcommon.Account, 0),
	}
	ts.servers, ts.keepers, ts.ctx = testkeeper.InitAllKeepers(t)
	ts.keepers.Epochstorage.SetEpochDetails(sdk.UnwrapSDKContext(ts.ctx), *epochstoragetypes.DefaultGenesis().EpochDetails)
	// Create a mock spec
	ts.spec = testcommon.CreateMockSpec() // basic stuff

	apiCollections := []*spectypes.ApiCollection{
		{
			CollectionData: spectypes.CollectionData{
				ApiInterface: "mandatory",
				InternalPath: "",
				Type:         "",
				AddOn:        "",
			},
			Enabled: true,
		},
		{
			CollectionData: spectypes.CollectionData{
				ApiInterface: "mandatory2",
				InternalPath: "",
				Type:         "",
				AddOn:        "",
			},
			Enabled: true,
		},
		{
			CollectionData: spectypes.CollectionData{
				ApiInterface: "mandatory2",
				InternalPath: "",
				Type:         "banana",
				AddOn:        "",
			},
			Enabled: true,
		},
		{
			CollectionData: spectypes.CollectionData{
				ApiInterface: "mandatory",
				InternalPath: "",
				Type:         "",
				AddOn:        "addon",
			},
			Enabled: true,
		},
		{
			CollectionData: spectypes.CollectionData{
				ApiInterface: "mandatory2",
				InternalPath: "",
				Type:         "",
				AddOn:        "addon",
			},
			Enabled: true,
		},
		{
			CollectionData: spectypes.CollectionData{
				ApiInterface: "mandatory",
				InternalPath: "",
				Type:         "",
				AddOn:        "unique-addon",
			},
			Enabled: true,
		},
		{
			CollectionData: spectypes.CollectionData{
				ApiInterface: "optional",
				InternalPath: "",
				Type:         "",
				AddOn:        "optional",
			},
			Enabled: true,
		},
	}

	ts.spec.ApiCollections = apiCollections
	ts.keepers.Spec.SetSpec(sdk.UnwrapSDKContext(ts.ctx), ts.spec)
	// Advance epoch
	ts.ctx = testkeeper.AdvanceEpoch(ts.ctx, ts.keepers)
	provider := testcommon.CreateNewAccount(ts.ctx, *ts.keepers, balance)
	ts.providers = append(ts.providers, &provider)

	getEndpoint := func(host string, apiInterfaces []string, addons []string, geoloc uint64) epochstoragetypes.Endpoint {
		return epochstoragetypes.Endpoint{
			IPPORT:        host,
			Geolocation:   geoloc,
			Addons:        addons,
			ApiInterfaces: apiInterfaces,
		}
	}
	type testEndpoint struct {
		name        string
		endpoints   []epochstoragetypes.Endpoint
		success     bool
		geolocation uint64
	}
	playbook := []testEndpoint{
		{
			name:        "empty single",
			endpoints:   append([]epochstoragetypes.Endpoint{}, getEndpoint("123", []string{}, []string{}, 1)),
			success:     true,
			geolocation: 1,
		},
		{
			name:        "partial apiInterface implementation",
			endpoints:   append([]epochstoragetypes.Endpoint{}, getEndpoint("123", []string{"mandatory"}, []string{}, 1)),
			success:     false,
			geolocation: 1,
		},
		{
			name:        "explicit",
			endpoints:   append([]epochstoragetypes.Endpoint{}, getEndpoint("123", []string{"mandatory", "mandatory2"}, []string{}, 1)),
			success:     true,
			geolocation: 1,
		},
		{
			name:        "divided explicit",
			endpoints:   append([]epochstoragetypes.Endpoint{getEndpoint("123", []string{"mandatory"}, []string{}, 1)}, getEndpoint("123", []string{"mandatory2"}, []string{}, 1)),
			success:     true,
			geolocation: 1,
		},
		{
			name:        "partial in each geolocation",
			endpoints:   append([]epochstoragetypes.Endpoint{getEndpoint("123", []string{"mandatory"}, []string{}, 1)}, getEndpoint("123", []string{"mandatory2"}, []string{}, 2)),
			success:     false,
			geolocation: 3,
		},
		{
			name: "empty multi geo",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{}, []string{}, 1),
				getEndpoint("123", []string{}, []string{}, 2)},
			success:     true,
			geolocation: 3,
		},
		{
			name: "explicit divided multi geo",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{"mandatory"}, []string{}, 1),
				getEndpoint("123", []string{"mandatory2"}, []string{}, 1),
				getEndpoint("123", []string{"mandatory"}, []string{}, 2),
				getEndpoint("123", []string{"mandatory2"}, []string{}, 2)},
			success:     true,
			geolocation: 3,
		},
		{
			name: "explicit divided multi geo in addons split",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{}, []string{"mandatory"}, 1),
				getEndpoint("123", []string{}, []string{"mandatory2"}, 1),
				getEndpoint("123", []string{}, []string{"mandatory"}, 2),
				getEndpoint("123", []string{}, []string{"mandatory2"}, 2)},
			success:     true,
			geolocation: 3,
		},
		{
			name: "explicit divided multi geo in addons together",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{}, []string{"mandatory", "mandatory2"}, 1),
				getEndpoint("123", []string{}, []string{"mandatory", "mandatory2"}, 2)},
			success:     true,
			geolocation: 3,
		},
		{
			name: "empty with addon partial-geo",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{}, []string{"addon"}, 1),
				getEndpoint("123", []string{}, []string{""}, 2)},
			success:     true,
			geolocation: 3,
		},
		{
			name: "empty with addon multi-geo",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{}, []string{"addon"}, 1),
				getEndpoint("123", []string{}, []string{"addon"}, 2)},
			success:     true,
			geolocation: 3,
		},
		{
			name: "empty with unique addon",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{}, []string{"addon", "unique-addon"}, 1),
				getEndpoint("123", []string{}, []string{"addon"}, 2)},
			success:     false,
			geolocation: 3,
		},
		{
			name: "explicit with unique addon partial geo",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{"mandatory"}, []string{"unique-addon"}, 1),
				getEndpoint("123", []string{"mandatory2"}, []string{}, 1),
				getEndpoint("123", []string{}, []string{"addon"}, 2)},
			success:     true,
			geolocation: 3,
		},
		{
			name: "explicit with addon + unique addon partial geo",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{"mandatory"}, []string{"addon", "unique-addon"}, 1),
				getEndpoint("123", []string{"mandatory2"}, []string{"addon"}, 1),
				getEndpoint("123", []string{}, []string{"addon"}, 2)},
			success:     true,
			geolocation: 3,
		},
		{
			name: "partial explicit and full emptry with addon + unique addon",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{"mandatory"}, []string{"addon", "unique-addon"}, 1),
				getEndpoint("123", []string{}, []string{"addon"}, 1)},
			success:     true,
			geolocation: 1,
		},
		{
			name: "empty + explicit optional",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{}, []string{}, 1),
				getEndpoint("123", []string{"optional"}, []string{}, 1)},
			success:     true,
			geolocation: 1,
		},
		{
			name: "empty + explicit optional in addon",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{}, []string{}, 1),
				getEndpoint("123", []string{}, []string{"optional"}, 1)},
			success:     true,
			geolocation: 1,
		},
		{
			name: "empty + explicit optional + optional addon",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{}, []string{}, 1),
				getEndpoint("123", []string{"optional"}, []string{"optional"}, 1)},
			success:     true,
			geolocation: 1,
		},
		{
			name: "explicit optional",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{"optional"}, []string{"optional"}, 1)},
			success:     false,
			geolocation: 1,
		},
		{
			name: "full partial geo",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{}, []string{"addon"}, 1),
				getEndpoint("123", []string{"mandatory"}, []string{"unique-addon"}, 1),
				getEndpoint("123", []string{"optional"}, []string{}, 1),
				getEndpoint("123", []string{}, []string{}, 2)},
			success:     true,
			geolocation: 3,
		},
		{
			name: "full multi geo",
			endpoints: []epochstoragetypes.Endpoint{
				getEndpoint("123", []string{}, []string{"addon"}, 1),
				getEndpoint("123", []string{"mandatory"}, []string{"unique-addon"}, 1),
				getEndpoint("123", []string{"optional"}, []string{}, 1),
				getEndpoint("123", []string{}, []string{"addon"}, 2),
				getEndpoint("123", []string{"mandatory"}, []string{"unique-addon"}, 2),
				getEndpoint("123", []string{"optional"}, []string{}, 2)},
			success:     true,
			geolocation: 3,
		},
	}

	for _, play := range playbook {
		t.Run(play.name, func(t *testing.T) {
			_, err := ts.servers.PairingServer.StakeProvider(ts.ctx, &types.MsgStakeProvider{Creator: ts.providers[0].Addr.String(), ChainID: ts.spec.Index, Amount: sdk.NewCoin(epochstoragetypes.TokenDenom, sdk.NewInt(stake)), Geolocation: play.geolocation, Endpoints: play.endpoints})
			if play.success {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
