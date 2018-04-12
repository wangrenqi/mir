package env

import (
	"mir-go/orm"
	"github.com/jinzhu/gorm"
)

type Environ struct {
	Db   *gorm.DB
	Maps interface{}
}

func InitEnviron() *Environ {
	db := orm.GetDB()
	defer db.Close()

	// TODO maps

	return &Environ{
		Db: db,
		//Maps:
	}
}
