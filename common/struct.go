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
