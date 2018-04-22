package util

import (
	"encoding/binary"
	"log"
)

func IndexToBytes(index int) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, uint16(index))
	return bytes
}

// 根据传入的索引 返回读完string后所在bytes 的下一个索引 及string
func ReadString(bytes []byte, index int) (int, string) {
	//[0, 1, 2, 5, 9, 10, 11, 12, 50, 23, 77, 99]
	if len(bytes) == 0 {
		return -1, ""
	}
	strLen := int(bytes[index])
	lastIndex := index + strLen + 1
	msg := string(bytes[index+1 : lastIndex])
	log.Println("ReadString:", bytes, "to string:", msg)
	return lastIndex, msg
}
