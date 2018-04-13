package proto

import (
	cp "mir-go/proto/client"
	sp "mir-go/proto/server"
)

type Packet struct {
	IsServer bool
	Index    int
	Data     interface{}
}

type Null struct{}

func getPacketInfo([]byte) (bool, int) {

	return true, 0
}

func ToPacket(bytes []byte) *Packet {
	isServer, index := getPacketInfo(bytes)
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
	return &Packet{isServer, index, data}
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
		case cp.DISCONNECT:
		}
	}
	return nil
}
