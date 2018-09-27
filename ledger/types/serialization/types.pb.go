// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ledger/types/serialization/types.proto

/*
Package serialization is a generated protocol buffer package.

It is generated from these files:
	ledger/types/serialization/types.proto

It has these top-level messages:
	PublicKey
	PrivateKey
	Signature
	Validator
	OverspendingProof
	Coin
	ReservedFund
	TransferRecord
	Account
	TxInput
	TxOutput
	Tx
	CoinbaseTx
	SendTx
	ReserveFundTx
	ReleaseFundTx
	ServicePaymentTx
	SlashTx
	Split
	SplitContract
	SplitContractTx
	UpdateValidatorsTx
*/
package serialization

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PublicKey struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *PublicKey) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PrivateKey struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PrivateKey) Reset()                    { *m = PrivateKey{} }
func (m *PrivateKey) String() string            { return proto.CompactTextString(m) }
func (*PrivateKey) ProtoMessage()               {}
func (*PrivateKey) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *PrivateKey) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Signature struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

func (m *Signature) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Validator struct {
	Id    []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Stake int64  `protobuf:"varint,2,opt,name=stake,proto3" json:"stake,omitempty"`
}

func (m *Validator) Reset()                    { *m = Validator{} }
func (m *Validator) String() string            { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()               {}
func (*Validator) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{3} }

func (m *Validator) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Validator) GetStake() int64 {
	if m != nil {
		return m.Stake
	}
	return 0
}

type OverspendingProof struct {
	ReserveSequence int64               `protobuf:"varint,1,opt,name=reserve_sequence,json=reserveSequence,proto3" json:"reserve_sequence,omitempty"`
	ServicePayments []*ServicePaymentTx `protobuf:"bytes,2,rep,name=service_payments,json=servicePayments" json:"service_payments,omitempty"`
}

func (m *OverspendingProof) Reset()                    { *m = OverspendingProof{} }
func (m *OverspendingProof) String() string            { return proto.CompactTextString(m) }
func (*OverspendingProof) ProtoMessage()               {}
func (*OverspendingProof) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{4} }

func (m *OverspendingProof) GetReserveSequence() int64 {
	if m != nil {
		return m.ReserveSequence
	}
	return 0
}

func (m *OverspendingProof) GetServicePayments() []*ServicePaymentTx {
	if m != nil {
		return m.ServicePayments
	}
	return nil
}

type Coin struct {
	Denom  string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *Coin) Reset()                    { *m = Coin{} }
func (m *Coin) String() string            { return proto.CompactTextString(m) }
func (*Coin) ProtoMessage()               {}
func (*Coin) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{5} }

func (m *Coin) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Coin) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type ReservedFund struct {
	Collateral      []*Coin           `protobuf:"bytes,1,rep,name=collateral" json:"collateral,omitempty"`
	InitialFund     []*Coin           `protobuf:"bytes,2,rep,name=initial_fund,json=initialFund" json:"initial_fund,omitempty"`
	UsedFund        []*Coin           `protobuf:"bytes,3,rep,name=used_fund,json=usedFund" json:"used_fund,omitempty"`
	ResourceIds     [][]byte          `protobuf:"bytes,4,rep,name=resource_ids,json=resourceIds" json:"resource_ids,omitempty"`
	EndBlockHeight  int64             `protobuf:"varint,5,opt,name=end_block_height,json=endBlockHeight,proto3" json:"end_block_height,omitempty"`
	ReserveSequence int64             `protobuf:"varint,6,opt,name=reserve_sequence,json=reserveSequence,proto3" json:"reserve_sequence,omitempty"`
	TransferRecord  []*TransferRecord `protobuf:"bytes,7,rep,name=transfer_record,json=transferRecord" json:"transfer_record,omitempty"`
}

func (m *ReservedFund) Reset()                    { *m = ReservedFund{} }
func (m *ReservedFund) String() string            { return proto.CompactTextString(m) }
func (*ReservedFund) ProtoMessage()               {}
func (*ReservedFund) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{6} }

func (m *ReservedFund) GetCollateral() []*Coin {
	if m != nil {
		return m.Collateral
	}
	return nil
}

func (m *ReservedFund) GetInitialFund() []*Coin {
	if m != nil {
		return m.InitialFund
	}
	return nil
}

func (m *ReservedFund) GetUsedFund() []*Coin {
	if m != nil {
		return m.UsedFund
	}
	return nil
}

func (m *ReservedFund) GetResourceIds() [][]byte {
	if m != nil {
		return m.ResourceIds
	}
	return nil
}

func (m *ReservedFund) GetEndBlockHeight() int64 {
	if m != nil {
		return m.EndBlockHeight
	}
	return 0
}

func (m *ReservedFund) GetReserveSequence() int64 {
	if m != nil {
		return m.ReserveSequence
	}
	return 0
}

func (m *ReservedFund) GetTransferRecord() []*TransferRecord {
	if m != nil {
		return m.TransferRecord
	}
	return nil
}

type TransferRecord struct {
	ServicePayment *ServicePaymentTx `protobuf:"bytes,1,opt,name=service_payment,json=servicePayment" json:"service_payment,omitempty"`
}

func (m *TransferRecord) Reset()                    { *m = TransferRecord{} }
func (m *TransferRecord) String() string            { return proto.CompactTextString(m) }
func (*TransferRecord) ProtoMessage()               {}
func (*TransferRecord) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{7} }

func (m *TransferRecord) GetServicePayment() *ServicePaymentTx {
	if m != nil {
		return m.ServicePayment
	}
	return nil
}

