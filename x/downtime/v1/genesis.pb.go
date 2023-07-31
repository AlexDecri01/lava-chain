// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: downtime/v1/genesis.proto

package v1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState is the genesis state of the downtime module.
type GenesisState struct {
	Params                     Params                       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Downtimes                  []*Downtime                  `protobuf:"bytes,2,rep,name=downtimes,proto3" json:"downtimes,omitempty"`
	DowntimesGarbageCollection []*DowntimeGarbageCollection `protobuf:"bytes,3,rep,name=downtimes_garbage_collection,json=downtimesGarbageCollection,proto3" json:"downtimes_garbage_collection,omitempty"`
	// last_block_time keeps track of when the last block time was set.
	// it's nullable because we might want it to be non existent.
	// we want it to exist when we have a genesis export-import migration scenario.
	LastBlockTime *time.Time `protobuf:"bytes,4,opt,name=last_block_time,json=lastBlockTime,proto3,stdtime" json:"last_block_time,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2acc27bef320760, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetDowntimes() []*Downtime {
	if m != nil {
		return m.Downtimes
	}
	return nil
}

func (m *GenesisState) GetDowntimesGarbageCollection() []*DowntimeGarbageCollection {
	if m != nil {
		return m.DowntimesGarbageCollection
	}
	return nil
}

func (m *GenesisState) GetLastBlockTime() *time.Time {
	if m != nil {
		return m.LastBlockTime
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lavanet.lava.downtime.v1.GenesisState")
}

func init() { proto.RegisterFile("downtime/v1/genesis.proto", fileDescriptor_a2acc27bef320760) }

var fileDescriptor_a2acc27bef320760 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x33, 0x6d, 0x29, 0xdf, 0x97, 0x2a, 0x42, 0x70, 0x11, 0x83, 0x4c, 0x4a, 0x41, 0xe8,
	0x6a, 0x86, 0xb6, 0x7b, 0x95, 0x28, 0x74, 0xe3, 0x42, 0xa2, 0x2b, 0x37, 0x61, 0x12, 0xc7, 0x31,
	0x38, 0xc9, 0x94, 0xce, 0xb4, 0xfa, 0x18, 0x7d, 0xac, 0x2e, 0xbb, 0x11, 0x5c, 0x55, 0x69, 0x5f,
	0x44, 0x66, 0xf2, 0xc7, 0x82, 0x14, 0x57, 0xf7, 0xde, 0x9c, 0xf3, 0xbb, 0x37, 0x39, 0xb1, 0x4f,
	0x1e, 0xc5, 0x6b, 0xae, 0xd2, 0x8c, 0xe2, 0xf9, 0x00, 0x33, 0x9a, 0x53, 0x99, 0x4a, 0x34, 0x99,
	0x0a, 0x25, 0x1c, 0x97, 0x93, 0x39, 0xc9, 0xa9, 0x42, 0xba, 0xa2, 0xca, 0x87, 0xe6, 0x03, 0xcf,
	0xdb, 0x85, 0x6a, 0xc1, 0x50, 0x9e, 0xcf, 0x84, 0x60, 0x9c, 0x62, 0x33, 0xc5, 0xb3, 0x27, 0xac,
	0x35, 0xa9, 0x48, 0x36, 0x29, 0x0d, 0xc7, 0x4c, 0x30, 0x61, 0x5a, 0xac, 0xbb, 0xe2, 0x69, 0xef,
	0xbd, 0x61, 0x1f, 0x8c, 0x8b, 0xf3, 0x77, 0x8a, 0x28, 0xea, 0x9c, 0xdb, 0xed, 0x09, 0x99, 0x92,
	0x4c, 0xba, 0xa0, 0x0b, 0xfa, 0x9d, 0x61, 0x17, 0xed, 0x7b, 0x1d, 0x74, 0x6b, 0x7c, 0x41, 0x6b,
	0xb9, 0xf6, 0xad, 0xb0, 0xa4, 0x9c, 0x4b, 0xfb, 0x7f, 0xe5, 0x91, 0x6e, 0xa3, 0xdb, 0xec, 0x77,
	0x86, 0xbd, 0xfd, 0x2b, 0xae, 0xcb, 0x3e, 0xfc, 0x81, 0x9c, 0x99, 0x7d, 0x5a, 0x0f, 0x11, 0x23,
	0xd3, 0x98, 0x30, 0x1a, 0x25, 0x82, 0x73, 0x9a, 0xa8, 0x54, 0xe4, 0x6e, 0xd3, 0x2c, 0x1d, 0xfd,
	0xbd, 0x74, 0x5c, 0xb0, 0x57, 0x35, 0x1a, 0xd6, 0x01, 0xca, 0x5f, 0x9a, 0x73, 0x63, 0x1f, 0x71,
	0x22, 0x55, 0x14, 0x73, 0x91, 0xbc, 0x44, 0xda, 0xe4, 0xb6, 0x4c, 0x02, 0x1e, 0x2a, 0xa2, 0x45,
	0x55, 0xb4, 0xe8, 0xbe, 0x8a, 0x36, 0xf8, 0xb7, 0x5c, 0xfb, 0x60, 0xf1, 0xe9, 0x83, 0xf0, 0x50,
	0xc3, 0x81, 0x66, 0xb5, 0x1a, 0x5c, 0x2c, 0x37, 0x10, 0xac, 0x36, 0x10, 0x7c, 0x6d, 0x20, 0x58,
	0x6c, 0xa1, 0xb5, 0xda, 0x42, 0xeb, 0x63, 0x0b, 0xad, 0x87, 0x33, 0x96, 0xaa, 0xe7, 0x59, 0x8c,
	0x12, 0x91, 0xe1, 0xf2, 0x13, 0x4c, 0xc5, 0x6f, 0x78, 0xe7, 0xf7, 0xc6, 0x6d, 0x73, 0x6d, 0xf4,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x54, 0x62, 0x34, 0x29, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastBlockTime != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.LastBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.LastBlockTime):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintGenesis(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DowntimesGarbageCollection) > 0 {
		for iNdEx := len(m.DowntimesGarbageCollection) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DowntimesGarbageCollection[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Downtimes) > 0 {
		for iNdEx := len(m.Downtimes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Downtimes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Downtimes) > 0 {
		for _, e := range m.Downtimes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DowntimesGarbageCollection) > 0 {
		for _, e := range m.DowntimesGarbageCollection {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LastBlockTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.LastBlockTime)
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Downtimes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Downtimes = append(m.Downtimes, &Downtime{})
			if err := m.Downtimes[len(m.Downtimes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DowntimesGarbageCollection", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DowntimesGarbageCollection = append(m.DowntimesGarbageCollection, &DowntimeGarbageCollection{})
			if err := m.DowntimesGarbageCollection[len(m.DowntimesGarbageCollection)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastBlockTime == nil {
				m.LastBlockTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.LastBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
