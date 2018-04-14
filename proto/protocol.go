package proto

import (
	cp "mir-go/proto/client"
	sp "mir-go/proto/server"
	"encoding/binary"
)

type Packet struct {
	IsServer bool
	Index    int
	Data     interface{}
}

type Null struct{}

func GetBytesLength(bytes []byte) int {
	if len(bytes) == 0 {
		return 0
	}
	return int(bytes[1]<<8 + bytes[0])
}

func getPacketIndex(bytes []byte) (isServer bool, index int) {
	length := GetBytesLength(bytes)
	if length == 0 {
		return false, -1
	}
	if length > len(bytes) || length < 2 {
		return false, -1
	}
	index = int(binary.LittleEndian.Uint16(bytes[2:4]))
	if index > 250 {
		return false, -1
	}
	return false, index
}

func getPacketData(isServer bool, index int) interface{} {
	var data interface{}
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
		case cp.DISCONNECT:
		default:
			data = &Null{}
		}
	}
	return data
}

func ToPacket(bytes []byte) *Packet {
	isServer, index := getPacketIndex(bytes)
	pkg := &Packet{isServer, index, getPacketData(isServer, index)}
	return pkg
}

func (pkg *Packet) ToBytes() []byte {
	if pkg.IsServer {
		switch pkg.Index {
		case sp.CONNECTED:
		case sp.CLIENT_VERSION:
		}
	} else {
		switch pkg.Index {
		case cp.CLIENT_VERSION:
			return []byte{24, 0, 0, 0, 16, 0, 0, 0, 196, 46, 198, 6, 217, 38, 102, 128, 242, 128, 185, 164, 66, 146, 36, 34}
		case cp.DISCONNECT:
		case cp.KEEPALIVE:
		case cp.NEW_ACCOUNT:
		case cp.CHANGE_PASSWORD:
		case cp.LOGIN:
			//data := pkg.Data.(*cp.Login)
			return []byte{15, 0, 5, 0, 3, 50, 50, 50, 6, 50, 50, 50, 50, 50, 50}
		case cp.NEW_CHARACTER:
		case cp.DELETE_CHARACTER:
		case cp.START_GAME:
			return []byte{8, 0, 8, 0, 2, 0, 0, 0}
		case cp.LOGOUT:
		case cp.TURN:
		case cp.WALK:
			data := pkg.Data.(*cp.Walk)
			// up upright right downright down downleft left upleft
			// 0 ~ 7
			return []byte{5, 0, 11, 0, byte(data.Dir)}
		case cp.RUN:
		case cp.CHAT:
		}
	}
	return nil
}