type Account struct {
	Sequence               int64           `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Balance                []*Coin         `protobuf:"bytes,2,rep,name=balance" json:"balance,omitempty"`
	ReservedFunds          []*ReservedFund `protobuf:"bytes,3,rep,name=reserved_funds,json=reservedFunds" json:"reserved_funds,omitempty"`
	LastUpdatedBlockHeight int64           `protobuf:"varint,4,opt,name=last_updated_block_height,json=lastUpdatedBlockHeight,proto3" json:"last_updated_block_height,omitempty"`
	PubKey                 *PublicKey      `protobuf:"bytes,5,opt,name=pub_key,json=pubKey" json:"pub_key,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{8} }

func (m *Account) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Account) GetBalance() []*Coin {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *Account) GetReservedFunds() []*ReservedFund {
	if m != nil {
		return m.ReservedFunds
	}
	return nil
}

func (m *Account) GetLastUpdatedBlockHeight() int64 {
	if m != nil {
		return m.LastUpdatedBlockHeight
	}
	return 0
}

func (m *Account) GetPubKey() *PublicKey {
	if m != nil {
		return m.PubKey
	}
	return nil
}

type TxInput struct {
	Address   []byte     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Coins     []*Coin    `protobuf:"bytes,2,rep,name=coins" json:"coins,omitempty"`
	Sequence  int64      `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Signature *Signature `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	Pubkey    *PublicKey `protobuf:"bytes,5,opt,name=pubkey" json:"pubkey,omitempty"`
}

func (m *TxInput) Reset()                    { *m = TxInput{} }
func (m *TxInput) String() string            { return proto.CompactTextString(m) }
func (*TxInput) ProtoMessage()               {}
func (*TxInput) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{9} }

func (m *TxInput) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *TxInput) GetCoins() []*Coin {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *TxInput) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *TxInput) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *TxInput) GetPubkey() *PublicKey {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

type TxOutput struct {
	Address []byte  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Coins   []*Coin `protobuf:"bytes,2,rep,name=coins" json:"coins,omitempty"`
}

func (m *TxOutput) Reset()                    { *m = TxOutput{} }
func (m *TxOutput) String() string            { return proto.CompactTextString(m) }
func (*TxOutput) ProtoMessage()               {}
func (*TxOutput) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{10} }

func (m *TxOutput) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *TxOutput) GetCoins() []*Coin {
	if m != nil {
		return m.Coins
	}
	return nil
}

type Tx struct {
	// Types that are valid to be assigned to Tx:
	//	*Tx_Coinbase
	//	*Tx_Send
	//	*Tx_Reserve
	//	*Tx_Release
	//	*Tx_ServicePayment
	//	*Tx_Slash
	//	*Tx_SplitContract
	//	*Tx_UpdateValidators
	Tx isTx_Tx `protobuf_oneof:"tx"`
}

func (m *Tx) Reset()                    { *m = Tx{} }
func (m *Tx) String() string            { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()               {}
func (*Tx) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{11} }

type isTx_Tx interface {
	isTx_Tx()
}

type Tx_Coinbase struct {
	Coinbase *CoinbaseTx `protobuf:"bytes,1,opt,name=coinbase,oneof"`
}
type Tx_Send struct {
	Send *SendTx `protobuf:"bytes,2,opt,name=send,oneof"`
}
type Tx_Reserve struct {
	Reserve *ReserveFundTx `protobuf:"bytes,3,opt,name=reserve,oneof"`
}
type Tx_Release struct {
	Release *ReleaseFundTx `protobuf:"bytes,4,opt,name=release,oneof"`
}
type Tx_ServicePayment struct {
	ServicePayment *ServicePaymentTx `protobuf:"bytes,5,opt,name=service_payment,json=servicePayment,oneof"`
}
type Tx_Slash struct {
	Slash *SlashTx `protobuf:"bytes,6,opt,name=slash,oneof"`
}
type Tx_SplitContract struct {
	SplitContract *SplitContractTx `protobuf:"bytes,7,opt,name=split_contract,json=splitContract,oneof"`
}
type Tx_UpdateValidators struct {
	UpdateValidators *UpdateValidatorsTx `protobuf:"bytes,8,opt,name=update_validators,json=updateValidators,oneof"`
}

func (*Tx_Coinbase) isTx_Tx()         {}
func (*Tx_Send) isTx_Tx()             {}
func (*Tx_Reserve) isTx_Tx()          {}
func (*Tx_Release) isTx_Tx()          {}
func (*Tx_ServicePayment) isTx_Tx()   {}
func (*Tx_Slash) isTx_Tx()            {}
func (*Tx_SplitContract) isTx_Tx()    {}
func (*Tx_UpdateValidators) isTx_Tx() {}

func (m *Tx) GetTx() isTx_Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *Tx) GetCoinbase() *CoinbaseTx {
	if x, ok := m.GetTx().(*Tx_Coinbase); ok {
		return x.Coinbase
	}
	return nil
}

func (m *Tx) GetSend() *SendTx {
	if x, ok := m.GetTx().(*Tx_Send); ok {
		return x.Send
	}
	return nil
}

func (m *Tx) GetReserve() *ReserveFundTx {
	if x, ok := m.GetTx().(*Tx_Reserve); ok {
		return x.Reserve
	}
	return nil
}

func (m *Tx) GetRelease() *ReleaseFundTx {
	if x, ok := m.GetTx().(*Tx_Release); ok {
		return x.Release
	}
	return nil
}

func (m *Tx) GetServicePayment() *ServicePaymentTx {
	if x, ok := m.GetTx().(*Tx_ServicePayment); ok {
		return x.ServicePayment
	}
	return nil
}

func (m *Tx) GetSlash() *SlashTx {
	if x, ok := m.GetTx().(*Tx_Slash); ok {
		return x.Slash
	}
	return nil
}

