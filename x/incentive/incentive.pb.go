// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/incentive/v1/incentive.proto

package incentive

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the incentive module.
type Params struct {
	// max_unbondings is the maximum amount of concurrent unbondings an address can have
	// of each bonded uToken denom. Zero is interpreted as no limit.
	MaxUnbondings uint32 `protobuf:"varint,1,opt,name=max_unbondings,json=maxUnbondings,proto3" json:"max_unbondings,omitempty"`
	// unbonding_duration is the unbonding duration (in seconds).
	UnbondingDuration int64 `protobuf:"varint,2,opt,name=unbonding_duration,json=unbondingDuration,proto3" json:"unbonding_duration,omitempty"`
	// emergency_unbond_fee is the portion of a bond that is paid when it is instantly
	// released using MsgEmergencyUnbond. For example, 0.01 is a 1% fee. Ranges 0-1.
	EmergencyUnbondFee github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=emergency_unbond_fee,json=emergencyUnbondFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"emergency_unbond_fee"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c99c623956e199b, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// IncentiveProgram defines a liquidity mining incentive program on a single
// locked uToken denom that will run for a set amount of time.
type IncentiveProgram struct {
	// ID uniquely identifies the incentive program after it has been created.
	// It is zero when the program is being proposed by governance, and is set
	// to its final value when the proposal passes.
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// start_time is the unix time (in seconds) at which the incentives begin.
	// If a program is passed after its intended start time, its start time
	// will be increased to the current time, with program duration unchanged.
	StartTime int64 `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// duration is the length of the incentive program from start time to
	// completion in seconds.
	Duration int64 `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	// uToken is the incentivized uToken collateral denom. Suppliers who collateralize
	// this asset then bond it to the incentive module are eligible for this program's
	// rewards.
	UToken string `protobuf:"bytes,4,opt,name=uToken,proto3" json:"uToken,omitempty"`
	// funded indicates whether a program bas been funded. This can happen when
	// a program passes if funding from community fund, or any time before the
	// program's start time if funding with MsgSponsor. A program that reaches
	// its start time without being funded is cancelled.
	Funded bool `protobuf:"varint,5,opt,name=funded,proto3" json:"funded,omitempty"`
	// total_rewards are total amount of rewards which can be distributed to
	// suppliers by this program. This is set to its final value when the program
	// is proposed by governance.
	TotalRewards types.Coin `protobuf:"bytes,6,opt,name=total_rewards,json=totalRewards,proto3" json:"total_rewards"`
	// remaining_rewards are total amount of this program's funded rewards
	// which have not yet been allocated to suppliers. This is zero until the
	// program is both passed by governance and funded, at which point it
	// starts at the same value as total_rewards then begins decreasing
	// to zero as the program runs to completion.
	RemainingRewards types.Coin `protobuf:"bytes,7,opt,name=remaining_rewards,json=remainingRewards,proto3" json:"remaining_rewards"`
}

func (m *IncentiveProgram) Reset()         { *m = IncentiveProgram{} }
func (m *IncentiveProgram) String() string { return proto.CompactTextString(m) }
func (*IncentiveProgram) ProtoMessage()    {}
func (*IncentiveProgram) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c99c623956e199b, []int{1}
}
func (m *IncentiveProgram) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentiveProgram) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentiveProgram.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentiveProgram) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentiveProgram.Merge(m, src)
}
func (m *IncentiveProgram) XXX_Size() int {
	return m.Size()
}
func (m *IncentiveProgram) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentiveProgram.DiscardUnknown(m)
}

var xxx_messageInfo_IncentiveProgram proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "umee.incentive.v1.Params")
	proto.RegisterType((*IncentiveProgram)(nil), "umee.incentive.v1.IncentiveProgram")
}

func init() { proto.RegisterFile("umee/incentive/v1/incentive.proto", fileDescriptor_8c99c623956e199b) }

