// source: cosmos/bank/v1beta1/authz.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SendAuthorization allows the grantee to spend up to spend_limit coins from
// the granter's account.
//
// Since: cosmos-sdk 0.43
type SendAuthorization struct {
	SpendLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=spend_limit,json=spendLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"spend_limit"`
	// allow_list specifies an optional list of addresses to whom the grantee can send tokens on behalf of the
	// granter. If omitted, any recipient is allowed.
	//
	// Since: cosmos-sdk 0.47
	AllowList []string `protobuf:"bytes,2,rep,name=allow_list,json=allowList,proto3" json:"allow_list,omitempty"`
}

func (m *SendAuthorization) Reset()         { *m = SendAuthorization{} }
func (m *SendAuthorization) String() string { return proto.CompactTextString(m) }
func (*SendAuthorization) ProtoMessage()    {}
func (*SendAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d2a37888ea779f, []int{0}
}

func (m *SendAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *SendAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *SendAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendAuthorization.Merge(m, src)
}

func (m *SendAuthorization) XXX_Size() int {
	return m.Size()
}

func (m *SendAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_SendAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_SendAuthorization proto.InternalMessageInfo

func (m *SendAuthorization) GetSpendLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.SpendLimit
	}
	return nil
}

func (m *SendAuthorization) GetAllowList() []string {
	if m != nil {
		return m.AllowList
	}
	return nil
}

func init() {

}

func init() {} // proto.RegisterFile("cosmos/bank/v1beta1/authz.proto", fileDescriptor_a4d2a37888ea779f) }

var fileDescriptor_a4d2a37888ea779f = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x4a, 0xcc, 0xcb, 0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4,
	0x4f, 0x2c, 0x2d, 0xc9, 0xa8, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86, 0x28, 0xd0,
	0x03, 0x29, 0xd0, 0x83, 0x2a, 0x90, 0x12, 0x4c, 0xcc, 0xcd, 0xcc, 0xcb, 0xd7, 0x07, 0x93, 0x10,
	0x75, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xa6, 0x3e, 0x88, 0x05, 0x15, 0x95, 0x84, 0xe8,
	0x8e, 0x87, 0x48, 0x40, 0x8d, 0x82, 0x48, 0xc9, 0xc1, 0x6d, 0x2e, 0x4e, 0x85, 0xdb, 0x9c, 0x9c,
	0x9f, 0x99, 0x07, 0x91, 0x57, 0xba, 0xc9, 0xc8, 0x25, 0x18, 0x9c, 0x9a, 0x97, 0xe2, 0x58, 0x5a,
	0x92, 0x91, 0x5f, 0x94, 0x59, 0x95, 0x58, 0x92, 0x99, 0x9f, 0x27, 0x54, 0xc8, 0xc5, 0x5d, 0x5c,
	0x90, 0x9a, 0x97, 0x12, 0x9f, 0x93, 0x99, 0x9b, 0x59, 0x22, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d,
	0x24, 0xa9, 0x07, 0x77, 0x64, 0x71, 0x2a, 0xcc, 0x91, 0x7a, 0xce, 0xf9, 0x99, 0x79, 0x4e, 0xa6,
	0x27, 0xee, 0xc9, 0x33, 0xac, 0xba, 0x2f, 0xaf, 0x91, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97,
	0x9c, 0x9f, 0x0b, 0x75, 0x06, 0x94, 0xd2, 0x2d, 0x4e, 0xc9, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d,
	0x06, 0x6b, 0x28, 0x5e, 0xf1, 0x7c, 0x83, 0x16, 0x63, 0x10, 0x17, 0xd8, 0x12, 0x1f, 0x90, 0x1d,
	0x42, 0xb2, 0x5c, 0x5c, 0x89, 0x39, 0x39, 0xf9, 0xe5, 0xf1, 0x39, 0x99, 0xc5, 0x25, 0x12, 0x4c,
	0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0x9c, 0x60, 0x11, 0x9f, 0xcc, 0xe2, 0x12, 0x2b, 0xa3, 0x53, 0x5b,
	0x74, 0x79, 0x51, 0x1c, 0xd9, 0xf5, 0x7c, 0x83, 0x96, 0x0c, 0x92, 0xe9, 0x18, 0xbe, 0x70, 0x72,
	0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96,
	0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x4d, 0xbc, 0xee, 0xac, 0x80,
	0xc4, 0x13, 0xd8, 0xb9, 0x49, 0x6c, 0xe0, 0x70, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2a,
	0xa1, 0xb7, 0x33, 0xc3, 0x01, 0x00, 0x00,
}

func (m *SendAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AllowList) > 0 {
		for iNdEx := len(m.AllowList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowList[iNdEx])
			copy(dAtA[i:], m.AllowList[iNdEx])
			i = encodeVarintAuthz(dAtA, i, uint64(len(m.AllowList[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.SpendLimit) > 0 {
		for iNdEx := len(m.SpendLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpendLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuthz(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthz(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthz(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *SendAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SpendLimit) > 0 {
		for _, e := range m.SpendLimit {
			l = e.Size()
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	if len(m.AllowList) > 0 {
		for _, s := range m.AllowList {
			l = len(s)
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	return n
}

func sovAuthz(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozAuthz(x uint64) (n int) {
	return sovAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *SendAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: SendAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpendLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpendLimit = append(m.SpendLimit, types.Coin{})
			if err := m.SpendLimit[len(m.SpendLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowList = append(m.AllowList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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

func skipAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
				return 0, ErrInvalidLengthAuthz
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthz
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthz
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthz        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthz          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthz = fmt.Errorf("proto: unexpected end of group")
)