func (m *Tx) GetSplitContract() *SplitContractTx {
	if x, ok := m.GetTx().(*Tx_SplitContract); ok {
		return x.SplitContract
	}
	return nil
}

func (m *Tx) GetUpdateValidators() *UpdateValidatorsTx {
	if x, ok := m.GetTx().(*Tx_UpdateValidators); ok {
		return x.UpdateValidators
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Tx) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Tx_OneofMarshaler, _Tx_OneofUnmarshaler, _Tx_OneofSizer, []interface{}{
		(*Tx_Coinbase)(nil),
		(*Tx_Send)(nil),
		(*Tx_Reserve)(nil),
		(*Tx_Release)(nil),
		(*Tx_ServicePayment)(nil),
		(*Tx_Slash)(nil),
		(*Tx_SplitContract)(nil),
		(*Tx_UpdateValidators)(nil),
	}
}

func _Tx_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Tx)
	// tx
	switch x := m.Tx.(type) {
	case *Tx_Coinbase:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Coinbase); err != nil {
			return err
		}
	case *Tx_Send:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Send); err != nil {
			return err
		}
	case *Tx_Reserve:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Reserve); err != nil {
			return err
		}
	case *Tx_Release:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Release); err != nil {
			return err
		}
	case *Tx_ServicePayment:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ServicePayment); err != nil {
			return err
		}
	case *Tx_Slash:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Slash); err != nil {
			return err
		}
	case *Tx_SplitContract:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SplitContract); err != nil {
			return err
		}
	case *Tx_UpdateValidators:
		_ = b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdateValidators); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Tx.Tx has unexpected type %T", x)
	}
	return nil
}

func _Tx_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Tx)
	switch tag {
	case 1: // tx.coinbase
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CoinbaseTx)
		err := b.DecodeMessage(msg)
		m.Tx = &Tx_Coinbase{msg}
		return true, err
	case 2: // tx.send
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SendTx)
		err := b.DecodeMessage(msg)
		m.Tx = &Tx_Send{msg}
		return true, err
	case 3: // tx.reserve
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ReserveFundTx)
		err := b.DecodeMessage(msg)
		m.Tx = &Tx_Reserve{msg}
		return true, err
	case 4: // tx.release
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ReleaseFundTx)
		err := b.DecodeMessage(msg)
		m.Tx = &Tx_Release{msg}
		return true, err
	case 5: // tx.service_payment
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ServicePaymentTx)
		err := b.DecodeMessage(msg)
		m.Tx = &Tx_ServicePayment{msg}
		return true, err
	case 6: // tx.slash
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SlashTx)
		err := b.DecodeMessage(msg)
		m.Tx = &Tx_Slash{msg}
		return true, err
	case 7: // tx.split_contract
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SplitContractTx)
		err := b.DecodeMessage(msg)
		m.Tx = &Tx_SplitContract{msg}
		return true, err
	case 8: // tx.update_validators
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdateValidatorsTx)
		err := b.DecodeMessage(msg)
		m.Tx = &Tx_UpdateValidators{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Tx_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Tx)
	// tx
	switch x := m.Tx.(type) {
	case *Tx_Coinbase:
		s := proto.Size(x.Coinbase)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_Send:
		s := proto.Size(x.Send)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_Reserve:
		s := proto.Size(x.Reserve)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_Release:
		s := proto.Size(x.Release)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_ServicePayment:
		s := proto.Size(x.ServicePayment)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_Slash:
		s := proto.Size(x.Slash)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_SplitContract:
		s := proto.Size(x.SplitContract)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_UpdateValidators:
		s := proto.Size(x.UpdateValidators)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CoinbaseTx struct {
	Proposer    *TxInput    `protobuf:"bytes,1,opt,name=proposer" json:"proposer,omitempty"`
	Outputs     []*TxOutput `protobuf:"bytes,2,rep,name=outputs" json:"outputs,omitempty"`
	BlockHeight int64       `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *CoinbaseTx) Reset()                    { *m = CoinbaseTx{} }
func (m *CoinbaseTx) String() string            { return proto.CompactTextString(m) }
func (*CoinbaseTx) ProtoMessage()               {}
func (*CoinbaseTx) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{12} }

func (m *CoinbaseTx) GetProposer() *TxInput {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *CoinbaseTx) GetOutputs() []*TxOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *CoinbaseTx) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type SendTx struct {
	Gas     int64       `protobuf:"varint,1,opt,name=gas,proto3" json:"gas,omitempty"`
	Fee     *Coin       `protobuf:"bytes,2,opt,name=fee" json:"fee,omitempty"`
	Inputs  []*TxInput  `protobuf:"bytes,3,rep,name=inputs" json:"inputs,omitempty"`
	Outputs []*TxOutput `protobuf:"bytes,4,rep,name=outputs" json:"outputs,omitempty"`
}

func (m *SendTx) Reset()                    { *m = SendTx{} }
func (m *SendTx) String() string            { return proto.CompactTextString(m) }
func (*SendTx) ProtoMessage()               {}
func (*SendTx) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{13} }

func (m *SendTx) GetGas() int64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func (m *SendTx) GetFee() *Coin {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *SendTx) GetInputs() []*TxInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *SendTx) GetOutputs() []*TxOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type ReserveFundTx struct {
	Gas         int64    `protobuf:"varint,1,opt,name=gas,proto3" json:"gas,omitempty"`
	Fee         *Coin    `protobuf:"bytes,2,opt,name=fee" json:"fee,omitempty"`
	Source      *TxInput `protobuf:"bytes,3,opt,name=source" json:"source,omitempty"`
	Collateral  []*Coin  `protobuf:"bytes,4,rep,name=collateral" json:"collateral,omitempty"`
	ResourceIds [][]byte `protobuf:"bytes,5,rep,name=resource_ids,json=resourceIds" json:"resource_ids,omitempty"`
	Duration    int64    `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (m *ReserveFundTx) Reset()                    { *m = ReserveFundTx{} }
func (m *ReserveFundTx) String() string            { return proto.CompactTextString(m) }
func (*ReserveFundTx) ProtoMessage()               {}
func (*ReserveFundTx) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{14} }

func (m *ReserveFundTx) GetGas() int64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func (m *ReserveFundTx) GetFee() *Coin {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *ReserveFundTx) GetSource() *TxInput {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *ReserveFundTx) GetCollateral() []*Coin {
	if m != nil {
		return m.Collateral
	}
	return nil
}

func (m *ReserveFundTx) GetResourceIds() [][]byte {
	if m != nil {
		return m.ResourceIds
	}
	return nil
}

func (m *ReserveFundTx) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type ReleaseFundTx struct {
	Gas             int64    `protobuf:"varint,1,opt,name=gas,proto3" json:"gas,omitempty"`
	Fee             *Coin    `protobuf:"bytes,2,opt,name=fee" json:"fee,omitempty"`
	Source          *TxInput `protobuf:"bytes,3,opt,name=source" json:"source,omitempty"`
	ReserveSequence int64    `protobuf:"varint,4,opt,name=reserve_sequence,json=reserveSequence,proto3" json:"reserve_sequence,omitempty"`
}

func (m *ReleaseFundTx) Reset()                    { *m = ReleaseFundTx{} }
func (m *ReleaseFundTx) String() string            { return proto.CompactTextString(m) }
func (*ReleaseFundTx) ProtoMessage()               {}
func (*ReleaseFundTx) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{15} }

