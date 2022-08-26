package utils

func UInt8(data []byte) uint8 {
	return uint8(data[0])
}

func UInt16(data []byte) uint16 {
	return uint16(data[0])<<8 +
		uint16(data[1])
}

func UInt32(data []byte) uint32 {
	return uint32(data[0])<<32 +
		uint32(data[1])<<16 +
		uint32(data[2])<<8 +
		uint32(data[3])
}
