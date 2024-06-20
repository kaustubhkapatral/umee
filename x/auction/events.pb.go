// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/auction/v1/events.proto

package auction

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
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

// EventRewardsAuctionResult is emitted at the end of each auction that has at least one bidder.
type EventRewardsAuctionResult struct {
	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Bidder string `protobuf:"bytes,2,opt,name=bidder,proto3" json:"bidder,omitempty"`
}

func (m *EventRewardsAuctionResult) Reset()         { *m = EventRewardsAuctionResult{} }
func (m *EventRewardsAuctionResult) String() string { return proto.CompactTextString(m) }
func (*EventRewardsAuctionResult) ProtoMessage()    {}
func (*EventRewardsAuctionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5998c755af8caa1, []int{0}
}
func (m *EventRewardsAuctionResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRewardsAuctionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRewardsAuctionResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRewardsAuctionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRewardsAuctionResult.Merge(m, src)
}
func (m *EventRewardsAuctionResult) XXX_Size() int {
	return m.Size()
}
func (m *EventRewardsAuctionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRewardsAuctionResult.DiscardUnknown(m)
}

var xxx_messageInfo_EventRewardsAuctionResult proto.InternalMessageInfo

// EventFundRewardsAuction is emitted when sending rewards to auction module
type EventFundRewardsAuction struct {
	Assets []types.Coin `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets"`
}

func (m *EventFundRewardsAuction) Reset()         { *m = EventFundRewardsAuction{} }
func (m *EventFundRewardsAuction) String() string { return proto.CompactTextString(m) }
func (*EventFundRewardsAuction) ProtoMessage()    {}
func (*EventFundRewardsAuction) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5998c755af8caa1, []int{1}
}
func (m *EventFundRewardsAuction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFundRewardsAuction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFundRewardsAuction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventFundRewardsAuction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFundRewardsAuction.Merge(m, src)
}
func (m *EventFundRewardsAuction) XXX_Size() int {
	return m.Size()
}
func (m *EventFundRewardsAuction) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFundRewardsAuction.DiscardUnknown(m)
}

var xxx_messageInfo_EventFundRewardsAuction proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventRewardsAuctionResult)(nil), "umee.auction.v1.EventRewardsAuctionResult")
	proto.RegisterType((*EventFundRewardsAuction)(nil), "umee.auction.v1.EventFundRewardsAuction")
}

func init() { proto.RegisterFile("umee/auction/v1/events.proto", fileDescriptor_b5998c755af8caa1) }

var fileDescriptor_b5998c755af8caa1 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x27, 0xfd, 0x7f, 0x0a, 0x46, 0x54, 0x18, 0x0a, 0x4e, 0x8b, 0xc4, 0xd2, 0x55, 0x5d,
	0x34, 0x71, 0x14, 0x74, 0xdd, 0x8a, 0xba, 0x8f, 0x3b, 0x41, 0x64, 0x66, 0x72, 0x19, 0x83, 0x36,
	0x91, 0x24, 0x33, 0xf5, 0x31, 0x7c, 0x18, 0x1f, 0xa2, 0xcb, 0xe2, 0xca, 0x95, 0x68, 0xfb, 0x22,
	0x32, 0x93, 0xb8, 0x70, 0x77, 0x0f, 0xdf, 0xe1, 0x9e, 0xc3, 0xc1, 0x07, 0xd5, 0x1c, 0x80, 0x65,
	0x55, 0xe1, 0xa4, 0x56, 0xac, 0x4e, 0x19, 0xd4, 0xa0, 0x9c, 0xa5, 0xcf, 0x46, 0x3b, 0x1d, 0xef,
	0x35, 0x94, 0x06, 0x4a, 0xeb, 0x74, 0xd0, 0x2b, 0x75, 0xa9, 0x5b, 0xc6, 0x9a, 0xcb, 0xdb, 0x06,
	0xfd, 0x42, 0xdb, 0xb9, 0xb6, 0xf7, 0x1e, 0x78, 0x11, 0x10, 0xf1, 0x8a, 0xe5, 0x99, 0x05, 0x56,
	0xa7, 0x39, 0xb8, 0x2c, 0x65, 0x85, 0x96, 0xca, 0xf3, 0xd1, 0x1d, 0xee, 0x5f, 0x36, 0x89, 0x1c,
	0x16, 0x99, 0x11, 0x76, 0xea, 0xa3, 0x38, 0xd8, 0xea, 0xc9, 0xc5, 0xbb, 0xb8, 0x23, 0x45, 0x82,
	0x86, 0x68, 0xbc, 0xc3, 0x3b, 0x52, 0xc4, 0xc7, 0xb8, 0x9b, 0x4b, 0x21, 0xc0, 0x24, 0x9d, 0x21,
	0x1a, 0x6f, 0xcd, 0x92, 0xf7, 0xb7, 0x49, 0x2f, 0xc4, 0x4d, 0x85, 0x30, 0x60, 0xed, 0x8d, 0x33,
	0x52, 0x95, 0x3c, 0xf8, 0x46, 0x1c, 0xef, 0xb7, 0xef, 0xaf, 0x2a, 0x25, 0xfe, 0x46, 0xc4, 0xe7,
	0xb8, 0x9b, 0x59, 0x0b, 0xce, 0x26, 0x68, 0xf8, 0x6f, 0xbc, 0x7d, 0xd2, 0xa7, 0xe1, 0x53, 0x53,
	0x95, 0x86, 0xaa, 0xf4, 0x42, 0x4b, 0x35, 0xfb, 0xbf, 0xfc, 0x3c, 0x8c, 0x78, 0xb0, 0xcf, 0xae,
	0x97, 0xdf, 0x24, 0x5a, 0xae, 0x09, 0x5a, 0xad, 0x09, 0xfa, 0x5a, 0x13, 0xf4, 0xba, 0x21, 0xd1,
	0x6a, 0x43, 0xa2, 0x8f, 0x0d, 0x89, 0x6e, 0x8f, 0x4a, 0xe9, 0x1e, 0xaa, 0x9c, 0x16, 0x7a, 0xce,
	0x9a, 0xf5, 0x26, 0x0a, 0xdc, 0x42, 0x9b, 0xc7, 0x56, 0xb0, 0xfa, 0x8c, 0xbd, 0xfc, 0xae, 0x9d,
	0x77, 0xdb, 0x09, 0x4e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xe5, 0xa1, 0x74, 0x84, 0x01,
	0x00, 0x00,
}

func (m *EventRewardsAuctionResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRewardsAuctionResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRewardsAuctionResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventFundRewardsAuction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFundRewardsAuction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventFundRewardsAuction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for iNdEx := len(m.Assets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Assets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventRewardsAuctionResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventFundRewardsAuction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for _, e := range m.Assets {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventRewardsAuctionResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRewardsAuctionResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRewardsAuctionResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventFundRewardsAuction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventFundRewardsAuction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFundRewardsAuction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Assets = append(m.Assets, types.Coin{})
			if err := m.Assets[len(m.Assets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)