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
}

type CharacterInfo struct {
	gorm.Model
	Name   string
	Level  int
	Class  int
	Gender int
}

func CreateDB() (*gorm.DB) {
	db, err := gorm.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbAddr+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	return db
}
