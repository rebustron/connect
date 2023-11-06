// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/alerts/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// AlertParams is the set of parameters for the x/Alerts module's Alerting. It
// defines whether or not Alerts can be submitted, and if so, the minimum
// bond amount required to submit an Alert.
type AlertParams struct {
	// Enabled is a boolean defining whether or not Alerts can be submitted
	// to the module
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// BondAmount is the minimum amount of bond required to submit an
	// Alert
	BondAmount types.Coin `protobuf:"bytes,2,opt,name=bond_amount,json=bondAmount,proto3" json:"bond_amount"`
	// MaxBlockAge defines the maximum age of an Alert before it is pruned, notice this is defined wrt.
	// the height that the Alert references, i.e Alerts are only relevant until Alert.Height + MaxBlockAge
	// is reached.
	MaxBlockAge uint64 `protobuf:"varint,3,opt,name=max_block_age,json=maxBlockAge,proto3" json:"max_block_age,omitempty"`
}

func (m *AlertParams) Reset()         { *m = AlertParams{} }
func (m *AlertParams) String() string { return proto.CompactTextString(m) }
func (*AlertParams) ProtoMessage()    {}
func (*AlertParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b18f585e730177c, []int{0}
}
func (m *AlertParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AlertParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AlertParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AlertParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertParams.Merge(m, src)
}
func (m *AlertParams) XXX_Size() int {
	return m.Size()
}
func (m *AlertParams) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertParams.DiscardUnknown(m)
}

var xxx_messageInfo_AlertParams proto.InternalMessageInfo

func (m *AlertParams) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *AlertParams) GetBondAmount() types.Coin {
	if m != nil {
		return m.BondAmount
	}
	return types.Coin{}
}

func (m *AlertParams) GetMaxBlockAge() uint64 {
	if m != nil {
		return m.MaxBlockAge
	}
	return 0
}

// PruningParams defines the criterion for pruning Alerts from the state.
type PruningParams struct {
	// Enabled defines whether Alerts are to be pruned
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// BlocksToPrune defines the number of blocks until an Alert will be pruned from state, notice this is
	// defined wrt. the current block height, i.e Alerts will be stored in state until current_height + BlocksToPrune
	// is reached.
	BlocksToPrune uint64 `protobuf:"varint,2,opt,name=blocks_to_prune,json=blocksToPrune,proto3" json:"blocks_to_prune,omitempty"`
}

func (m *PruningParams) Reset()         { *m = PruningParams{} }
func (m *PruningParams) String() string { return proto.CompactTextString(m) }
func (*PruningParams) ProtoMessage()    {}
func (*PruningParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b18f585e730177c, []int{1}
}
func (m *PruningParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PruningParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PruningParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PruningParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PruningParams.Merge(m, src)
}
func (m *PruningParams) XXX_Size() int {
	return m.Size()
}
func (m *PruningParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PruningParams.DiscardUnknown(m)
}

var xxx_messageInfo_PruningParams proto.InternalMessageInfo

func (m *PruningParams) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *PruningParams) GetBlocksToPrune() uint64 {
	if m != nil {
		return m.BlocksToPrune
	}
	return 0
}

