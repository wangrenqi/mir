package com

import (
	"path/filepath"
	"io/ioutil"
	"sync/atomic"
	"net"
	"github.com/jinzhu/gorm"
)

var objectId uint32 = 0
var Maps map[uint32]Map
var AOIs map[uint32][]AOIEntity

const WIDTH = 20

func InitEnviron() *Environ {
	db := GetDB()
	db.AutoMigrate(&AccountInfo{}, &CharacterInfo{}, &RespawnInfo{}, &MonsterInfo{})
	maps := InitMaps(MapFilesPath)
	aoi := make(map[uint32][]AOIEntity)
	for i, m := range *maps {
		m.Index = i
		LoadNPC(&m, db)
		LoadMonster(&m, db)
		aoi[i] = InitAOIEntities(&m)
		AOIs = aoi
	}
	return &Environ{DB: db, Maps: maps, AOI: &aoi}
}

func GetMapObjectId() uint32 {
	res := objectId
	atomic.AddUint32(&objectId, 1)
	return res
}

func InitMaps(path string) *map[uint32]Map {
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
	Maps = make(map[uint32]Map)
	Maps[index] = tmp
	return &Maps
}

func GetMaps() *map[uint32]Map {
	return &Maps
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
	points := make([]*Point, 0)
	for i := center.X - int32(spread); i < center.X+int32(spread); i ++ {
		for j := center.Y - int32(spread); j < center.Y+int32(spread); j ++ {
			p := m.PointProxy[string(i)+","+string(j)]
			points = append(points, p)
		}
	}
	// TODO check that point has not object
	mapLen := len(points)
	for {
		randInt := RandomInt(0, mapLen)
		p := points[randInt]
		if p.Valid {
			p.Valid = false
			return p
		}
	}
}

// 返回a的conn以及附近8个aoi area 的conn
func (this AOIEntity) GetNearlyPlayerConnections() []net.Conn {
	aois := this.GetNearlyEightAOIs()
	conns := make([]net.Conn, 0)
	for _, aoi := range aois {
		for _, conn := range *aoi.Connections {
			conns = append(conns, conn)
		}
	}
	return conns
}

func (this AOIEntity) ValidPoint(m Map, point Point) bool {
	// TODO check aoi has monster NPC object
	p := m.PointProxy[string(point.X)+","+string(point.Y)]
	return p.Valid
}

// 根据一个 aoi 返回附近8个 aoi
func (this AOIEntity) GetNearlyEightAOIs() []AOIEntity {
	aois := make([]AOIEntity, 0)
	for _, a := range AOIs[this.MapIndex] {
		if (a.X == this.X-20 && a.Y == this.Y-20) || (a.X == this.X && a.Y == this.Y-20) || (a.X == this.X+20 && a.Y == this.Y-20) {
			aois = append(aois, a)
		}
		if (a.X == this.X-20 && a.Y == this.Y) || (a.X == this.X+20 && a.Y == this.Y) {
			aois = append(aois, a)
		}
		if (a.X == this.X-20 && a.Y == this.Y+20) || (a.X == this.X && a.Y == this.Y+20) || (a.X == this.X+20 && a.Y == this.Y+20) {
			aois = append(aois, a)
		}
	}
	return aois
}

