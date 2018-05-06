package env

import (
	"net"
)

type AOIEntity struct {
	Index       uint16
	MapIndex    uint32
	X           uint16
	Y           uint16
	Width       uint16
	Connections map[int32]net.Conn // client id : conn
}

func (m Map) GetAOIEntities() []AOIEntity {
	// TODO 把map object 切成很多个aoi entity
	return nil
}
