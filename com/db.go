package com

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func GetDB() *gorm.DB {
	db, err := gorm.Open(dialect, dbUser+":"+dbPassword+"@tcp("+dbAddr+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

type AccountInfo struct {
	Index     uint32 `gorm:"primary_key"`
	AccountID string // 这个才是登陆用户名...........WTF???
	Password  string

	CharacterInfos []CharacterInfo
}

type CharacterInfo struct {
	Index         uint32 `gorm:"primary_key"`
	AccountInfoID uint32 // AcountInfo.AccountID
	Name          string
	Level         uint16
	Class         byte
	Gender        byte
	Hair          byte
	GuildIndex    int32
	CreationIP    string

	// TODO
	//public bool Banned;
	//public string BanReason = string.Empty;
	//public DateTime ExpiryDate;
	// ...
	CurrentMapIndex  uint32
	CurrentLocationX int32
	CurrentLocationY int32
	Direction        MirDirection
	//public int BindMapIndex;
	//public Point BindLocation;

	HP         uint16
	MP         uint16
	Experience uint64

	//public AttackMode AMode;
	//public PetMode PMode;
	// ...
}

// many maps to many monster types
type RespawnInfo struct {
	Index        uint32 `gorm:"primary_key"`
	MapIndex     uint32
	MonsterIndex uint32
	Count        uint32
	Spread       uint32
	LocationX    int32
	LocationY    int32
}

// monster type
type MonsterInfo struct {
	Index        uint32 `gorm:"primary_key"`
	MonsterIndex uint32
	Name         string
	Level        uint16
	HP           uint16
	Experience   uint32
	// TODO ...
	// uint16 MinAC, MaxAC, MinMAC, MaxMAC, MinDC, MaxDC, MinMC, MaxMC, MinSC, MaxSC;
}

type NPCInfo struct {
	Index     uint32 `gorm:"primary_key"`
	Name      string
	MapIndex  uint32
	LocationX int32
	LocationY int32
	Image     uint16
}
