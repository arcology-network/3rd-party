package keys

import (
	amino "github.com/arcology-network/3rd-party/tm/go-amino"
	crypto "github.com/arcology-network/3rd-party/tm/go-crypto"
)

var cdc = amino.NewCodec()

func init() {
	crypto.RegisterAmino(cdc)
}
