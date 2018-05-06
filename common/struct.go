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

type RandomItemStat struct {
	MaxDuraChance            byte
	MaxDuraStatChance        byte
	MaxDuraMaxStat           byte
	MaxAcChance              byte
	MaxAcStatChance          byte
	MaxAcMaxStat             byte
	MaxMacChance             byte
	MaxMacStatChance         byte
	MaxMacMaxStat            byte
	MaxDcChance              byte
	MaxDcStatChance          byte
	MaxDcMaxStat             byte
	MaxMcChance              byte
	MaxMcStatChance          byte
	MaxMcMaxStat             byte
	MaxScChance              byte
	MaxScStatChance          byte
	MaxScMaxStat             byte
	AccuracyChance           byte
	AccuracyStatChance       byte
	AccuracyMaxStat          byte
	AgilityChance            byte
	AgilityStatChance        byte
	AgilityMaxStat           byte
	HpChance                 byte
	HpStatChance             byte
	HpMaxStat                byte
	MpChance                 byte
	MpStatChance             byte
	MpMaxStat                byte
	StrongChance             byte
	StrongStatChance         byte
	StrongMaxStat            byte
	MagicResistChance        byte
	MagicResistStatChance    byte
	MagicResistMaxStat       byte
	PoisonResistChance       byte
	PoisonResistStatChance   byte
	PoisonResistMaxStat      byte
	HpRecovChance            byte
	HpRecovStatChance        byte
	HpRecovMaxStat           byte
	MpRecovChance            byte
	MpRecovStatChance        byte
	MpRecovMaxStat           byte
	PoisonRecovChance        byte
	PoisonRecovStatChance    byte
	PoisonRecovMaxStat       byte
	CriticalRateChance       byte
	CriticalRateStatChance   byte
	CriticalRateMaxStat      byte
	CriticalDamageChance     byte
	CriticalDamageStatChance byte
	CriticalDamageMaxStat    byte
	FreezeChance             byte
	FreezeStatChance         byte
	FreezeMaxStat            byte
	PoisonAttackChance       byte
	PoisonAttackStatChance   byte
	PoisonAttackMaxStat      byte
	AttackSpeedChance        byte
	AttackSpeedStatChance    byte
	AttackSpeedMaxStat       byte
	LuckChance               byte
	LuckStatChance           byte
	LuckMaxStat              byte
	CurseChance              byte
}

func (self *RandomItemStat) ToBytes() []byte {
	// TODO
	return nil
}

func GetWeaponRandomItemStat() RandomItemStat {
	return RandomItemStat{
		MaxDuraChance:     2,
		MaxDuraStatChance: 13,
		MaxDuraMaxStat:    13,

		MaxDcChance:     15,
		MaxDcStatChance: 15,
		MaxDcMaxStat:    13,

		MaxMcChance:     20,
		MaxMcStatChance: 15,
		MaxMcMaxStat:    13,

		MaxScChance:     20,
		MaxScStatChance: 15,
		MaxScMaxStat:    13,

		AttackSpeedChance:     60,
		AttackSpeedStatChance: 30,
		AttackSpeedMaxStat:    3,

		StrongChance:     24,
		StrongStatChance: 20,
		StrongMaxStat:    2,

		AccuracyChance:     30,
		AccuracyStatChance: 20,
		AccuracyMaxStat:    2,
	}
}

// TODO
func GetArmourRandomItemStat() RandomItemStat    { return RandomItemStat{} }
func GetHelmetRandomItemStat() RandomItemStat    { return RandomItemStat{} }
func GetBeltBootsRandomItemStat() RandomItemStat { return RandomItemStat{} }
func GetNecklaceRandomItemStat() RandomItemStat  { return RandomItemStat{} }
func GetBraceletRandomItemStat() RandomItemStat  { return RandomItemStat{} }
func GetRingRandomItemStat() RandomItemStat      { return RandomItemStat{} }
func GetMountRandomItemStat() RandomItemStat     { return RandomItemStat{} }
