package proto

import (
	cp "mir-go/proto/client"
	sp "mir-go/proto/server"
	"encoding/binary"
	"bytes"
)

type Packet struct {
	Index int
	Data  interface{}
}

type Null struct{}

//func GetBytesLength(bytes []byte) int {
//	if len(bytes) == 0 {
//		return 0
//	}
//	return int(bytes[1]<<8 + bytes[0])
//}

//func getPacketIndex(bytes []byte) (isServer bool, index int) {
//	length := GetBytesLength(bytes)
//	if length == 0 {
//		return false, -1
//	}
//	if length > len(bytes) || length < 2 {
//		return false, -1
//	}
//	index = int(binary.LittleEndian.Uint16(bytes[2:4]))
//	if index > 250 {
//		return false, -1
//	}
//	return false, index
//}

func BytesToStruct(bytes []byte, isServer bool) (int, interface{}) {
	var data interface{}
	index := int(binary.LittleEndian.Uint16(bytes[0:2]))

	if isServer {
		switch index {
		case sp.CONNECTED:
		case sp.CLIENT_VERSION:
			data = &sp.ClientVersion{}
		default:
			data = &Null{}
		}
	} else {
		switch index {
		case cp.CLIENT_VERSION:
			data = &cp.ClientVersion{}
		case cp.DISCONNECT:
			data = &cp.Disconnect{}
		case cp.KEEPALIVE:
			data = &cp.KeepAlive{}
		case cp.NEW_ACCOUNT:
			data = &cp.NewAccount{}
		case cp.CHANGE_PASSWORD:
			data = &cp.ChangePassword{}
		case cp.LOGIN:
			data = &cp.Login{}
		case cp.NEW_CHARACTER:
			data = &cp.NewCharacter{}
		case cp.DELETE_CHARACTER:
			data = &cp.DeleteCharacter{}
		case cp.START_GAME:
			data = &cp.StartGame{}
		case cp.LOGOUT:
			data = &cp.Logout{}
		case cp.TURN:
			data = &cp.Turn{}
		case cp.WALK:
			data = &cp.Walk{}
		case cp.RUN:
			data = &cp.Run{}
		case cp.CHAT:
			data = &cp.Chat{}
		default:
			data = &Null{}
		}
	}
	return index, data
}

func (pkg *Packet) ToBytes(isServer bool) []byte {
	if isServer {
		switch pkg.Index {
		case sp.CONNECTED:
		case sp.CLIENT_VERSION:
		}
	} else {
		switch pkg.Index {
		case cp.CLIENT_VERSION:
			//24, 0 (22 + 2)
			return []byte{0, 0, 16, 0, 0, 0, 196, 46, 198, 6, 217, 38, 102, 128, 242, 128, 185, 164, 66, 146, 36, 34}
		case cp.DISCONNECT:
		case cp.KEEPALIVE:
		case cp.NEW_ACCOUNT:
		case cp.CHANGE_PASSWORD:
		case cp.LOGIN:
			//data := pkg.Data.(*cp.Login)
			// 15, 0 (13 + 2)
			return []byte{5, 0, 3, 50, 50, 50, 6, 50, 50, 50, 50, 50, 50}
		case cp.NEW_CHARACTER:
		case cp.DELETE_CHARACTER:
		case cp.START_GAME:
			// 8, 0 (6 + 2)
			return []byte{8, 0, 2, 0, 0, 0}
		case cp.LOGOUT:
		case cp.TURN:
		case cp.WALK:
			data := pkg.Data.(*cp.Walk)
			// up upright right downright down downleft left upleft
			// 5, 0 (3 + 2)
			return []byte{11, 0, byte(data.Dir)}
		case cp.RUN:
		case cp.CHAT:
			// 20, 0 (18 + 2)
			return []byte{13, 0, 15, 228, 189, 160, 229, 165, 189, 229, 149, 138, 54, 54, 54, 239, 189, 129}
		}
	}
	return nil
}

// 封包
func Pack(data []byte) []byte {
	length := len(data) + 2
	header := make([]byte, 2)
	binary.LittleEndian.PutUint16(header, uint16(length))
	return append(header, data...)
}

// 解包
func UnPack(buffer []byte, readerChan chan []byte) []byte {
	// TODO
	data := buffer[:]
	readerChan <- data

	return nil
}

//字节转换成整形
func ByteToInt(n []byte) int {

	return 0
}

//整数转换成字节
func IntToBytes(n int) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}
