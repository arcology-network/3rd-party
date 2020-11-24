package crypto

import (
	"github.com/HPISTechnologies/3rd-party/eth/crypto/v2/secp256k1"
)

// Ecrecover returns the uncompressed public key that created the given signature.
func Ecrecover(hash, sig []byte) ([]byte, error) {
	return secp256k1.RecoverPubkey(hash, sig)
}
