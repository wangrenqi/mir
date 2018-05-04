package object

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
