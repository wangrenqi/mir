package env

import (
	"mir/orm"
	"github.com/jinzhu/gorm"
)

type Environ struct {
	Db   *gorm.DB
	Maps interface{}
}

func InitEnviron() *Environ {
	db := orm.GetDB()

	db.AutoMigrate(&orm.Account{}, &orm.SelectInfo{})

	// TODO maps

	return &Environ{
		Db: db,
		//Maps:
	}
}
