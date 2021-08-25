package fuzz_binary

import (
	"github.com/arcology/3rd-party/tm/go-amino"
	"github.com/arcology/3rd-party/tm/go-amino/tests"
)

//-------------------------------------
// Non-interface go-fuzz tests
// See https://github.com/dvyukov/go-fuzz
// (Test that deserialize never panics)

func Fuzz(data []byte) int {
	cdc := amino.NewCodec()
	cst := tests.ComplexSt{}
	err := cdc.UnmarshalBinary(data, &cst)
	if err != nil {
		return 0
	}
	return 1
}
