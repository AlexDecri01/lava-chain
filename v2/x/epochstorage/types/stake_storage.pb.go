// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/epochstorage/stake_storage.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type StakeStorage struct {
	Index          string       `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	StakeEntries   []StakeEntry `protobuf:"bytes,2,rep,name=stakeEntries,proto3" json:"stakeEntries"`
	EpochBlockHash []byte       `protobuf:"bytes,3,opt,name=epochBlockHash,proto3" json:"epochBlockHash,omitempty"`
}

func (m *StakeStorage) Reset()         { *m = StakeStorage{} }
func (m *StakeStorage) String() string { return proto.CompactTextString(m) }
func (*StakeStorage) ProtoMessage()    {}
func (*StakeStorage) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7b78aecc265fd4, []int{0}
}
func (m *StakeStorage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakeStorage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakeStorage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakeStorage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeStorage.Merge(m, src)
}
func (m *StakeStorage) XXX_Size() int {
	return m.Size()
}
func (m *StakeStorage) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeStorage.DiscardUnknown(m)
}

var xxx_messageInfo_StakeStorage proto.InternalMessageInfo

func (m *StakeStorage) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StakeStorage) GetStakeEntries() []StakeEntry {
	if m != nil {
		return m.StakeEntries
	}
	return nil
}

func (m *StakeStorage) GetEpochBlockHash() []byte {
	if m != nil {
		return m.EpochBlockHash
	}
	return nil
}

func init() {
	proto.RegisterType((*StakeStorage)(nil), "lavanet.lava.epochstorage.StakeStorage")
}

func init() {
	proto.RegisterFile("lavanet/lava/epochstorage/stake_storage.proto", fileDescriptor_be7b78aecc265fd4)
}

var fileDescriptor_be7b78aecc265fd4 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcd, 0x49, 0x2c, 0x4b,
	0xcc, 0x4b, 0x2d, 0xd1, 0x07, 0xd1, 0xfa, 0xa9, 0x05, 0xf9, 0xc9, 0x19, 0xc5, 0x25, 0xf9, 0x45,
	0x89, 0xe9, 0xa9, 0xfa, 0xc5, 0x25, 0x89, 0xd9, 0xa9, 0xf1, 0x50, 0x9e, 0x5e, 0x41, 0x51, 0x7e,
	0x49, 0xbe, 0x90, 0x24, 0x54, 0xb9, 0x1e, 0x88, 0xd6, 0x43, 0x56, 0x2e, 0xa5, 0x4d, 0xc8, 0xa4,
	0xd4, 0xbc, 0x92, 0xa2, 0x4a, 0x88, 0x39, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xa6, 0x3e,
	0x88, 0x05, 0x11, 0x55, 0x9a, 0xcb, 0xc8, 0xc5, 0x13, 0x0c, 0x52, 0x1b, 0x0c, 0xd1, 0x28, 0x24,
	0xc2, 0xc5, 0x9a, 0x99, 0x97, 0x92, 0x5a, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1,
	0x08, 0xf9, 0x73, 0xf1, 0x80, 0x4d, 0x74, 0xcd, 0x2b, 0x29, 0xca, 0x4c, 0x2d, 0x96, 0x60, 0x52,
	0x60, 0xd6, 0xe0, 0x36, 0x52, 0xd5, 0xc3, 0xe9, 0x36, 0xbd, 0x60, 0x98, 0xf2, 0x4a, 0x27, 0x96,
	0x13, 0xf7, 0xe4, 0x19, 0x82, 0x50, 0x0c, 0x10, 0x52, 0xe3, 0xe2, 0x03, 0x2b, 0x77, 0xca, 0xc9,
	0x4f, 0xce, 0xf6, 0x48, 0x2c, 0xce, 0x90, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x09, 0x42, 0x13, 0x75,
	0xf2, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c,
	0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xfd, 0xf4, 0xcc, 0x92,
	0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x94, 0x70, 0x28, 0x33, 0xd2, 0xaf, 0x40, 0x0d,
	0x8c, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x8f, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x82, 0xbe, 0xa5, 0xcb, 0x80, 0x01, 0x00, 0x00,
}

func (m *StakeStorage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakeStorage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakeStorage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EpochBlockHash) > 0 {
		i -= len(m.EpochBlockHash)
		copy(dAtA[i:], m.EpochBlockHash)
		i = encodeVarintStakeStorage(dAtA, i, uint64(len(m.EpochBlockHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.StakeEntries) > 0 {
		for iNdEx := len(m.StakeEntries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StakeEntries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStakeStorage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintStakeStorage(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStakeStorage(dAtA []byte, offset int, v uint64) int {
	offset -= sovStakeStorage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StakeStorage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovStakeStorage(uint64(l))
	}
	if len(m.StakeEntries) > 0 {
		for _, e := range m.StakeEntries {
			l = e.Size()
			n += 1 + l + sovStakeStorage(uint64(l))
		}
	}
	l = len(m.EpochBlockHash)
	if l > 0 {
		n += 1 + l + sovStakeStorage(uint64(l))
	}
	return n
}

func sovStakeStorage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStakeStorage(x uint64) (n int) {
	return sovStakeStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StakeStorage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStakeStorage
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
			return fmt.Errorf("proto: StakeStorage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakeStorage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStakeStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakeStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeEntries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeStorage
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
				return ErrInvalidLengthStakeStorage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakeEntries = append(m.StakeEntries, StakeEntry{})
			if err := m.StakeEntries[len(m.StakeEntries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochBlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStakeStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStakeStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochBlockHash = append(m.EpochBlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.EpochBlockHash == nil {
				m.EpochBlockHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStakeStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStakeStorage
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
func skipStakeStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStakeStorage
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
					return 0, ErrIntOverflowStakeStorage
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
					return 0, ErrIntOverflowStakeStorage
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
				return 0, ErrInvalidLengthStakeStorage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStakeStorage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStakeStorage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStakeStorage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStakeStorage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStakeStorage = fmt.Errorf("proto: unexpected end of group")
)