// 返回 aoi 所有 monster object
func (this AOIEntity) GetMonsterObjects() []MonsterObject {
	maps := GetMaps()
	thisMap := (*maps)[this.MapIndex]
	monsterObjects := (*thisMap.Objects)["monster"].([]MonsterObject)
	res := make([]MonsterObject, 0)
	for _, m := range monsterObjects {
		if m.CurrentLocation.X > int32(this.X) && m.CurrentLocation.Y > int32(this.Y) && m.CurrentLocation.X < int32(this.X)+WIDTH && m.CurrentLocation.Y < int32(this.Y)+WIDTH {
			res = append(res, m)
		}
	}
	return res
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
	aois := make([]AOIEntity, 0)
	index := uint16(0)
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j ++ {
			aois = append(aois, AOIEntity{
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
			for _, a := range aois {
				if x > int(a.X) && x < int(a.X)+WIDTH && y > int(a.Y) && y < int(a.Y)+WIDTH {
					a.Points = append(a.Points, m.PointProxy[string(x)+","+string(y)])
				}
			}
		}
	}
	return aois
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

// 把 db 中的 monster info 转成内存中的 monster object
func MonsterInfoToMonsterObject(info MonsterInfo, mapInfo Map) *MonsterObject {
	obj := &MonsterObject{}
	obj.ObjectID = GetMapObjectId()
	obj.Name = info.Name
	obj.Image = info.Image
	//CurrentMap Map
	//ExplosionInflictedTime int64 ??
	//ExplosionInflictedStage int64 ??
	//SpawnThread int32 ??
	obj.MonsterIndex = info.MonsterIndex
	obj.CurrentMapIndex = mapInfo.Index
	obj.CurrentLocation = Point{}
	obj.Direction = MirDirection(RandomInt(0, 7))
	obj.Level = info.Level
	obj.Health = uint32(info.HP)
	obj.MaxHealth = uint32(info.HP)
	obj.PercentHealth = 100
	return obj
}

func LoadNPC(m *Map, db *gorm.DB) {
	// TODO
}

var MapRespawnCount map[uint32]map[uint32]uint32 // MapIndex MonsterIndex MonsterCount

func LoadMonster(m *Map, db *gorm.DB) {
	MapRespawnCount = make(map[uint32]map[uint32]uint32)
	var respawnInfos []RespawnInfo
	db.Where(&RespawnInfo{MapIndex: m.Index}).Find(&respawnInfos)

	var monsterObjects []MonsterObject
	(*m.Objects)["monster"] = monsterObjects
	for _, respawnInfo := range respawnInfos {
		respawnCount := respawnInfo.Count
		var monsterInfo MonsterInfo
		db.Where(&MonsterInfo{Index: monsterInfo.MonsterIndex}).Find(&monsterInfo)
		if monsterInfo.Index == 0 {
			continue
		}
		MapAddMonster(m.Index, respawnInfo, respawnCount, db)
	}
}

func GetMapExistedMonsterCount(mapIndex uint32, monsterIndex uint32) uint32 {
	if MapRespawnCount[mapIndex] == nil {
		return 0
	}
	return MapRespawnCount[mapIndex][monsterIndex]
}

func GetMapRespawnInfos(mapIndex uint32, db *gorm.DB) []RespawnInfo {
	var respawnInfos []RespawnInfo
	db.Where(&RespawnInfo{MapIndex: mapIndex}).Find(&respawnInfos)
	return respawnInfos
}

func MapAddMonster(mapIndex uint32, respawnInfo RespawnInfo, addCount uint32, db *gorm.DB) {
	// TODO spawn time
	// if not spawn time then pass
	m := Maps[mapIndex]
	monsterObjects := (*m.Objects)["monster"].([]MonsterObject)
	var monsterInfo MonsterInfo
	db.Where(&MonsterInfo{MonsterIndex: respawnInfo.MonsterIndex}).Find(&monsterInfo)
	for i := uint32(0); i < addCount; i++ {
		randPoint := GetRandomPoint(&m, Point{X: respawnInfo.LocationX, Y: respawnInfo.LocationY}, respawnInfo.Spread)
		monsterObject := MonsterInfoToMonsterObject(monsterInfo, m)
		monsterObject.MapObject.CurrentLocation = *randPoint
		//randPoint.Valid = false
		monsterObjects = append(monsterObjects, *monsterObject)
	}
	existedCount := uint32(len(monsterObjects))
	monsterCount := MapRespawnCount[m.Index]
	if monsterCount == nil {
		MapRespawnCount[m.Index] = make(map[uint32]uint32)
	}
	MapRespawnCount[m.Index][monsterInfo.Index] = existedCount + addCount
	(*m.Objects)["monster"] = monsterObjects
}

func GetMapMonsterObject(mapIndex uint32, point Point) MonsterObject {
	m := (*GetMaps())[mapIndex]
	monsters := (*m.Objects)["monster"].([]MonsterObject)
	for _, obj := range monsters {
		if point.X == obj.CurrentLocation.X && point.Y == obj.CurrentLocation.Y {
			return obj
		}
	}
	return MonsterObject{}
}

func GetMapNPCObject(mapIndex uint32, point Point) NPCObject {
	return NPCObject{}
}
