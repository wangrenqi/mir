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

type Account struct {
	gorm.Model
	//AccountID string
	Password string
	UserName string

	Characters []SelectInfo
}

type SelectInfo struct {
	gorm.Model
	//Index      int32 `gorm:"primary_key"`  replace as gorm.Model Id
	Name       string
	Level      int16
	Class      byte
	Gender     byte
	LastAccess int64

	AccountID uint
}

func (self *SelectInfo) ToBytes() []byte {
	// TODO 把select info (character info)转成bytes
	return make([]byte, 0)
}

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbAddr+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
