//go:build amd64 || ppc64el || arm64 || loong64

package lowlevel

// This is a littleEndian architecture
var hostByteOrder ByteOrder = -1