func (m *ReleaseFundTx) GetGas() int64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func (m *ReleaseFundTx) GetFee() *Coin {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *ReleaseFundTx) GetSource() *TxInput {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *ReleaseFundTx) GetReserveSequence() int64 {
	if m != nil {
		return m.ReserveSequence
	}
	return 0
}

type ServicePaymentTx struct {
	Gas             int64    `protobuf:"varint,1,opt,name=gas,proto3" json:"gas,omitempty"`
	Fee             *Coin    `protobuf:"bytes,2,opt,name=fee" json:"fee,omitempty"`
	Source          *TxInput `protobuf:"bytes,3,opt,name=source" json:"source,omitempty"`
	Target          *TxInput `protobuf:"bytes,4,opt,name=target" json:"target,omitempty"`
	PaymentSequence int64    `protobuf:"varint,5,opt,name=PaymentSequence,proto3" json:"PaymentSequence,omitempty"`
	ReserveSequence int64    `protobuf:"varint,6,opt,name=ReserveSequence,proto3" json:"ReserveSequence,omitempty"`
	ResourceId      []byte   `protobuf:"bytes,7,opt,name=ResourceId,proto3" json:"ResourceId,omitempty"`
}

func (m *ServicePaymentTx) Reset()                    { *m = ServicePaymentTx{} }
func (m *ServicePaymentTx) String() string            { return proto.CompactTextString(m) }
func (*ServicePaymentTx) ProtoMessage()               {}
func (*ServicePaymentTx) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{16} }

func (m *ServicePaymentTx) GetGas() int64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func (m *ServicePaymentTx) GetFee() *Coin {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *ServicePaymentTx) GetSource() *TxInput {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *ServicePaymentTx) GetTarget() *TxInput {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *ServicePaymentTx) GetPaymentSequence() int64 {
	if m != nil {
		return m.PaymentSequence
	}
	return 0
}

func (m *ServicePaymentTx) GetReserveSequence() int64 {
	if m != nil {
		return m.ReserveSequence
	}
	return 0
}

func (m *ServicePaymentTx) GetResourceId() []byte {
	if m != nil {
		return m.ResourceId
	}
	return nil
}

type SlashTx struct {
	Proposer        *TxInput `protobuf:"bytes,1,opt,name=proposer" json:"proposer,omitempty"`
	SlashedAddress  []byte   `protobuf:"bytes,2,opt,name=slashed_address,json=slashedAddress,proto3" json:"slashed_address,omitempty"`
	ReserveSequence int64    `protobuf:"varint,3,opt,name=reserve_sequence,json=reserveSequence,proto3" json:"reserve_sequence,omitempty"`
	SlashProof      []byte   `protobuf:"bytes,4,opt,name=slash_proof,json=slashProof,proto3" json:"slash_proof,omitempty"`
}

func (m *SlashTx) Reset()                    { *m = SlashTx{} }
func (m *SlashTx) String() string            { return proto.CompactTextString(m) }
func (*SlashTx) ProtoMessage()               {}
func (*SlashTx) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{17} }

func (m *SlashTx) GetProposer() *TxInput {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *SlashTx) GetSlashedAddress() []byte {
	if m != nil {
		return m.SlashedAddress
	}
	return nil
}

func (m *SlashTx) GetReserveSequence() int64 {
	if m != nil {
		return m.ReserveSequence
	}
	return 0
}

func (m *SlashTx) GetSlashProof() []byte {
	if m != nil {
		return m.SlashProof
	}
	return nil
}

