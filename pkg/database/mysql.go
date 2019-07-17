package database

import (
	"github.com/bilibili/kratos/pkg/log"
	"github.com/jinzhu/gorm"

	// database driver
	_ "github.com/go-sql-driver/mysql"
)

// MySQLConfig mysql config.
type MySQLConfig struct {
	Addr    string   // for trace
	DSN     string   // write data source name.
	ReadDSN []string // read data source name.
	Active  int      // pool
	Idle    int      // pool
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *MySQLConfig) (db *gorm.DB) {
	var err error
	db, err = gorm.Open("mysql", c.DSN)

	if err != nil {
		log.Error("models.Setup err: %v", err)
	}

	db.LogMode(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return
}
