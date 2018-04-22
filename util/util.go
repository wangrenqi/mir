package util

import "encoding/binary"

func IndexToBytes(index int) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, uint16(index))
	return bytes
}
