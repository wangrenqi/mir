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
	x := self.X
	y := self.Y
	switch direction {
	case UP:
		y = y - distance
	case UP_RIGHT:
		x = x + distance
		y = y - distance
	case RIGHT:
		x = x + distance
	case DOWN_RIGHT:
		x = x + distance
		y = y + distance
	case DOWN:
		y = y + distance
	case DOWN_LEFT:
		x = x - distance
		y = y + distance
	case LEFT:
		x = x - distance
	case UP_LEFT:
		x = x - distance
		y = y - distance
	}
	return Point{X: x, Y: y}
}
