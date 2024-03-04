package lowlevel

//

type ByteOrder	int8


func Endianness() ByteOrder {
    return hostByteOrder
}
