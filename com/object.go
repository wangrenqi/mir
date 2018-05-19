package com

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
	Image        uint16
	// TODO
}
