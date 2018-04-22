package proto

import (
	cp "mir/proto/client"
	"encoding/binary"
	"log"
)

type Packet struct {
	Index int
	Data  interface{}
}

type Null struct{}

func (self *Null) ToBytes() []byte {
	return make([]byte, 0)
}

func BytesToStruct(bytes []byte) (int, interface{}) {
	var data interface{}
	index := int(binary.LittleEndian.Uint16(bytes[0:2]))
	log.Println("BytesToStruct bytes ->", bytes, "index ->", index)
	bytes = bytes[2:]
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
	return index, data
}

type Parser interface {
	ToBytes() []byte
}

func (pkg *Packet) ToBytes() []byte {
	var parser Parser
	switch pkg.Index {
	case cp.CLIENT_VERSION:
		parser = pkg.Data.(*cp.ClientVersion)
	case cp.DISCONNECT:
		parser = pkg.Data.(*cp.Disconnect)
	case cp.KEEPALIVE:
		parser = pkg.Data.(*cp.KeepAlive)
	case cp.NEW_ACCOUNT:
		parser = pkg.Data.(*cp.NewAccount)
	case cp.CHANGE_PASSWORD:
		parser = pkg.Data.(*cp.ChangePassword)
	case cp.LOGIN:
		parser = pkg.Data.(*cp.Login)
	case cp.NEW_CHARACTER:
		parser = pkg.Data.(*cp.NewCharacter)
	case cp.DELETE_CHARACTER:
		parser = pkg.Data.(*cp.DeleteCharacter)
	case cp.START_GAME:
		parser = pkg.Data.(*cp.StartGame)
	case cp.LOGOUT:
		parser = pkg.Data.(*cp.Logout)
	case cp.TURN:
		parser = pkg.Data.(*cp.Turn)
	case cp.WALK:
		parser = pkg.Data.(*cp.Walk)
	case cp.RUN:
		parser = pkg.Data.(*cp.Run)
	case cp.CHAT:
		parser = pkg.Data.(*cp.Chat)
	default:
		parser = &Null{}
	}
	return parser.ToBytes()
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
