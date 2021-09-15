package types

import (
	"math/big"
	"testing"

	ethCommon "github.com/arcology-network/3rd-party/eth/common"
	ethRlp "github.com/arcology-network/3rd-party/eth/rlp"
)

func BenchmarkTransactionAsMessage(b *testing.B) {
	rawTx := ethCommon.FromHex("f86a8086d55698372431831e848094f0109fc8df283027b6285cc889f5aa624eac1f55843b9aca008025a009ebb6ca057a0535d6186462bc0b465b561c94a295bdb0621fc19208ab149a9ca0440ffd775ce91a833ab410777204d5341a6f9fa91216a6f3ee2c051fea6a0428")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100000; j++ {
			var tx Transaction
			if err := ethRlp.DecodeBytes(rawTx, &tx); err != nil {
				b.Error(err)
			}
			ethCommon.RlpHash(&tx)
			if _, err := tx.AsMessage(NewEIP155Signer(new(big.Int).SetInt64(1))); err != nil {
				b.Error(err)
			}
		}
	}
}
