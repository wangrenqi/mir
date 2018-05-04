package env

import (
	"mir/orm"
	"github.com/jinzhu/gorm"
)

type Environ struct {
	Db   *gorm.DB
	Maps *map[uint16]Map
}

func InitEnviron() *Environ {
	db := orm.GetDB()

	db.AutoMigrate(&orm.AccountInfo{}, &orm.CharacterInfo{})

	maps := GetMaps(MapFilesPath)

	return &Environ{
		Db:   db,
		Maps: maps,
	}
}
