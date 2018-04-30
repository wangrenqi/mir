package engine

import (
	"mir/orm"
	"github.com/jinzhu/gorm"
	"path/filepath"
	"io/ioutil"
	"log"
	cm "mir/common"
)

type Environ struct {
	Db   *gorm.DB
	Maps interface{}
}

type Point struct {
	//Index int // Map Points数组里的索引
	X     int
	Y     int
	Valid bool
}

type Map struct {
	Witdh  uint16
	Height uint16
	Points interface{} // Point
	Object interface{}
}

var mapFilesPath = "/Users/Mccree/gopath/src/mir/engine/maps"

func InitEnviron() *Environ {
	db := orm.GetDB()

	db.AutoMigrate(&orm.Account{}, &orm.SelectInfo{})

	maps := getMaps(mapFilesPath)

	return &Environ{
		Db:   db,
		Maps: maps,
	}
}

func getMaps(path string) []Map {
	// TODO for map in path, loop read map and return []Map
	fileBytes, err := filepath.Abs(path + "/0.map")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadFile(fileBytes)
	if err != nil {
		panic(err)
	}
	typ := FindType(bytes)
	log.Println("map typ ->", typ)
	tmp := getMapV1(bytes)

	//saveToFile(tmp)

	maps := make([]Map, 0)
	return append(maps, tmp)
}

func saveToFile(tmp Map) {
	points := tmp.Points.([]Point)
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
	check(err)
	//fmt.Println("saved...")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getMapV1(bytes []byte) Map {
	offset := 21
	w := cm.BytesToUint16(bytes[offset : offset+2])
	offset += 2
	xor := cm.BytesToUint16(bytes[offset : offset+2])
	offset += 2
	h := cm.BytesToUint16(bytes[offset : offset+2])
	width := w ^ xor
	height := h ^ xor
	//fmt.Println(width, height)

	offset = 54
	index := 0
	points := make([]Point, int(width)*int(height))
	for i := 0; i < int(width); i ++ {
		for j := 0; j < int(height); j ++ {
			valid := true

			if (cm.BytesToUint32(bytes[offset:offset+4])^0xAA38AA38)&0x20000000 != 0 {
				valid = false
			}
			if ((cm.BytesToUint16(bytes[offset+6:offset+8]) ^ xor) & 0x8000) != 0 {
				valid = false
			}
			p := Point{X: i, Y: j, Valid: valid}
			points[index] = p
			index ++
			offset += 15
		}
	}
	m := Map{Witdh: width, Height: height, Points: points}
	return m
}

func FindType(bytes []byte) int {
	// TODO 直接从c# 源码copy ，未测试
	// TODO 用switch case
	//c# custom map format
	if (bytes[2] == 0x43) && (bytes[3] == 0x23) {
		return 100
	}
	//wemade mir3 maps have no title they just start with blank bytes
	if bytes[0] == 0 {
		return 5
	}
	//shanda mir3 maps start with title: (C) SNDA, MIR3.
	if (bytes[0] == 0x0F) && (bytes[5] == 0x53) && (bytes[14] == 0x33) {
		return 6
	}
	//wemades antihack map (laby maps) title start with: Mir2 AntiHack
	if (bytes[0] == 0x15) && (bytes[4] == 0x32) && (bytes[6] == 0x41) && (bytes[19] == 0x31) {
		return 4
	}
	//wemades 2010 map format i guess title starts with: Map 2010 Ver 1.0
	if (bytes[0] == 0x10) && (bytes[2] == 0x61) && (bytes[7] == 0x31) && (bytes[14] == 0x31) {
		return 1
	}
	//shanda's 2012 format and one of shandas(wemades) older formats share same header info, only difference is the filesize
	if (bytes[4] == 0x0F) && (bytes[18] == 0x0D) && (bytes[19] == 0x0A) {
		W := int(bytes[0] + (bytes[1] << 8))
		H := int(bytes[2] + (bytes[3] << 8))
		if len(bytes) > (52 + (W * H * 14)) {
			return 3
		} else {
			return 2
		}
	}
	//3/4 heroes map format (myth/lifcos i guess)
	if (bytes[0] == 0x0D) && (bytes[1] == 0x4C) && (bytes[7] == 0x20) && (bytes[11] == 0x6D) {
		return 7
	}
	return 0
}
