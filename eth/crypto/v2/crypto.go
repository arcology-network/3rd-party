package crypto

import (
	"hash"

	"golang.org/x/crypto/sha3"
)

type KeccakState interface {
	hash.Hash
	Read([]byte) (int, error)
}

func Keccak256(data ...[]byte) []byte {
	b := make([]byte, 32)
	d := sha3.NewLegacyKeccak256().(KeccakState)
	for _, b := range data {
		d.Write(b)
	}
	d.Read(b)
	return b
}