// Params is the set of parameters for the x/Alerts module.
type Params struct {
	// AlertParams is the set of parameters for the x/Alerts module's Alerting.
	AlertParams AlertParams `protobuf:"bytes,1,opt,name=alert_params,json=alertParams,proto3" json:"alert_params"`
	// ConclusionVerificationParams is the set of parameters for the x/Alerts module's conclusion verification.
	ConclusionVerificationParams *types1.Any `protobuf:"bytes,2,opt,name=conclusion_verification_params,json=conclusionVerificationParams,proto3" json:"conclusion_verification_params,omitempty"`
	// PruningParams is the set of parameters for the x/Alerts module's pruning.
	PruningParams PruningParams `protobuf:"bytes,3,opt,name=pruning_params,json=pruningParams,proto3" json:"pruning_params"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b18f585e730177c, []int{2}
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

func (m *Params) GetAlertParams() AlertParams {
	if m != nil {
		return m.AlertParams
	}
	return AlertParams{}
}

func (m *Params) GetConclusionVerificationParams() *types1.Any {
	if m != nil {
		return m.ConclusionVerificationParams
	}
	return nil
}

func (m *Params) GetPruningParams() PruningParams {
	if m != nil {
		return m.PruningParams
	}
	return PruningParams{}
}

// GenesisState is the state that must be provided at genesis. It contains params for the module, and the set
// initial Alerts.
type GenesisState struct {
	// Params is the set of x/Alerts parameters
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// Alerts is the set of Alerts that have been submitted to the module
	Alerts []AlertWithStatus `protobuf:"bytes,2,rep,name=alerts,proto3" json:"alerts"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b18f585e730177c, []int{3}
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

func (m *GenesisState) GetAlerts() []AlertWithStatus {
	if m != nil {
		return m.Alerts
	}
	return nil
}

func init() {
	proto.RegisterType((*AlertParams)(nil), "slinky.alerts.v1.AlertParams")
	proto.RegisterType((*PruningParams)(nil), "slinky.alerts.v1.PruningParams")
	proto.RegisterType((*Params)(nil), "slinky.alerts.v1.Params")
	proto.RegisterType((*GenesisState)(nil), "slinky.alerts.v1.GenesisState")
}

func init() { proto.RegisterFile("slinky/alerts/v1/genesis.proto", fileDescriptor_7b18f585e730177c) }

var fileDescriptor_7b18f585e730177c = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0xad, 0xdb, 0xaa, 0x80, 0x33, 0xe5, 0x11, 0xcd, 0x22, 0x33, 0x62, 0x32, 0xa5, 0x0b, 0x54,
	0x90, 0x6a, 0xab, 0x45, 0x62, 0x8b, 0xda, 0x0a, 0xd8, 0xb0, 0x18, 0x02, 0x02, 0x89, 0x4d, 0xe4,
	0x64, 0x3c, 0x19, 0xab, 0x89, 0x1d, 0xc5, 0x4e, 0xd5, 0x7e, 0x01, 0x2b, 0x24, 0xc4, 0x57, 0xb0,
	0x42, 0x2c, 0xf8, 0x88, 0x11, 0xab, 0x59, 0xb2, 0x42, 0xa8, 0x5d, 0xf0, 0x1b, 0x28, 0xb6, 0x3b,
	0xea, 0x50, 0x1e, 0x9b, 0xaa, 0xe7, 0xdc, 0x7b, 0xcf, 0xb9, 0xf7, 0xb4, 0x86, 0xbe, 0x4c, 0x19,
	0x9f, 0x2e, 0x30, 0x49, 0x69, 0xa1, 0x24, 0x9e, 0x0d, 0x70, 0x42, 0x39, 0x95, 0x4c, 0xa2, 0xbc,
	0x10, 0x4a, 0xb8, 0x37, 0x4d, 0x1d, 0x99, 0x3a, 0x9a, 0x0d, 0xf6, 0xfd, 0x58, 0xc8, 0x4c, 0x48,
	0x1c, 0x11, 0x49, 0xf1, 0x6c, 0x10, 0x51, 0x45, 0x06, 0x38, 0x16, 0x8c, 0x9b, 0x89, 0xfd, 0x5b,
	0x24, 0x63, 0x5c, 0x60, 0xfd, 0x69, 0xa9, 0x3d, 0x33, 0x12, 0x6a, 0x84, 0x0d, 0xb0, 0xa5, 0xdd,
	0x44, 0x24, 0xc2, 0xf0, 0xd5, 0x37, 0xcb, 0x1e, 0x6c, 0x6d, 0x65, 0xfd, 0xad, 0x5e, 0x22, 0x44,
	0x92, 0x52, 0xac, 0x51, 0x54, 0x9e, 0x60, 0xc2, 0x17, 0xa6, 0xd4, 0xfd, 0x00, 0xa0, 0x33, 0xaa,
	0x7a, 0x8f, 0x48, 0x41, 0x32, 0xe9, 0x7a, 0xf0, 0x0a, 0xe5, 0x24, 0x4a, 0xe9, 0xb1, 0x07, 0x3a,
	0xa0, 0x77, 0x35, 0x58, 0x43, 0xf7, 0x31, 0x74, 0x22, 0xc1, 0x8f, 0x43, 0x92, 0x89, 0x92, 0x2b,
	0xaf, 0xde, 0x01, 0x3d, 0x67, 0xb8, 0x87, 0xec, 0x76, 0xd5, 0x75, 0xc8, 0x5e, 0x87, 0x26, 0x82,
	0xf1, 0xf1, 0xb5, 0xb3, 0xef, 0x87, 0xb5, 0x8f, 0x3f, 0x3f, 0xdf, 0x07, 0x01, 0xac, 0x06, 0x47,
	0x7a, 0xce, 0xed, 0xc2, 0x76, 0x46, 0xe6, 0x61, 0x94, 0x8a, 0x78, 0x1a, 0x92, 0x84, 0x7a, 0x8d,
	0x0e, 0xe8, 0x35, 0x03, 0x27, 0x23, 0xf3, 0x71, 0xc5, 0x8d, 0x12, 0xda, 0x7d, 0x0e, 0xdb, 0x47,
	0x45, 0xc9, 0x19, 0x4f, 0xfe, 0xbb, 0xd5, 0x5d, 0x78, 0x43, 0x4b, 0xc9, 0x50, 0x89, 0x30, 0x2f,
	0x4a, 0x4e, 0xf5, 0x66, 0xcd, 0xa0, 0x6d, 0xe8, 0x97, 0xa2, 0x52, 0xa2, 0xdd, 0x4f, 0x75, 0xd8,
	0xb2, 0x62, 0x4f, 0xe0, 0x8e, 0x4e, 0x27, 0xcc, 0x35, 0xd6, 0x8a, 0xce, 0xf0, 0x00, 0xfd, 0xfe,
	0xcb, 0xa1, 0x8d, 0x5c, 0xc6, 0xcd, 0xea, 0x9a, 0xc0, 0x21, 0x1b, 0x51, 0xbd, 0x03, 0xd0, 0x8f,
	0x05, 0x8f, 0xd3, 0x52, 0x32, 0xc1, 0xc3, 0x19, 0x2d, 0xd8, 0x09, 0x8b, 0x89, 0xaa, 0x80, 0x95,
	0x36, 0x21, 0xed, 0x22, 0x93, 0x3f, 0x5a, 0xe7, 0x8f, 0x46, 0x7c, 0x31, 0x1e, 0x7c, 0xfd, 0xd2,
	0xef, 0x6f, 0x79, 0x4e, 0x2e, 0x04, 0x5f, 0x6d, 0xe8, 0x19, 0xc7, 0xe0, 0x76, 0xfc, 0x8f, 0xaa,
	0xfb, 0x0c, 0x5e, 0xcf, 0x4d, 0x6a, 0x6b, 0xfb, 0x86, 0xb6, 0x3f, 0xdc, 0xbe, 0xec, 0x52, 0xba,
	0xf6, 0xb6, 0x76, 0xbe, 0x49, 0x76, 0xdf, 0x02, 0xb8, 0xf3, 0xd4, 0xfc, 0xb5, 0x5f, 0x28, 0xa2,
	0xa8, 0xfb, 0x10, 0xb6, 0x2e, 0x05, 0xe6, 0xfd, 0x41, 0x76, 0x53, 0xcf, 0x76, 0xbb, 0x8f, 0x60,
	0xcb, 0x74, 0x78, 0xf5, 0x4e, 0xa3, 0xe7, 0x0c, 0xef, 0xfc, 0x25, 0xe8, 0xd7, 0x4c, 0x9d, 0x56,
	0x4e, 0xe5, 0x85, 0x80, 0x69, 0x18, 0x4f, 0xce, 0x96, 0x3e, 0x38, 0x5f, 0xfa, 0xe0, 0xc7, 0xd2,
	0x07, 0xef, 0x57, 0x7e, 0xed, 0x7c, 0xe5, 0xd7, 0xbe, 0xad, 0xfc, 0xda, 0x9b, 0x7b, 0x09, 0x53,
	0xa7, 0x65, 0x84, 0x62, 0x91, 0x61, 0x39, 0x65, 0x79, 0x3f, 0xa3, 0x33, 0x6c, 0x9f, 0xc2, 0x7c,
	0xfd, 0x18, 0xd4, 0x22, 0xa7, 0x32, 0x6a, 0xe9, 0xec, 0x1f, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0xce, 0x74, 0xdb, 0x08, 0xc0, 0x03, 0x00, 0x00,
}

