package types

import (
	"github.com/cosmos/cosmos-sdk/v42/codec"
	"github.com/cosmos/cosmos-sdk/v42/codec/types"
)

const (
	// MsgInterfaceProtoName defines the protobuf name of the cosmos Msg interface
	MsgInterfaceProtoName = "cosmos.base.v42beta1.Msg"
)

// RegisterLegacyAminoCodec registers the sdk message type.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*Msg)(nil), nil)
	cdc.RegisterInterface((*Tx)(nil), nil)
}

// RegisterInterfaces registers the sdk message type.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(MsgInterfaceProtoName, (*Msg)(nil))
}
