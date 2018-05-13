package com

import (
	"path/filepath"
	"io/ioutil"
	"sync/atomic"
	"net"
)

var objectId uint32 = 0

const WIDTH = 20

func InitEnviron() *Environ {
	db := GetDB()
	db.AutoMigrate(&AccountInfo{}, &CharacterInfo{}, &RespawnInfo{}, &MonsterInfo{})
	maps := GetMaps(MapFilesPath)
	aoi := make(map[uint32][]AOIEntity)
	for i, m := range *maps {
		m.Index = i
		LoadNPC(&m, db)
		LoadMonster(&m, db)
		aoi[i] = InitAOIEntities(&m)
	}
	return &Environ{DB: db, Maps: maps, AOI: &aoi}
}

func GetMapObjectId() uint32 {
	res := objectId
	atomic.AddUint32(&objectId, 1)
	return res
}

func GetMaps(path string) *map[uint32]Map {
	// TODO for map in path, loop read map and return []Map
	fileBytes, err := filepath.Abs(path + "/0.map")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadFile(fileBytes)
	if err != nil {
		panic(err)
	}
	//typ := FindType(bytes)
	//log.Println("map typ ->", typ)
	tmp := GetMapV1(bytes)

	//saveToFile(tmp)

	index := uint32(12289) // TODO
	maps := make(map[uint32]Map)
	maps[index] = tmp
	return &maps
}

func SaveToFile(tmp Map) {
	points := *tmp.Points
	str := ""
	index := 0
	for _, p := range points {
		if p.Valid == true {
			str = str + " "
		} else {
			str = str + "*"
		}
		index = index + 1
		if index == 700 {
			index = 0
			str = str + "\n"
		}
	}
	err := ioutil.WriteFile("output.txt", []byte(str), 0644)
	if err != nil {
		panic(err)
	}
	//fmt.Println("saved...")
}

func GetMapV1(bytes []byte) Map {
	offset := 21
	w := BytesToUint16(bytes[offset : offset+2])
	offset += 2
	xor := BytesToUint16(bytes[offset : offset+2])
	offset += 2
	h := BytesToUint16(bytes[offset : offset+2])
	width := w ^ xor
	height := h ^ xor
	//fmt.Println(width, height)
	pointProxy := make(map[string]*Point, 0)

	offset = 54
	index := 0
	points := make([]Point, int(width)*int(height))
	for i := 0; i < int(width); i ++ {
		for j := 0; j < int(height); j ++ {
			valid := true

			if (BytesToUint32(bytes[offset:offset+4])^0xAA38AA38)&0x20000000 != 0 {
				valid = false
			}
			if ((BytesToUint16(bytes[offset+6:offset+8]) ^ xor) & 0x8000) != 0 {
				valid = false
			}
			p := Point{X: int32(i), Y: int32(j), Valid: valid}
			pointProxy[string(i)+","+string(j)] = &p
			points[index] = p
			index ++
			offset += 15
		}
	}
	objects := make(map[string]interface{}, 0)
	m := Map{Width: width, Height: height, Points: &points, Objects: &objects, PointProxy: pointProxy}
	return m
}

func GetStartPoint() Point {
	// TODO Random
	return Point{X: 287, Y: 612}
}

func GetRandomPoint(m *Map, center Point, spread uint32) *Point {
	// TODO !!!优化算法 根据给定点，取该点spread范围内所有点，而不是map上所有点
	points := *m.Points
	mapLen := len(points)
	for {
		randInt := RandomInt(0, mapLen)
		p := points[randInt]
		if p.Valid {
			return &p
		}
	}
}

// TODO 返回a的conn以及附近8个aoi area 的conn
func (a AOIEntity) GetNearlyPlayerConnections() []net.Conn {
	return nil
}

func (self AOIEntity) ValidPoint(m Map, point Point) bool {
	// TODO check aoi has monster NPC object
	p := m.PointProxy[string(point.X)+","+string(point.Y)]
	return p.Valid
}

func points() []*Point {
	return make([]*Point, 0)
}

func connections() *map[int32]net.Conn {
	conns := make(map[int32]net.Conn, 0)
	return &conns
}

// 把map object 切成很多个aoi entity
func InitAOIEntities(m *Map) []AOIEntity {
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
func GetAOIEntity(aoi []AOIEntity, p Point) *AOIEntity {
	//log.Println(p.X, p.Y)
	for _, a := range aoi {
		//log.Println("--->", a.X, a.Y)
		if int32(a.X) <= p.X && int32(a.X)+WIDTH > p.X && int32(a.Y) <= p.Y && int32(a.Y)+WIDTH > p.Y {
			return &a
		}
	}
	return nil
}