func (m *AlertParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AlertParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AlertParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxBlockAge != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxBlockAge))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.BondAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PruningParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PruningParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PruningParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlocksToPrune != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BlocksToPrune))
		i--
		dAtA[i] = 0x10
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
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
		size, err := m.PruningParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.ConclusionVerificationParams != nil {
		{
			size, err := m.ConclusionVerificationParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.AlertParams.MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Alerts) > 0 {
		for iNdEx := len(m.Alerts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Alerts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *AlertParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	l = m.BondAmount.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.MaxBlockAge != 0 {
		n += 1 + sovGenesis(uint64(m.MaxBlockAge))
	}
	return n
}

func (m *PruningParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	if m.BlocksToPrune != 0 {
		n += 1 + sovGenesis(uint64(m.BlocksToPrune))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AlertParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.ConclusionVerificationParams != nil {
		l = m.ConclusionVerificationParams.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.PruningParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Alerts) > 0 {
		for _, e := range m.Alerts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AlertParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AlertParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AlertParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondAmount", wireType)
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
			if err := m.BondAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBlockAge", wireType)
			}
			m.MaxBlockAge = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxBlockAge |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *PruningParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PruningParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PruningParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksToPrune", wireType)
			}
			m.BlocksToPrune = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksToPrune |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlertParams", wireType)
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
			if err := m.AlertParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConclusionVerificationParams", wireType)
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
			if m.ConclusionVerificationParams == nil {
				m.ConclusionVerificationParams = &types1.Any{}
			}
			if err := m.ConclusionVerificationParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PruningParams", wireType)
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
			if err := m.PruningParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Alerts", wireType)
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
			m.Alerts = append(m.Alerts, AlertWithStatus{})
			if err := m.Alerts[len(m.Alerts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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