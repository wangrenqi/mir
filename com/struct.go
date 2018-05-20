package com

import (
	"net"
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron"
)

type Client struct {
	Id        int32
	Conn      net.Conn
	ReqChan   <-chan []byte
	Env       *Environ
	Status    int
	Info      map[string]interface{}
	Player    *PlayerObject
	AOIEntity *AOIEntity
}

type AOIEntity struct {
	Index       uint16
	MapIndex    uint32
	X           uint16
	Y           uint16
	Connections *map[int32]net.Conn // client id : conn
	Points      []*Point
}

type Map struct {
	Index      uint32
	Width      uint16
	Height     uint16
	Points     *[]Point
	PointProxy map[string]*Point
	Objects    *map[string]interface{}
}

type Environ struct {
	DB   *gorm.DB
	Maps *map[uint32]Map
	AOI  *map[uint32][]AOIEntity
	Cron *cron.Cron
}

type Point struct {
	X     int32
	Y     int32
	Valid bool
}

func (self *Point) ToBytes() []byte {
	XBytes := Uint32ToBytes(uint32(self.X))
	YBytes := Uint32ToBytes(uint32(self.Y))
	return append(XBytes, YBytes...)
}

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

type ItemInfo struct {
	Index            uint32
	Name             string
	Type             ItemType
	Grade            ItemGrade
	RequiredType     RequiredType   // default Level
	RequiredClass    RequiredClass  // default None
	RequiredGender   RequiredGender // default None
	Set              ItemSet
	Shape            uint16
	Weight           byte
	Light            byte
	RequiredAmount   byte
	Image            uint16
	Durability       uint16
	Price            uint32
	StackSize        uint32 //default 1;
	MinAC            byte
	MaxAC            byte
	MinMAC           byte
	MaxMAC           byte
	MinDC            byte
	MaxDC            byte
	MinMC            byte
	MaxMC            byte
	MinSC            byte
	MaxSC            byte
	Accuracy         byte
	Agility          byte
	HP               uint16
	MP               uint16
	AttackSpeed      int8 // 需要是负数
	Luck             int8
	BagWeight        byte
	HandWeight       byte
	WearWeight       byte
	StartItem        bool
	Effect           byte
	Strong           byte
	MagicResist      byte
	PoisonResist     byte
	HealthRecovery   byte
	SpellRecovery    byte
	PoisonRecovery   byte
	HPrate           byte
	MPrate           byte
	CriticalRate     byte
	CriticalDamage   byte
	NeedIdentify     bool
	ShowGroupPickup  bool
	GlobalDropNotify bool
	ClassBased       bool
	LevelBased       bool
	CanMine          bool
	CanFastRun       bool
	CanAwakening     bool
	MaxAcRate        byte
	MaxMacRate       byte
	Holy             byte
	Freezing         byte
	PoisonAttack     byte
	HpDrainRate      byte
	Bind             uint16 // BindMode 这个枚举太大了，直接用uint16 // default none
	Reflect          byte
	Unique           uint16 // SpecialItemMode ?? // default None;
	RandomStatsId    byte
	RandomStats      RandomItemStat
	ToolTip          string //default ""
}

