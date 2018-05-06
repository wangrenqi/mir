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

	for _, m := range *maps {
		m.LoadNPC()
		m.LoadMonster()
	}

	return &Environ{
		Db:   db,
		Maps: maps,
	}
}
