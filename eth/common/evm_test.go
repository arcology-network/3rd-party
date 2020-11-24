package common

import (
	"fmt"
	"testing"
)

func Test_Uint32(t *testing.T) {
	code := AlignToEvmForInt(32)
	fmt.Printf("code=%v\n", code)
}

func Test_strings(t *testing.T) {
	code := AlignToEvmForString("sdfdsfsd2123ffffffffssswwwwwww999999")
	fmt.Printf("code=%v\n", code)
}
