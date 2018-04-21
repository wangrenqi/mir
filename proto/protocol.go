package proto

import (
	cp "mir-go/proto/client"
	sp "mir-go/proto/server"
	"encoding/binary"
	"log"
)

type Packet struct {
	Index int
	Data  interface{}
}

type Null struct{}

func BytesToStruct(bytes []byte, isServer bool) (int, interface{}) {
	log.Println(bytes)
	var data interface{}
	index := int(binary.LittleEndian.Uint16(bytes[0:2]))

	bytes = bytes[2:]
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
			data = cp.GetClientVersion(bytes)
		case cp.DISCONNECT:
			data = cp.GetDisconnect(bytes)
		case cp.KEEPALIVE:
			data = cp.GetKeepAlive(bytes)
		case cp.NEW_ACCOUNT:
			data = cp.GetNewAccount(bytes)
		case cp.CHANGE_PASSWORD:
			data = cp.GetChangePassword(bytes)
		case cp.LOGIN:
			data = cp.GetLogin(bytes)
		case cp.NEW_CHARACTER:
			data = cp.GetNewCharacter(bytes)
		case cp.DELETE_CHARACTER:
			data = cp.GetDeleteCharacter(bytes)
		case cp.START_GAME:
			data = cp.GetStartGame(bytes)
		case cp.LOGOUT:
			data = cp.GetLogout(bytes)
		case cp.TURN:
			data = cp.GetTurn(bytes)
		case cp.WALK:
			data = cp.GetWalk(bytes)
		case cp.RUN:
			data = cp.GetRun(bytes)
		case cp.CHAT:
			data = cp.GetChat(bytes)
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
	bufLen := len(buffer)

	var i int
	for i = 0; i < bufLen; i = i + 1 {
		if bufLen < 4 {
			break
		}
		dataLen := int(buffer[i+1]<<8 + buffer[i])
		readerChan <- buffer[2+i : dataLen+i]
		i = i + dataLen - 1
	}
	if i >= bufLen {
		return make([]byte, 0)
	}
	return buffer[i:]
}
