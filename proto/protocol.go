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

func getPacketInfo([]byte) (bool, int) {

	return true, 0
}

func ToPacket(bytes []byte) *Packet {
	isServer, index := getPacketInfo(bytes)
	if isServer {
		switch index {
		case sp.CONNECTED:
		case sp.CLIENT_VERSION:
			return &Packet{true, sp.CLIENT_VERSION, &sp.ClientVersion{}}
		}
	} else {
		switch index {
		case cp.CLIENT_VERSION:
			return &Packet{false, cp.CLIENT_VERSION, &cp.ClientVersion{}}
		case cp.DISCONNECT:
		}
	}
	return &Packet{}
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
