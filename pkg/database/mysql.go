package database

// http://gorm.book.jasperxu.com/database.html#dbc

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
	MaxOpen int      // pool
	MaxIdle int      // pool
	Debug   bool
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *MySQLConfig) (db *gorm.DB) {
	var err error
	db, err = gorm.Open("mysql", c.DSN)

	if err != nil {
		log.Error("models.Setup err: %v", err)
	}

	db.LogMode(c.Debug)

	db.SingularTable(true)

	db.DB().SetMaxIdleConns(c.MaxIdle)
	db.DB().SetMaxOpenConns(c.MaxOpen)
	return
}

func MysqlConn() (client *gorm.DB) {
	c := &MySQLConfig{
		DSN:     "jdcj:Jdcjxiaozi527@tcp(39.96.187.72:33066)/post?charset=utf8&parseTime=True&loc=Local",
		MaxOpen: 500,
		MaxIdle: 10,
		Debug:   true,
	}
	client = NewMySQL(c)
	return
}
