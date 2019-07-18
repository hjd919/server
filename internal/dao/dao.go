package dao

import (
	"github.com/hjd919/server/pkg/database"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/jinzhu/gorm"
)

// Dao Dao.
type Dao struct {
	db *gorm.DB
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New new a Dao and return.
func New() *Dao {
	var (
		dc struct {
			Demo *database.MySQLConfig
		}
	)
	checkErr(paladin.Get("mysql.toml").UnmarshalTOML(&dc))
	return &Dao{
		// mysql
		db: database.NewMySQL(dc.Demo),
	}
}
