package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	cm "mir/common"
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
	Class      cm.MirClass
	Gender     cm.MirGender
	LastAccess int64

	AccountID uint
}

// TODO 这个方法不应该放在orm 而是proto
func (self *SelectInfo) ToBytes() []byte {
	// index(int32 4byte)
	indexBytes := cm.Uint32ToBytes(uint32(self.ID))
	// name (string)
	nameBytes := cm.StringToBytes(self.Name)
	// level (int16 2byte)
	levelBytes := cm.Uint16ToBytes(uint16(self.Level))
	// class (byte)
	class := self.Class
	classBytes := []byte{byte(class)}
	// gender (byte)
	gender := self.Gender
	genderBytes := []byte{byte(gender)}
	// lastAccess (int64 8byte)
	lastAccessBytes := cm.Uint64ToBytes(uint64(0))
	result := make([]byte, 0)
	for _, r := range [][]byte{indexBytes, nameBytes, levelBytes, classBytes, genderBytes, lastAccessBytes} {
		result = append(result, r...)
	}
	return result
}

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbAddr+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
