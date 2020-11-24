package common

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"testing"
	"time"
)

func PrepareData(max int) []Hash {
	hashes := make([]Hash, max)
	for i := range hashes {
		buffer := make([]byte, 4)
		binary.BigEndian.PutUint32(buffer, uint32(i))
		bytes := sha256.Sum256(buffer)
		hashes[i] = (BytesToHash(bytes[0:32]))
	}
	return hashes
}

func TestChecksum(t *testing.T) {
	hashes := PrepareData(500000)
	t0 := time.Now()
	Hashes(hashes).Checksum()
	fmt.Println("Checksum():", time.Now().Sub(t0))

	bytes := Hashes(hashes).Flatten()
	t0 = time.Now()
	sha256.Sum256(bytes)
	fmt.Println("Msha256.Sum256():", time.Now().Sub(t0))

	t0 = time.Now()
	md5.Sum(bytes)
	fmt.Println("md5.Sum():", time.Now().Sub(t0))
}

func TestEthHashFlatten(t *testing.T) {
	hashes := PrepareData(500000)
	t0 := time.Now()
	Hashes(hashes).Flatten()
	fmt.Println("Flatten():", time.Now().Sub(t0))
}