// TODO
func (self *ItemInfo) ToBytes() []byte {
	indexBytes := Uint32ToBytes(self.Index)
	nameBytes := StringToBytes(self.Name)
	typeBytes := []byte{byte(self.Type)}
	gradeBytes := []byte{byte(self.Grade)}
	requiredTypeBytes := []byte{byte(self.RequiredType)}
	requiredClassBytes := []byte{byte(self.RequiredClass)}
	requiredGenderBytes := []byte{byte(self.RequiredGender)}
	setBytes := []byte{byte(self.Set)}
	shapeBytes := Uint16ToBytes(self.Shape)
	weightBytes := []byte{self.Weight}
	lightBytes := []byte{self.Light}
	requiredAmountBytes := []byte{self.RequiredAmount}
	imageBytes := Uint16ToBytes(self.Image)
	durabilityBytes := Uint16ToBytes(self.Durability)
	priceBytes := Uint32ToBytes(self.Price)
	stackSizeBytes := Uint32ToBytes(self.StackSize)
	minACBytes := []byte{self.MinAC}
	maxACBytes := []byte{self.MaxAC}
	minMACBytes := []byte{self.MinMAC}
	maxMACBytes := []byte{self.MaxMAC}
	minDCBytes := []byte{self.MinDC}
	maxDCBytes := []byte{self.MaxDC}
	minMCBytes := []byte{self.MinMC}
	maxMCBytes := []byte{self.MaxMC}
	minSCBytes := []byte{self.MinSC}
	maxSCBytes := []byte{self.MaxSC}
	accuracyBytes := []byte{self.Accuracy}
	agilityBytes := []byte{self.Agility}
	hPBytes := Uint16ToBytes(self.HP)
	mPBytes := Uint16ToBytes(self.MP)
	attackSpeedBytes := []byte{byte(self.AttackSpeed)} // TODO int8 可能为负数，先当正数处理
	luckBytes := []byte{byte(self.Luck)}               // 同上
	bagWeightBytes := []byte{self.BagWeight}
	handWeightBytes := []byte{self.HandWeight}
	wearWeightBytes := []byte{self.WearWeight}
	startItemBytes := BoolToBytes(self.StartItem)
	effectBytes := []byte{self.Effect}
	strongBytes := []byte{self.Strong}
	magicResistBytes := []byte{self.MagicResist}
	poisonResistBytes := []byte{self.PoisonResist}
	healthRecoveryBytes := []byte{self.HealthRecovery}
	spellRecoveryBytes := []byte{self.SpellRecovery}
	poisonRecoveryBytes := []byte{self.PoisonRecovery}
	hPrateBytes := []byte{self.HPrate}
	mPrateBytes := []byte{self.MPrate}
	criticalRateBytes := []byte{self.CriticalRate}
	criticalDamageBytes := []byte{self.CriticalDamage}
	needIdentifyBytes := BoolToBytes(self.NeedIdentify)
	showGroupPickupBytes := BoolToBytes(self.ShowGroupPickup)
	globalDropNotifyBytes := BoolToBytes(self.GlobalDropNotify)
	classBasedBytes := BoolToBytes(self.ClassBased)
	levelBasedBytes := BoolToBytes(self.LevelBased)
	canMineBytes := BoolToBytes(self.CanMine)
	canFastRunBytes := BoolToBytes(self.CanFastRun)
	canAwakeningBytes := BoolToBytes(self.CanAwakening)
	maxAcRateBytes := []byte{self.MaxAcRate}
	maxMacRateBytes := []byte{self.MaxMacRate}
	holyBytes := []byte{self.Holy}
	freezingBytes := []byte{self.Freezing}
	poisonAttackBytes := []byte{self.PoisonAttack}
	hpDrainRateBytes := []byte{self.HpDrainRate}
	bindBytes := Uint16ToBytes(self.Bind)
	reflectBytes := []byte{self.Reflect}
	uniqueBytes := Uint16ToBytes(self.Unique)
	randomStatsIdBytes := []byte{self.RandomStatsId}
	randomStatsBytes := self.RandomStats.ToBytes()
	toolTipBytes := StringToBytes(self.ToolTip)
	result := make([]byte, 0)
	for _, r := range [][]byte{indexBytes, nameBytes, typeBytes, gradeBytes, requiredTypeBytes, requiredClassBytes,
		requiredGenderBytes, setBytes, shapeBytes, weightBytes, lightBytes, requiredAmountBytes, imageBytes,
		durabilityBytes, priceBytes, stackSizeBytes, minACBytes, maxACBytes, minMACBytes, maxMACBytes, minDCBytes,
		maxDCBytes, minMCBytes, maxMCBytes, minSCBytes, maxSCBytes, accuracyBytes, agilityBytes, hPBytes, mPBytes,
		attackSpeedBytes, luckBytes, bagWeightBytes, handWeightBytes, wearWeightBytes, startItemBytes, effectBytes,
		strongBytes, magicResistBytes, poisonResistBytes, healthRecoveryBytes, spellRecoveryBytes, poisonRecoveryBytes,
		hPrateBytes, mPrateBytes, criticalRateBytes, criticalDamageBytes, needIdentifyBytes, showGroupPickupBytes,
		globalDropNotifyBytes, classBasedBytes, levelBasedBytes, canMineBytes, canFastRunBytes, canAwakeningBytes,
		maxAcRateBytes, maxMacRateBytes, holyBytes, freezingBytes, poisonAttackBytes, hpDrainRateBytes, bindBytes,
		reflectBytes, uniqueBytes, randomStatsIdBytes, randomStatsBytes, toolTipBytes,
	} {
		result = append(result, r...)
	}
	return result
}

type UserItem struct {
	UniqueID       uint64
	ItemIndex      uint32
	Info           ItemInfo
	CurrentDura    uint16
	MaxDura        uint16
	Count          uint32
	GemCount       uint32
	AC             byte
	MAC            byte
	DC             byte
	MC             byte
	SC             byte
	Accuracy       byte
	Agility        byte
	HP             byte
	MP             byte
	Strong         byte
	MagicResist    byte
	PoisonResist   byte
	HealthRecovery byte
	ManaRecovery   byte
	PoisonRecovery byte
	CriticalRate   byte
	CriticalDamage byte
	Freezing       byte
	PoisonAttack   byte
	AttackSpeed    byte
	Luck           byte
	RefinedValue   RefinedValue
	RefineAdded    byte
	DuraChanged    bool
	SoulBoundId    uint32
	Identified     bool
	Cursed         bool
	WeddingRing    uint32
	//public UserItem[] Slots = new UserItem[5];
	//public DateTime BuybackExpiryDate;
	//public ExpireInfo ExpireInfo;
	//public RentalInformation RentalInformation;
	//public Awake Awake = new Awake();
}

// TODO
func (self *UserItem) ToBytes() []byte {
	return nil
}
