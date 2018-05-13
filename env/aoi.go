package env

import (
	"net"
	cm "mir/common"
)

const WIDTH = 20

type AOIEntity struct {
	Index       uint16
	MapIndex    uint32
	X           uint16
	Y           uint16
	Connections *map[int32]net.Conn // client id : conn
	Points      []*cm.Point
}

func points() []*cm.Point {
	return make([]*cm.Point, 0)
}

func connections() *map[int32]net.Conn {
	conns := make(map[int32]net.Conn, 0)
	return &conns
}

// 把map object 切成很多个aoi entity
func (m *Map) InitAOIEntities() []AOIEntity {
	columns, rows := 0, 0
	if m.Width%WIDTH == 0 {
		columns = int(m.Width) / WIDTH
	} else {
		columns = int(m.Width)/WIDTH + 1
	}
	if m.Height%WIDTH == 0 {
		rows = int(m.Height) / WIDTH
	} else {
		rows = int(m.Height)/WIDTH + 1
	}
	aoi := make([]AOIEntity, 0)
	index := uint16(0)
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j ++ {
			aoi = append(aoi, AOIEntity{
				Index:       index,
				MapIndex:    m.Index,
				X:           uint16(i * WIDTH),
				Y:           uint16(j * WIDTH),
				Connections: connections(),
				Points:      points(),
			})
			index = index + 1
		}
	}

	for x := 0; x < int(m.Width); x ++ {
		for y := 0; y < int(m.Height); y ++ {
			for _, a := range aoi {
				if x > int(a.X) && x < int(a.X)+WIDTH && y > int(a.Y) && y < int(a.Y)+WIDTH {
					a.Points = append(a.Points, m.PointProxy[string(x)+","+string(y)])
				}
			}
		}
	}
	return aoi
}

// 获得point 所在的aoi
func GetAOIEntity(aoi []AOIEntity, p cm.Point) *AOIEntity {
	//log.Println(p.X, p.Y)
	for _, a := range aoi {
		//log.Println("--->", a.X, a.Y)
		if int32(a.X) <= p.X && int32(a.X)+WIDTH > p.X && int32(a.Y) <= p.Y && int32(a.Y)+WIDTH > p.Y {
			return &a
		}
	}
	return nil
}

// TODO 返回a的conn以及附近8个aoi area 的conn
func (a AOIEntity) GetNearlyPlayerConnections() []net.Conn {
	return nil
}

func (self AOIEntity) ValidPoint(point cm.Point) bool {
	return true
}
