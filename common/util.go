package common

import (
	"encoding/binary"
	"log"
)

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

func StringToBytes(s string) []byte {
	// 根据传入的string 返回字节数组，索引为0（即数组开头）的值为数组的长度
	stringBytes := []byte(s)
	stringBytesLenBytes := []byte{byte(len(stringBytes))}
	return append(stringBytesLenBytes, stringBytes...)
}

func Uint16ToBytes(num uint16) []byte {
	res := make([]byte, 2)
	binary.LittleEndian.PutUint16(res, uint16(num))
	return res
}

func Uint32ToBytes(num uint32) []byte {
	res := make([]byte, 4)
	binary.LittleEndian.PutUint32(res, uint32(num))
	return res
}

func Uint64ToBytes(num uint64) []byte {
	res := make([]byte, 8)
	binary.LittleEndian.PutUint64(res, uint64(num))
	return res
}
