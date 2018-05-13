package com

import "github.com/jinzhu/gorm"

type MapObject struct {
	ObjectID uint32
	//Race ObjectType
	Name string
	//CurrentMap Map
	//ExplosionInflictedTime int64 ??
	//ExplosionInflictedStage int64 ??
	//SpawnThread int32 ??
	CurrentMapIndex uint32
	CurrentLocation Point
	Direction       MirDirection
	Level           uint16
	Health          uint32
	MaxHealth       uint32
	PercentHealth   byte
	// TODO
	//public ushort MinAC, MaxAC, MinMAC, MaxMAC;
	//public ushort MinDC, MaxDC, MinMC, MaxMC, MinSC, MaxSC;
	// ...
}

type PlayerObject struct {
	MapObject
	HP uint16
	MP uint16
}

func (self *PlayerObject) CanWalk() bool {
	return true
}

func (self *PlayerObject) CanMove() bool {
	return true
}

func (self *PlayerObject) CanRun() bool {
	return true
}

type MonsterObject struct {
	MapObject
	MonsterIndex uint32
	// TODO
}

func monsterInfoToMonsterObject(info MonsterInfo, mapInfo Map) MonsterObject {
	obj := MonsterObject{}
	obj.ObjectID = GetMapObjectId()
	obj.Name = info.Name
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

func LoadMonster(m *Map, db *gorm.DB) {
	var respawnInfos []RespawnInfo
	db.Where(&RespawnInfo{MapIndex: m.Index}).Find(&respawnInfos)

	var monsterObjects []MonsterObject
	for _, respawnInfo := range respawnInfos {
		respawnCount := respawnInfo.Count
		// get monster info by monster index
		var monsterInfo MonsterInfo
		db.Where(&MonsterInfo{Index: monsterInfo.MonsterIndex}).Find(&monsterInfo)
		// create monster object and save to map object
		if monsterInfo.Index == 0 {
			continue
		}
		for i := uint32(0); i < respawnCount; i++ {
			monsterObject := monsterInfoToMonsterObject(monsterInfo, *m)
			// random monster object point base on respawnInfo.spread
			randPoint := GetRandomPoint(m, Point{X: respawnInfo.LocationX, Y: respawnInfo.LocationY}, respawnInfo.Spread)
			monsterObject.MapObject.CurrentLocation = *randPoint
			randPoint.Valid = false
			monsterObjects = append(monsterObjects, monsterObject)
		}
	}
	(*m.Objects)["monster"] = monsterObjects
}
