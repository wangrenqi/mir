package env

import (
	"net"
	cm "mir/common"
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

// TODO 根据点坐标和地图索引，返回该点在地图上所在的区域
func GetAOIEntity(mapIndex uint32, point cm.Point) *AOIEntity {
	return nil
}

// TODO 返回a的conn以及附近8个aoi area 的conn
func (a AOIEntity) GetNearlyPlayerConnections() []net.Conn {
	return nil
}