type Split struct {
	Address    []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Percentage int64  `protobuf:"varint,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (m *Split) Reset()                    { *m = Split{} }
func (m *Split) String() string            { return proto.CompactTextString(m) }
func (*Split) ProtoMessage()               {}
func (*Split) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{18} }

func (m *Split) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Split) GetPercentage() int64 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

type SplitContract struct {
	InitiatorAddress []byte   `protobuf:"bytes,1,opt,name=initiator_address,json=initiatorAddress,proto3" json:"initiator_address,omitempty"`
	ResourceId       []byte   `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Splits           []*Split `protobuf:"bytes,3,rep,name=splits" json:"splits,omitempty"`
	EndBlockHeight   int64    `protobuf:"varint,4,opt,name=end_block_height,json=endBlockHeight,proto3" json:"end_block_height,omitempty"`
}

func (m *SplitContract) Reset()                    { *m = SplitContract{} }
func (m *SplitContract) String() string            { return proto.CompactTextString(m) }
func (*SplitContract) ProtoMessage()               {}
func (*SplitContract) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{19} }

func (m *SplitContract) GetInitiatorAddress() []byte {
	if m != nil {
		return m.InitiatorAddress
	}
	return nil
}

func (m *SplitContract) GetResourceId() []byte {
	if m != nil {
		return m.ResourceId
	}
	return nil
}

func (m *SplitContract) GetSplits() []*Split {
	if m != nil {
		return m.Splits
	}
	return nil
}

func (m *SplitContract) GetEndBlockHeight() int64 {
	if m != nil {
		return m.EndBlockHeight
	}
	return 0
}

type SplitContractTx struct {
	Gas        int64    `protobuf:"varint,1,opt,name=gas,proto3" json:"gas,omitempty"`
	Fee        *Coin    `protobuf:"bytes,2,opt,name=fee" json:"fee,omitempty"`
	ResourceId []byte   `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Initiator  *TxInput `protobuf:"bytes,4,opt,name=initiator" json:"initiator,omitempty"`
	Splits     []*Split `protobuf:"bytes,5,rep,name=splits" json:"splits,omitempty"`
	Duration   int64    `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (m *SplitContractTx) Reset()                    { *m = SplitContractTx{} }
func (m *SplitContractTx) String() string            { return proto.CompactTextString(m) }
func (*SplitContractTx) ProtoMessage()               {}
func (*SplitContractTx) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{20} }

func (m *SplitContractTx) GetGas() int64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func (m *SplitContractTx) GetFee() *Coin {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *SplitContractTx) GetResourceId() []byte {
	if m != nil {
		return m.ResourceId
	}
	return nil
}

func (m *SplitContractTx) GetInitiator() *TxInput {
	if m != nil {
		return m.Initiator
	}
	return nil
}

func (m *SplitContractTx) GetSplits() []*Split {
	if m != nil {
		return m.Splits
	}
	return nil
}

func (m *SplitContractTx) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type UpdateValidatorsTx struct {
	Gas        int64        `protobuf:"varint,1,opt,name=gas,proto3" json:"gas,omitempty"`
	Fee        *Coin        `protobuf:"bytes,2,opt,name=fee" json:"fee,omitempty"`
	Proposer   *TxInput     `protobuf:"bytes,3,opt,name=proposer" json:"proposer,omitempty"`
	Validators []*Validator `protobuf:"bytes,4,rep,name=validators" json:"validators,omitempty"`
}

func (m *UpdateValidatorsTx) Reset()                    { *m = UpdateValidatorsTx{} }
func (m *UpdateValidatorsTx) String() string            { return proto.CompactTextString(m) }
func (*UpdateValidatorsTx) ProtoMessage()               {}
func (*UpdateValidatorsTx) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{21} }

func (m *UpdateValidatorsTx) GetGas() int64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func (m *UpdateValidatorsTx) GetFee() *Coin {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *UpdateValidatorsTx) GetProposer() *TxInput {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *UpdateValidatorsTx) GetValidators() []*Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func init() {
	proto.RegisterType((*PublicKey)(nil), "serialization.PublicKey")
	proto.RegisterType((*PrivateKey)(nil), "serialization.PrivateKey")
	proto.RegisterType((*Signature)(nil), "serialization.Signature")
	proto.RegisterType((*Validator)(nil), "serialization.Validator")
	proto.RegisterType((*OverspendingProof)(nil), "serialization.OverspendingProof")
	proto.RegisterType((*Coin)(nil), "serialization.Coin")
	proto.RegisterType((*ReservedFund)(nil), "serialization.ReservedFund")
	proto.RegisterType((*TransferRecord)(nil), "serialization.TransferRecord")
	proto.RegisterType((*Account)(nil), "serialization.Account")
	proto.RegisterType((*TxInput)(nil), "serialization.TxInput")
	proto.RegisterType((*TxOutput)(nil), "serialization.TxOutput")
	proto.RegisterType((*Tx)(nil), "serialization.Tx")
	proto.RegisterType((*CoinbaseTx)(nil), "serialization.CoinbaseTx")
	proto.RegisterType((*SendTx)(nil), "serialization.SendTx")
	proto.RegisterType((*ReserveFundTx)(nil), "serialization.ReserveFundTx")
	proto.RegisterType((*ReleaseFundTx)(nil), "serialization.ReleaseFundTx")
	proto.RegisterType((*ServicePaymentTx)(nil), "serialization.ServicePaymentTx")
	proto.RegisterType((*SlashTx)(nil), "serialization.SlashTx")
	proto.RegisterType((*Split)(nil), "serialization.Split")
	proto.RegisterType((*SplitContract)(nil), "serialization.SplitContract")
	proto.RegisterType((*SplitContractTx)(nil), "serialization.SplitContractTx")
	proto.RegisterType((*UpdateValidatorsTx)(nil), "serialization.UpdateValidatorsTx")
}

