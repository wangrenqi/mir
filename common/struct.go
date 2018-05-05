package common

type Point struct {
	X     int32
	Y     int32
	Valid bool
}

func (self *Point) ToBytes() []byte {
	// TODO 未验证
	XBytes := Uint32ToBytes(uint32(self.X))
	YBytes := Uint32ToBytes(uint32(self.Y))
	return append(XBytes, YBytes...)
}

// TODO
func (self *Point) Move(direction MirDirection, distance int32) Point {
	return Point{}
}
