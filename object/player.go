package object

type PlayerObject struct {
	MapObject
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