func init() { proto.RegisterFile("ledger/types/serialization/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 1237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x4f, 0x73, 0xdb, 0x44,
	0x14, 0xaf, 0x2c, 0xff, 0x7d, 0x76, 0x6c, 0x77, 0x29, 0x41, 0x0d, 0x90, 0xa4, 0x9a, 0x01, 0xd2,
	0x29, 0x75, 0xda, 0xb4, 0x53, 0xca, 0x31, 0xe9, 0x4c, 0x49, 0xdb, 0x43, 0x32, 0x8a, 0xe1, 0xc0,
	0x45, 0xb3, 0x96, 0x36, 0x8e, 0x26, 0x8a, 0x24, 0x76, 0x57, 0x19, 0x9b, 0x0f, 0xc0, 0x81, 0x0f,
	0xc0, 0x95, 0x03, 0x7c, 0x02, 0x66, 0x38, 0xf1, 0x35, 0xf8, 0x16, 0xc0, 0x27, 0xe0, 0xc2, 0x68,
	0x77, 0xa5, 0x48, 0xb2, 0x1c, 0x3c, 0x81, 0xe9, 0xcd, 0xfb, 0xf6, 0xf7, 0x56, 0xef, 0xb7, 0xef,
	0xf7, 0xde, 0x5b, 0xc3, 0xc7, 0x3e, 0x71, 0xa7, 0x84, 0xee, 0xf2, 0x79, 0x44, 0xd8, 0x2e, 0x23,
	0xd4, 0xc3, 0xbe, 0xf7, 0x2d, 0xe6, 0x5e, 0x18, 0x48, 0xdb, 0x28, 0xa2, 0x21, 0x0f, 0xd1, 0x5a,
	0x61, 0x6b, 0xe3, 0xe1, 0xd4, 0xe3, 0x67, 0xf1, 0x64, 0xe4, 0x84, 0x17, 0xbb, 0xd3, 0x70, 0x1a,
	0xee, 0x0a, 0xd4, 0x24, 0x3e, 0x15, 0x2b, 0xb1, 0x10, 0xbf, 0xa4, 0xb7, 0xb9, 0x05, 0x9d, 0xe3,
	0x78, 0xe2, 0x7b, 0xce, 0x1b, 0x32, 0x47, 0x08, 0xea, 0x2e, 0xe6, 0xd8, 0xa8, 0x6d, 0x6b, 0x3b,
	0x3d, 0x4b, 0xfc, 0x36, 0xb7, 0x01, 0x8e, 0xa9, 0x77, 0x89, 0x39, 0x59, 0x86, 0xd8, 0x82, 0xce,
	0x89, 0x37, 0x0d, 0x30, 0x8f, 0x29, 0xa9, 0x04, 0x3c, 0x86, 0xce, 0x57, 0xd8, 0xf7, 0x5c, 0xcc,
	0x43, 0x8a, 0xfa, 0x50, 0xf3, 0x5c, 0x43, 0x13, 0xdb, 0x35, 0xcf, 0x45, 0x77, 0xa0, 0xc1, 0x38,
	0x3e, 0x27, 0xc2, 0x43, 0xb7, 0xe4, 0xc2, 0xfc, 0x5e, 0x83, 0xdb, 0x47, 0x97, 0x84, 0xb2, 0x88,
	0x04, 0xae, 0x17, 0x4c, 0x8f, 0x69, 0x18, 0x9e, 0xa2, 0xfb, 0x30, 0xa4, 0x84, 0x11, 0x7a, 0x49,
	0x6c, 0x46, 0xbe, 0x89, 0x49, 0xe0, 0x10, 0x71, 0x92, 0x6e, 0x0d, 0x94, 0xfd, 0x44, 0x99, 0xd1,
	0x6b, 0x18, 0x26, 0x06, 0xcf, 0x21, 0x76, 0x84, 0xe7, 0x17, 0x24, 0xe0, 0xcc, 0xa8, 0x6d, 0xeb,
	0x3b, 0xdd, 0xbd, 0xad, 0x51, 0xe1, 0xc2, 0x46, 0x27, 0x12, 0x76, 0x2c, 0x51, 0xe3, 0x99, 0x35,
	0x60, 0x05, 0x0b, 0x33, 0x9f, 0x42, 0xfd, 0x45, 0xe8, 0x05, 0x49, 0xa8, 0x2e, 0x09, 0xc2, 0x0b,
	0xf1, 0xcd, 0x8e, 0x25, 0x17, 0x68, 0x1d, 0x9a, 0xf8, 0x22, 0x8c, 0x03, 0xae, 0x18, 0xa8, 0x95,
	0xf9, 0x57, 0x0d, 0x7a, 0x96, 0x8c, 0xca, 0x7d, 0x19, 0x07, 0x2e, 0x7a, 0x02, 0xe0, 0x84, 0xbe,
	0x8f, 0x39, 0xa1, 0xd8, 0x37, 0x34, 0x11, 0xcc, 0x3b, 0xa5, 0x60, 0x92, 0xef, 0x58, 0x39, 0x18,
	0x7a, 0x06, 0x3d, 0x2f, 0xf0, 0xb8, 0x87, 0x7d, 0xfb, 0x34, 0x0e, 0x5c, 0xc5, 0xa1, 0xd2, 0xad,
	0xab, 0x80, 0xe2, 0x63, 0x8f, 0xa0, 0x13, 0x33, 0xe2, 0x4a, 0x27, 0x7d, 0xb9, 0x53, 0x3b, 0x41,
	0x09, 0x8f, 0x7b, 0xd0, 0xa3, 0x84, 0x85, 0x31, 0x75, 0x88, 0xed, 0xb9, 0xcc, 0xa8, 0x6f, 0xeb,
	0x3b, 0x3d, 0xab, 0x9b, 0xda, 0x5e, 0xb9, 0x0c, 0xed, 0xc0, 0x90, 0x04, 0xae, 0x3d, 0xf1, 0x43,
	0xe7, 0xdc, 0x3e, 0x23, 0xde, 0xf4, 0x8c, 0x1b, 0x0d, 0x41, 0xba, 0x4f, 0x02, 0xf7, 0x20, 0x31,
	0x1f, 0x0a, 0x6b, 0x65, 0xa6, 0x9a, 0xd5, 0x99, 0x7a, 0x09, 0x03, 0x4e, 0x71, 0xc0, 0x4e, 0x09,
	0xb5, 0x29, 0x71, 0x42, 0xea, 0x1a, 0x2d, 0x11, 0xef, 0x87, 0xa5, 0x78, 0xc7, 0x0a, 0x65, 0x09,
	0x90, 0xd5, 0xe7, 0x85, 0xb5, 0xf9, 0x35, 0xf4, 0x8b, 0x08, 0x74, 0x08, 0x83, 0x92, 0x06, 0x44,
	0xe6, 0x56, 0x90, 0x40, 0xbf, 0x28, 0x01, 0xf3, 0xbb, 0x1a, 0xb4, 0xf6, 0x1d, 0x27, 0xc9, 0x2b,
	0xda, 0x80, 0x76, 0x49, 0x7c, 0xd9, 0x1a, 0x3d, 0x84, 0xd6, 0x04, 0xfb, 0x38, 0xd9, 0xba, 0x26,
	0x51, 0x29, 0x06, 0x1d, 0x40, 0x5f, 0xdd, 0x86, 0x4c, 0x14, 0x53, 0x99, 0x7a, 0xbf, 0xe4, 0x95,
	0x97, 0x91, 0xb5, 0x46, 0x73, 0x2b, 0x86, 0x3e, 0x87, 0xbb, 0x3e, 0x66, 0xdc, 0x8e, 0x23, 0x17,
	0x73, 0x52, 0x4a, 0x4e, 0x5d, 0xc4, 0xb7, 0x9e, 0x00, 0xbe, 0x94, 0xfb, 0xf9, 0x24, 0x3d, 0x86,
	0x56, 0x14, 0x4f, 0xec, 0x73, 0x32, 0x17, 0x59, 0xec, 0xee, 0x19, 0xa5, 0xef, 0x66, 0x9d, 0xc1,
	0x6a, 0x46, 0xf1, 0xe4, 0x0d, 0x99, 0x9b, 0xbf, 0x6b, 0xd0, 0x1a, 0xcf, 0x5e, 0x05, 0x51, 0xcc,
	0x91, 0x01, 0x2d, 0xec, 0xba, 0x94, 0x30, 0xa6, 0xca, 0x39, 0x5d, 0xa2, 0xfb, 0xd0, 0x70, 0x42,
	0x2f, 0x60, 0xd7, 0x5d, 0x82, 0x44, 0x14, 0x6e, 0x53, 0x2f, 0xdd, 0xe6, 0x33, 0xe8, 0xb0, 0xb4,
	0xb1, 0x08, 0x2a, 0x8b, 0x11, 0x66, 0x8d, 0xc7, 0xba, 0x82, 0xa2, 0x47, 0x90, 0x84, 0xbb, 0x2a,
	0xad, 0x73, 0x32, 0x37, 0x8f, 0xa0, 0x3d, 0x9e, 0x1d, 0xc5, 0xfc, 0xff, 0xa2, 0x65, 0xfe, 0xad,
	0x43, 0x6d, 0x3c, 0x43, 0x9f, 0x41, 0x3b, 0x59, 0x4f, 0x30, 0x23, 0x4a, 0x7a, 0x77, 0x2b, 0x9c,
	0x92, 0xed, 0xf1, 0xec, 0xf0, 0x96, 0x95, 0x81, 0xd1, 0x03, 0xa8, 0x33, 0x22, 0xca, 0x3d, 0x71,
	0x7a, 0x77, 0x41, 0xaf, 0x81, 0x2b, 0x1c, 0x04, 0x08, 0x3d, 0x87, 0x96, 0xd2, 0x84, 0xb8, 0xc2,
	0xee, 0xde, 0x07, 0xd5, 0xfa, 0x49, 0x04, 0x23, 0xdc, 0x52, 0xb8, 0xf4, 0xf4, 0x49, 0x12, 0x5e,
	0x7d, 0x89, 0xa7, 0xd8, 0xcd, 0x7b, 0x0a, 0x03, 0x7a, 0xbd, 0x58, 0x5b, 0x8d, 0x95, 0x6a, 0xeb,
	0xf0, 0x56, 0xb9, 0xba, 0xd0, 0x08, 0x1a, 0xcc, 0xc7, 0xec, 0x4c, 0x74, 0x88, 0xee, 0xde, 0x7a,
	0xf9, 0x84, 0x64, 0x4f, 0x38, 0x4a, 0x18, 0xfa, 0x02, 0xfa, 0x2c, 0xf2, 0x3d, 0x6e, 0x3b, 0x61,
	0xc0, 0x29, 0x76, 0xb8, 0xd1, 0x12, 0x8e, 0x9b, 0x65, 0xc7, 0x04, 0xf4, 0x42, 0x61, 0xc4, 0x01,
	0x6b, 0x2c, 0x6f, 0x42, 0xc7, 0x70, 0x5b, 0x96, 0x8d, 0x7d, 0x99, 0xce, 0x27, 0x66, 0xb4, 0xc5,
	0x59, 0xf7, 0x4a, 0x67, 0xc9, 0xf2, 0xc9, 0xc6, 0x18, 0x13, 0xc7, 0x0d, 0xe3, 0x92, 0xf5, 0xa0,
	0x0e, 0x35, 0x3e, 0x33, 0x7f, 0xd0, 0x00, 0xae, 0x12, 0x8b, 0xf6, 0xa0, 0x1d, 0xd1, 0x30, 0x0a,
	0x19, 0xa1, 0x4a, 0x05, 0x65, 0x8a, 0xaa, 0xa4, 0xac, 0x0c, 0x97, 0xd4, 0x66, 0x28, 0xf4, 0x98,
	0xaa, 0xed, 0xbd, 0x05, 0x17, 0xa9, 0x57, 0x2b, 0xc5, 0x25, 0x0d, 0xbc, 0x50, 0xfc, 0xb2, 0x9c,
	0xba, 0x93, 0xab, 0x8a, 0x37, 0x7f, 0xd6, 0xa0, 0x29, 0xc5, 0x83, 0x86, 0xa0, 0x4f, 0x31, 0x53,
	0x1d, 0x2c, 0xf9, 0x89, 0x3e, 0x02, 0xfd, 0x94, 0x10, 0x25, 0xb9, 0x4a, 0x71, 0x27, 0xfb, 0x68,
	0x04, 0x4d, 0x2f, 0x10, 0x81, 0xc9, 0x66, 0xb5, 0x8c, 0x8b, 0x42, 0xe5, 0x99, 0xd4, 0x57, 0x63,
	0x62, 0xfe, 0xa9, 0xc1, 0x5a, 0x41, 0xb3, 0xff, 0x29, 0x5a, 0x39, 0xbf, 0x54, 0x69, 0x2c, 0x8d,
	0x56, 0xa2, 0x4a, 0x43, 0xba, 0xbe, 0xda, 0x90, 0x2e, 0x8f, 0xce, 0xc6, 0xe2, 0xe8, 0xdc, 0x80,
	0xb6, 0x1b, 0x53, 0xe1, 0xaf, 0x06, 0x61, 0xb6, 0x36, 0x7f, 0x12, 0x74, 0x73, 0x85, 0xf6, 0xf6,
	0xe8, 0x56, 0xcd, 0xe9, 0x7a, 0xe5, 0x9c, 0x36, 0x7f, 0xac, 0xc1, 0xb0, 0x5c, 0xcc, 0x6f, 0x2f,
	0xd0, 0x11, 0x34, 0x39, 0xa6, 0x53, 0xc2, 0x55, 0xa3, 0x5a, 0x8a, 0x97, 0x28, 0xb4, 0x03, 0x03,
	0x15, 0x65, 0x4a, 0x40, 0xbd, 0x54, 0xca, 0xe6, 0x04, 0x69, 0x15, 0xa9, 0xa6, 0x2f, 0x95, 0x92,
	0x19, 0x6d, 0x02, 0x58, 0x59, 0x4a, 0x45, 0xcf, 0xe9, 0x59, 0x39, 0x8b, 0xf9, 0x8b, 0x06, 0x2d,
	0xd5, 0xac, 0x6e, 0x54, 0xf3, 0x9f, 0xc0, 0x40, 0x34, 0x38, 0xe2, 0xda, 0xe9, 0x04, 0x92, 0xcf,
	0xe8, 0xbe, 0x32, 0xef, 0x67, 0x83, 0x68, 0x31, 0x6b, 0x7a, 0xf5, 0xeb, 0x6a, 0x0b, 0xba, 0xc2,
	0xd9, 0x8e, 0x92, 0x17, 0xb4, 0xb8, 0xbc, 0x9e, 0x05, 0xc2, 0x24, 0xde, 0xd4, 0xe6, 0x3e, 0x34,
	0x44, 0x9f, 0xbc, 0x66, 0xee, 0x6d, 0x02, 0x44, 0x84, 0x3a, 0x24, 0xe0, 0x78, 0x9a, 0xbe, 0xd3,
	0x73, 0x16, 0xf3, 0x57, 0x0d, 0xd6, 0x0a, 0xbd, 0x16, 0x3d, 0x80, 0xdb, 0xf2, 0x31, 0xca, 0x43,
	0x6a, 0x17, 0x4f, 0x1d, 0x66, 0x1b, 0x29, 0x9b, 0x2d, 0xe8, 0xe6, 0xaa, 0x47, 0x51, 0x86, 0xab,
	0xe2, 0x41, 0x9f, 0x42, 0x53, 0xf4, 0xed, 0xb4, 0xe3, 0xdc, 0xa9, 0xea, 0xf3, 0x96, 0xc2, 0x54,
	0x3e, 0x52, 0xeb, 0x55, 0x8f, 0x54, 0xf3, 0x0f, 0x0d, 0x06, 0xa5, 0x19, 0x71, 0x73, 0x41, 0x97,
	0x58, 0xe8, 0x0b, 0x2c, 0x9e, 0x42, 0x27, 0xa3, 0xfe, 0x2f, 0x22, 0xbe, 0x02, 0xe6, 0xb8, 0x37,
	0x56, 0xe0, 0x7e, 0x5d, 0x97, 0xf9, 0x4d, 0x03, 0xb4, 0x38, 0xc5, 0x6e, 0x4e, 0x38, 0xaf, 0x70,
	0x7d, 0x45, 0x85, 0x3f, 0x07, 0xc8, 0x4d, 0x5a, 0xd9, 0x5d, 0xcb, 0xaf, 0xb3, 0x2c, 0x3a, 0x2b,
	0x87, 0x9d, 0x34, 0xc5, 0xdf, 0xd5, 0x27, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x61, 0x66, 0xa2,
	0xeb, 0x16, 0x0f, 0x00, 0x00,
}
