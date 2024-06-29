//go:build ppc64 || mips

package lowlevel

// This is a BigEndian architecture
var hostByteOrder ByteOrder = +1
