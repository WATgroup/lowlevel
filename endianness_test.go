package lowlevel

import (
	"fmt"
	"testing"
)

func TestEndianness(t *testing.T) {

	fmt.Println("Endianness:", Endianness())

}
