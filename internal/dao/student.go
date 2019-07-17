package dao

import "github.com/hjd919/server/internal/model"

func (d *Dao) Test() string {
	// d.mc.Close()
	// d.redis.Close()
	// d.db.Where()
	// d.db.Find()
	var a model.Aaa
	d.db.First(&a)
	return a.Appid
}
