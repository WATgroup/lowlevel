package lowlevel

//

type ByteOrder int

func Endianness() ByteOrder {
	return hostByteOrder
}

func (e ByteOrder) String() string {

	switch e {
	case +1:
		return "Big Endian"
	case -1:
		return "litte Endian"
	default:
		return `unknown`
	}
}
