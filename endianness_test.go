package lowlevel

import (
    "testing"
    "fmt"
)


func TestEndianness(t *testing.T) {

    fmt.Println("Endianness:", Endianness())

}