var fileDescriptor_8c99c623956e199b = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0xce, 0xb4, 0xb5, 0xee, 0x8e, 0x76, 0xd9, 0x0e, 0x8b, 0xc4, 0x82, 0x69, 0x5c, 0x50, 0x02,
	0xd2, 0x84, 0x2a, 0x78, 0xf0, 0x58, 0x83, 0x50, 0xf0, 0xb0, 0x84, 0xf5, 0xe2, 0x25, 0x4e, 0x92,
	0x77, 0xe3, 0x50, 0x67, 0x66, 0x99, 0x4c, 0xba, 0xdd, 0x7f, 0xe1, 0x4f, 0xf0, 0xe6, 0x3f, 0x91,
	0x1e, 0xf7, 0x28, 0x1e, 0x16, 0x6d, 0x2f, 0xfe, 0x0c, 0x49, 0x32, 0x4d, 0x3d, 0xee, 0x69, 0xe6,
	0x79, 0xde, 0xf7, 0x79, 0x78, 0xbf, 0xf0, 0xd3, 0x92, 0x03, 0x04, 0x4c, 0xa4, 0x20, 0x34, 0x5b,
	0x42, 0xb0, 0x9c, 0xee, 0x81, 0x7f, 0xa9, 0xa4, 0x96, 0x64, 0x58, 0xa5, 0xf8, 0x7b, 0x76, 0x39,
	0x1d, 0x39, 0xa9, 0x2c, 0xb8, 0x2c, 0x82, 0x84, 0x16, 0x95, 0x24, 0x01, 0x4d, 0xa7, 0x41, 0x2a,
	0x99, 0x68, 0x24, 0xa3, 0x93, 0x5c, 0xe6, 0xb2, 0xfe, 0x06, 0xd5, 0xaf, 0x61, 0x4f, 0x7f, 0x20,
	0xdc, 0x3f, 0xa3, 0x8a, 0xf2, 0x82, 0x3c, 0xc3, 0x47, 0x9c, 0xae, 0xe2, 0x52, 0x24, 0x52, 0x64,
	0x4c, 0xe4, 0x85, 0x8d, 0x5c, 0xe4, 0x0d, 0xa2, 0x01, 0xa7, 0xab, 0x0f, 0x2d, 0x49, 0x26, 0x98,
	0xb4, 0x29, 0x71, 0x56, 0x2a, 0xaa, 0x99, 0x14, 0x76, 0xc7, 0x45, 0x5e, 0x37, 0x1a, 0xb6, 0x91,
	0xd0, 0x04, 0xc8, 0x27, 0x7c, 0x02, 0x1c, 0x54, 0x0e, 0x22, 0xbd, 0x36, 0xde, 0xf1, 0x05, 0x80,
	0xdd, 0x75, 0x91, 0x77, 0x38, 0xf3, 0xd7, 0xb7, 0x63, 0xeb, 0xd7, 0xed, 0xf8, 0x79, 0xce, 0xf4,
	0xe7, 0x32, 0xf1, 0x53, 0xc9, 0x03, 0xd3, 0x47, 0xf3, 0x4c, 0x8a, 0x6c, 0x11, 0xe8, 0xeb, 0x4b,
	0x28, 0xfc, 0x10, 0xd2, 0x88, 0xb4, 0x5e, 0x4d, 0x45, 0xef, 0x00, 0xde, 0xf4, 0xfe, 0x7e, 0x1b,
	0xa3, 0xd3, 0xef, 0x1d, 0x7c, 0x3c, 0xdf, 0xcd, 0xe3, 0x4c, 0xc9, 0x5c, 0x51, 0x4e, 0x8e, 0x70,
	0x67, 0x1e, 0x9a, 0x36, 0x3a, 0xf3, 0x90, 0x3c, 0xc1, 0xb8, 0xd0, 0x54, 0xe9, 0x58, 0x33, 0x0e,
	0xa6, 0xe6, 0xc3, 0x9a, 0x39, 0x67, 0x1c, 0xc8, 0x08, 0x1f, 0xb4, 0x0d, 0x75, 0xeb, 0x60, 0x8b,
	0xc9, 0x23, 0xdc, 0x2f, 0xcf, 0xe5, 0x02, 0x84, 0xdd, 0xab, 0x2a, 0x8f, 0x0c, 0xaa, 0xf8, 0x8b,
	0x52, 0x64, 0x90, 0xd9, 0xf7, 0x5c, 0xe4, 0x1d, 0x44, 0x06, 0x91, 0x10, 0x0f, 0xb4, 0xd4, 0xf4,
	0x4b, 0xac, 0xe0, 0x8a, 0xaa, 0xac, 0xb0, 0xfb, 0x2e, 0xf2, 0x1e, 0xbc, 0x7c, 0xec, 0x37, 0x7d,
	0xf9, 0xd5, 0x9a, 0x7c, 0xb3, 0x26, 0xff, 0xad, 0x64, 0x62, 0xd6, 0xab, 0x66, 0x11, 0x3d, 0xac,
	0x55, 0x51, 0x23, 0x22, 0xef, 0xf1, 0x50, 0x01, 0xa7, 0x4c, 0x54, 0xc3, 0xde, 0x39, 0xdd, 0xbf,
	0x9b, 0xd3, 0x71, 0xab, 0x34, 0x6e, 0xcd, 0xa4, 0x66, 0xf3, 0xf5, 0x1f, 0xc7, 0x5a, 0x6f, 0x1c,
	0x74, 0xb3, 0x71, 0xd0, 0xef, 0x8d, 0x83, 0xbe, 0x6e, 0x1d, 0xeb, 0x66, 0xeb, 0x58, 0x3f, 0xb7,
	0x8e, 0xf5, 0xf1, 0xc5, 0x7f, 0x9b, 0xa8, 0x8e, 0x6c, 0x22, 0x40, 0x5f, 0x49, 0xb5, 0xa8, 0x41,
	0xb0, 0x7c, 0x1d, 0xac, 0xf6, 0xc7, 0x98, 0xf4, 0xeb, 0x23, 0x7a, 0xf5, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x14, 0xe7, 0x13, 0x90, 0xb2, 0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MaxUnbondings != that1.MaxUnbondings {
		return false
	}
	if this.UnbondingDuration != that1.UnbondingDuration {
		return false
	}
	if !this.EmergencyUnbondFee.Equal(that1.EmergencyUnbondFee) {
		return false
	}
	return true
}
func (this *IncentiveProgram) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IncentiveProgram)
	if !ok {
		that2, ok := that.(IncentiveProgram)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if this.StartTime != that1.StartTime {
		return false
	}
	if this.Duration != that1.Duration {
		return false
	}
	if this.UToken != that1.UToken {
		return false
	}
	if this.Funded != that1.Funded {
		return false
	}
	if !this.TotalRewards.Equal(&that1.TotalRewards) {
		return false
	}
	if !this.RemainingRewards.Equal(&that1.RemainingRewards) {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.EmergencyUnbondFee.Size()
		i -= size
		if _, err := m.EmergencyUnbondFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.UnbondingDuration != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.UnbondingDuration))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxUnbondings != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.MaxUnbondings))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *IncentiveProgram) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentiveProgram) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentiveProgram) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.RemainingRewards.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.TotalRewards.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.Funded {
		i--
		if m.Funded {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.UToken) > 0 {
		i -= len(m.UToken)
		copy(dAtA[i:], m.UToken)
		i = encodeVarintIncentive(dAtA, i, uint64(len(m.UToken)))
		i--
		dAtA[i] = 0x22
	}
	if m.Duration != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTime != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintIncentive(dAtA []byte, offset int, v uint64) int {
	offset -= sovIncentive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxUnbondings != 0 {
		n += 1 + sovIncentive(uint64(m.MaxUnbondings))
	}
	if m.UnbondingDuration != 0 {
		n += 1 + sovIncentive(uint64(m.UnbondingDuration))
	}
	l = m.EmergencyUnbondFee.Size()
	n += 1 + l + sovIncentive(uint64(l))
	return n
}

func (m *IncentiveProgram) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovIncentive(uint64(m.ID))
	}
	if m.StartTime != 0 {
		n += 1 + sovIncentive(uint64(m.StartTime))
	}
	if m.Duration != 0 {
		n += 1 + sovIncentive(uint64(m.Duration))
	}
	l = len(m.UToken)
	if l > 0 {
		n += 1 + l + sovIncentive(uint64(l))
	}
	if m.Funded {
		n += 2
	}
	l = m.TotalRewards.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.RemainingRewards.Size()
	n += 1 + l + sovIncentive(uint64(l))
	return n
}

func sovIncentive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIncentive(x uint64) (n int) {
	return sovIncentive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxUnbondings", wireType)
			}
			m.MaxUnbondings = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxUnbondings |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingDuration", wireType)
			}
			m.UnbondingDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingDuration |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmergencyUnbondFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EmergencyUnbondFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func (m *IncentiveProgram) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: IncentiveProgram: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentiveProgram: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funded", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Funded = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemainingRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func skipIncentive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
				return 0, ErrInvalidLengthIncentive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIncentive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIncentive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIncentive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIncentive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIncentive = fmt.Errorf("proto: unexpected end of group")
)
