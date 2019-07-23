package dao

import (
	"github.com/hjd919/server/pkg/database"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/jinzhu/gorm"
)

// Dao Dao.
type Dao struct {
	DB  *gorm.DB
	MDB *mongo.Database
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
		dm struct {
			Mongodb *database.MDBConfig
		}
	)

	checkErr(paladin.Get("mysql.toml").UnmarshalTOML(&dc))
	checkErr(paladin.Get("mongodb.toml").UnmarshalTOML(&dm))
	return &Dao{
		// mysql
		DB:  database.NewMySQL(dc.Demo),
		MDB: database.NewMDB(dm.Mongodb),
	}
}
