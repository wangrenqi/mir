package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dialect = "mysql"

var dbName = "mir"
var dbUser = "root"
var dbPassword = "root"
var dbAddr = "localhost:3306"

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbAddr+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

type AccountInfo struct {
	Index     uint32 `gorm:"primary_key"`
	AccountID string
	Password  string
	UserName  string
}

type CharacterInfo struct {
	Index         uint32 `gorm:"primary_key"`
	AccountInfoID uint
	Name          string
	Level         uint16
	Class         byte
	Gender        byte
	Hair          byte
	GuildIndex    int32
	CreationIP    string
	//public bool Banned;
	//public string BanReason = string.Empty;
	//public DateTime ExpiryDate;
	// ...
	//public MirDirection Direction;
	//public int BindMapIndex;
	//public Point BindLocation;

	HP         uint16
	MP         uint16
	Experience uint64

	//public AttackMode AMode;
	//public PetMode PMode;
	// ...
}
