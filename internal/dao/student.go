package dao

import "github.com/hjd919/server/internal/model"

func (d *Dao) Test() string {
	var a model.Aaa
	d.db.First(&a)
	return a.Appid
}
