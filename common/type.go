package common

type MirDirection byte

const (
	UP         MirDirection = iota
	UP_RIGHT
	RIGHT
	DOWN_RIGHT
	DOWN
	DOWN_LEFT
	LEFT
	UP_LEFT
)

type MirGender byte

const (
	MALE   MirGender = iota
	FEMALE
)

type MirClass byte

const (
	WARRIOR  MirClass = iota
	WIZARD
	TAOIST
	ASSASSIN
	ARCHER
)

// public enum LevelEffects : byte
// {
// 		None = 0,
// 		ist = 0x0001,
// 		RedDragon = 0x0002,
// 		BlueDragon = 0x0004
// }