package packet

import (
	"fmt"
	"testing"

	"github.com/cpusoft/goutil/jsonutil"
)

func TestParseBytesToTxtModel(t *testing.T) {
	bytess := []byte{0x76, 0x3d, 0x73, 0x70, 0x66, 0x31, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x3a, 0x73, 0x70, 0x66, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x71, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x20, 0x2d, 0x61, 0x6c, 0x6c}
	a, newOffsetFromStart, err := ParseBytesToTxtModel(bytess, 999)
	fmt.Println(jsonutil.MarshalJson(a), newOffsetFromStart, err)
}
