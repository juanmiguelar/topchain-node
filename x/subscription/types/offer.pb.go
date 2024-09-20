// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: topchain/subscription/offer.proto

package types

import (
	fmt "fmt"
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

type Offer_Status int32

const (
	Offer_UNSPECIFIED Offer_Status = 0
)

var Offer_Status_name = map[int32]string{
	0: "UNSPECIFIED",
}

var Offer_Status_value = map[string]int32{
	"UNSPECIFIED": 0,
}

func (x Offer_Status) String() string {
	return proto.EnumName(Offer_Status_name, int32(x))
}

func (Offer_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ff67a1f2f1b26c98, []int{0, 0}
}

type Offer struct {
	Id              string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Requester       string       `protobuf:"bytes,2,opt,name=requester,proto3" json:"requester,omitempty"`
	CroId           string       `protobuf:"bytes,3,opt,name=cro_id,json=croId,proto3" json:"cro_id,omitempty"`
	SubscriptionIds []string     `protobuf:"bytes,4,rep,name=subscription_ids,json=subscriptionIds,proto3" json:"subscription_ids,omitempty"`
	Status          Offer_Status `protobuf:"varint,5,opt,name=status,proto3,enum=topchain.subscription.Offer_Status" json:"status,omitempty"`
	InitialAmount   uint64       `protobuf:"varint,6,opt,name=initial_amount,json=initialAmount,proto3" json:"initial_amount,omitempty"`
	AvailableAmount uint64       `protobuf:"varint,7,opt,name=available_amount,json=availableAmount,proto3" json:"available_amount,omitempty"`
	StartBlock      uint64       `protobuf:"varint,8,opt,name=start_block,json=startBlock,proto3" json:"start_block,omitempty"`
	EndBlock        uint64       `protobuf:"varint,9,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
}

func (m *Offer) Reset()         { *m = Offer{} }
func (m *Offer) String() string { return proto.CompactTextString(m) }
func (*Offer) ProtoMessage()    {}
func (*Offer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff67a1f2f1b26c98, []int{0}
}
func (m *Offer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Offer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Offer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Offer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Offer.Merge(m, src)
}
func (m *Offer) XXX_Size() int {
	return m.Size()
}
func (m *Offer) XXX_DiscardUnknown() {
	xxx_messageInfo_Offer.DiscardUnknown(m)
}

var xxx_messageInfo_Offer proto.InternalMessageInfo

func (m *Offer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Offer) GetRequester() string {
	if m != nil {
		return m.Requester
	}
	return ""
}

func (m *Offer) GetCroId() string {
	if m != nil {
		return m.CroId
	}
	return ""
}

func (m *Offer) GetSubscriptionIds() []string {
	if m != nil {
		return m.SubscriptionIds
	}
	return nil
}

func (m *Offer) GetStatus() Offer_Status {
	if m != nil {
		return m.Status
	}
	return Offer_UNSPECIFIED
}

func (m *Offer) GetInitialAmount() uint64 {
	if m != nil {
		return m.InitialAmount
	}
	return 0
}

func (m *Offer) GetAvailableAmount() uint64 {
	if m != nil {
		return m.AvailableAmount
	}
	return 0
}

func (m *Offer) GetStartBlock() uint64 {
	if m != nil {
		return m.StartBlock
	}
	return 0
}

func (m *Offer) GetEndBlock() uint64 {
	if m != nil {
		return m.EndBlock
	}
	return 0
}

func init() {
	proto.RegisterEnum("topchain.subscription.Offer_Status", Offer_Status_name, Offer_Status_value)
	proto.RegisterType((*Offer)(nil), "topchain.subscription.Offer")
}

func init() { proto.RegisterFile("topchain/subscription/offer.proto", fileDescriptor_ff67a1f2f1b26c98) }

var fileDescriptor_ff67a1f2f1b26c98 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0xdb, 0x6e, 0xeb, 0xbb, 0x66, 0xbc, 0xdb, 0x08, 0x0c, 0x22, 0x6a, 0xad, 0x13, 0xa1,
	0x5e, 0x3a, 0xd0, 0x83, 0x07, 0x4f, 0x4e, 0x27, 0xf4, 0xa2, 0xd2, 0xe1, 0xc5, 0x4b, 0x49, 0x9b,
	0x0c, 0x83, 0xb5, 0xa9, 0x49, 0x2a, 0xfa, 0x2d, 0xfc, 0x58, 0x1e, 0x77, 0xf4, 0xa6, 0x6c, 0x5f,
	0x44, 0x16, 0xbb, 0x39, 0xc1, 0xeb, 0xef, 0xf9, 0xfd, 0xff, 0x3c, 0xf0, 0x80, 0x5d, 0xc5, 0x8b,
	0xf4, 0x0e, 0xb3, 0x7c, 0x20, 0xcb, 0x44, 0xa6, 0x82, 0x15, 0x8a, 0xf1, 0x7c, 0xc0, 0x27, 0x13,
	0x2a, 0x82, 0x42, 0x70, 0xc5, 0x61, 0x6f, 0xa9, 0x04, 0xeb, 0x4a, 0xff, 0xc3, 0x02, 0x8d, 0xab,
	0x85, 0x06, 0xdb, 0xc0, 0x62, 0x04, 0x99, 0x9e, 0xe9, 0x3b, 0x91, 0xc5, 0x08, 0xdc, 0x02, 0x8e,
	0xa0, 0x8f, 0x25, 0x95, 0x8a, 0x0a, 0x64, 0x69, 0xfc, 0x03, 0x60, 0x0f, 0xd8, 0xa9, 0xe0, 0x31,
	0x23, 0xa8, 0xa6, 0xa3, 0x46, 0x2a, 0x78, 0x48, 0xe0, 0x01, 0xe8, 0xae, 0xbf, 0x8f, 0x19, 0x91,
	0xa8, 0xee, 0xd5, 0x7c, 0x27, 0xea, 0xac, 0xf3, 0x90, 0x48, 0x78, 0x02, 0x6c, 0xa9, 0xb0, 0x2a,
	0x25, 0x6a, 0x78, 0xa6, 0xdf, 0x3e, 0xdc, 0x0b, 0xfe, 0x6c, 0x18, 0xe8, 0x76, 0xc1, 0x58, 0xab,
	0x51, 0x75, 0x02, 0xf7, 0x41, 0x9b, 0xe5, 0x4c, 0x31, 0x9c, 0xc5, 0xf8, 0x81, 0x97, 0xb9, 0x42,
	0xb6, 0x67, 0xfa, 0xf5, 0xe8, 0x7f, 0x45, 0x4f, 0x35, 0x5c, 0xd4, 0xc1, 0x4f, 0x98, 0x65, 0x38,
	0xc9, 0xe8, 0x52, 0xfc, 0xa7, 0xc5, 0xce, 0x8a, 0x57, 0xea, 0x0e, 0x68, 0x49, 0x85, 0x85, 0x8a,
	0x93, 0x8c, 0xa7, 0xf7, 0xa8, 0xa9, 0x2d, 0xa0, 0xd1, 0x70, 0x41, 0xe0, 0x26, 0x70, 0x68, 0x4e,
	0xaa, 0xd8, 0xd1, 0x71, 0x93, 0xe6, 0x44, 0x87, 0xfd, 0x0d, 0x60, 0x7f, 0x37, 0x84, 0x1d, 0xd0,
	0xba, 0xb9, 0x1c, 0x5f, 0x8f, 0xce, 0xc2, 0x8b, 0x70, 0x74, 0xde, 0x35, 0x86, 0xc7, 0x6f, 0x33,
	0xd7, 0x9c, 0xce, 0x5c, 0xf3, 0x73, 0xe6, 0x9a, 0xaf, 0x73, 0xd7, 0x98, 0xce, 0x5d, 0xe3, 0x7d,
	0xee, 0x1a, 0xb7, 0xdb, 0xab, 0xd5, 0x9e, 0x7f, 0xef, 0xa6, 0x5e, 0x0a, 0x2a, 0x13, 0x5b, 0x0f,
	0x77, 0xf4, 0x15, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x73, 0xa4, 0xfb, 0xdd, 0x01, 0x00, 0x00,
}

func (m *Offer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Offer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Offer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndBlock != 0 {
		i = encodeVarintOffer(dAtA, i, uint64(m.EndBlock))
		i--
		dAtA[i] = 0x48
	}
	if m.StartBlock != 0 {
		i = encodeVarintOffer(dAtA, i, uint64(m.StartBlock))
		i--
		dAtA[i] = 0x40
	}
	if m.AvailableAmount != 0 {
		i = encodeVarintOffer(dAtA, i, uint64(m.AvailableAmount))
		i--
		dAtA[i] = 0x38
	}
	if m.InitialAmount != 0 {
		i = encodeVarintOffer(dAtA, i, uint64(m.InitialAmount))
		i--
		dAtA[i] = 0x30
	}
	if m.Status != 0 {
		i = encodeVarintOffer(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.SubscriptionIds) > 0 {
		for iNdEx := len(m.SubscriptionIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SubscriptionIds[iNdEx])
			copy(dAtA[i:], m.SubscriptionIds[iNdEx])
			i = encodeVarintOffer(dAtA, i, uint64(len(m.SubscriptionIds[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.CroId) > 0 {
		i -= len(m.CroId)
		copy(dAtA[i:], m.CroId)
		i = encodeVarintOffer(dAtA, i, uint64(len(m.CroId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Requester) > 0 {
		i -= len(m.Requester)
		copy(dAtA[i:], m.Requester)
		i = encodeVarintOffer(dAtA, i, uint64(len(m.Requester)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintOffer(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOffer(dAtA []byte, offset int, v uint64) int {
	offset -= sovOffer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Offer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovOffer(uint64(l))
	}
	l = len(m.Requester)
	if l > 0 {
		n += 1 + l + sovOffer(uint64(l))
	}
	l = len(m.CroId)
	if l > 0 {
		n += 1 + l + sovOffer(uint64(l))
	}
	if len(m.SubscriptionIds) > 0 {
		for _, s := range m.SubscriptionIds {
			l = len(s)
			n += 1 + l + sovOffer(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovOffer(uint64(m.Status))
	}
	if m.InitialAmount != 0 {
		n += 1 + sovOffer(uint64(m.InitialAmount))
	}
	if m.AvailableAmount != 0 {
		n += 1 + sovOffer(uint64(m.AvailableAmount))
	}
	if m.StartBlock != 0 {
		n += 1 + sovOffer(uint64(m.StartBlock))
	}
	if m.EndBlock != 0 {
		n += 1 + sovOffer(uint64(m.EndBlock))
	}
	return n
}

func sovOffer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOffer(x uint64) (n int) {
	return sovOffer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Offer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOffer
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
			return fmt.Errorf("proto: Offer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Offer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
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
				return ErrInvalidLengthOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requester", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
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
				return ErrInvalidLengthOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requester = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CroId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
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
				return ErrInvalidLengthOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CroId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
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
				return ErrInvalidLengthOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubscriptionIds = append(m.SubscriptionIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Offer_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialAmount", wireType)
			}
			m.InitialAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvailableAmount", wireType)
			}
			m.AvailableAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AvailableAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBlock", wireType)
			}
			m.StartBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlock", wireType)
			}
			m.EndBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOffer
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
func skipOffer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOffer
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
					return 0, ErrIntOverflowOffer
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
					return 0, ErrIntOverflowOffer
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
				return 0, ErrInvalidLengthOffer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOffer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOffer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOffer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOffer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOffer = fmt.Errorf("proto: unexpected end of group")
)
