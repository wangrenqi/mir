package engine

import (
	cm "mir/common"
)

type MapObject struct {
	ObjectID uint32
	//Race ObjectType
	Name string
	//CurrentMap Map
	//ExplosionInflictedTime int64 ??
	//ExplosionInflictedStage int64 ??
	//SpawnThread int32 ??
	//CurrentMapIndex int32
	CurrentLocation cm.Point
	Direction       cm.MirDirection
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
}

func (self *PlayerObject) CanWalk() bool {
	return true
}

func (self *PlayerObject) CanMove() bool {
	return true
}
