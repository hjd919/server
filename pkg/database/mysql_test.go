package database

import (
	// database driver
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 链接
func TestMysqlConn(t *testing.T) {
	db := MysqlConn()
	defer db.Close()
}

type Post struct {
	ID        int `gorm:"primary_key"`
	Num       int
	Name      string
	CreatedAt time.Time
}

// 增加
func TestMysqlAdd(t *testing.T) {
	db := MysqlConn()
	defer db.Close()
	var err error
	// var post Post
	post := Post{
		Num:  22,
		Name: "hjd",
	}
	err = db.Create(&post).Error
	if err != nil {
		log.Printf("TestMysqlAdd-Create-%v", err.Error())
		return
	}
	log.Printf("TestMysqlAdd-Create-%v", post)
}

/*
// 获取第一个匹配记录
db.Where("name = ?", "jinzhu").First(&user)
//// SELECT * FROM users WHERE name = 'jinzhu' limit 1;

// 获取所有匹配记录
db.Where("name = ?", "jinzhu").Find(&users)
//// SELECT * FROM users WHERE name = 'jinzhu';

db.Where("name <> ?", "jinzhu").Find(&users)

// IN
db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)

// LIKE
db.Where("name LIKE ?", "%jin%").Find(&users)

// AND
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)

// Time
db.Where("updated_at > ?", lastWeek).Find(&users)

db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
*/

// 查询-单条
func TestMysqlQuery(t *testing.T) {
	db := MysqlConn()
	defer db.Close()
	var err error
	// where struct
	where := Post{
		Name: "hjd",
	}
	var res Post
	dbwhere := db.Where(where)
	err = dbwhere.First(&res).Error
	if err != nil {
		log.Printf("TestMysqlQuery-First-%v", err.Error())
		return
	}
	log.Printf("TestMysqlQuery-First-%v", res)
	// where map
	where3 := map[string]interface{}{
		"id": "1",
	}
	var res3 Post
	dbwhere3 := db.Where(where3)
	err = dbwhere3.First(&res3).Error
	if err != nil {
		log.Printf("TestMysqlQuery-First-%v", err.Error())
		return
	}
	log.Printf("TestMysqlQuery-First-%v", res3)
	// where query+args
	query := "id = ?"
	args := []interface{}{2}
	var res4 Post
	dbwhere4 := db.Where(query, args...)
	err = dbwhere4.First(&res4).Error
	if err != nil {
		log.Printf("TestMysqlQuery-First-%v", err.Error())
		return
	}
	log.Printf("TestMysqlQuery-First-%v", res4)
}

// 查询-多条
func TestMysqlQueryMany(t *testing.T) {
	db := MysqlConn()
	defer db.Close()
	var err error
	var posts []Post
	// where := map[string]interface{}{}
	query := "id > ?"
	args := []interface{}{2}
	dbwhere := db.Where(query, args...)
	err = dbwhere.Find(&posts).Error
	if err != nil {
		log.Printf("TestMysqlQueryMany-Find-%v", err.Error())
		return
	}
	log.Printf("TestMysqlQueryMany-Find-%v", posts)
}

// 删除
func TestMysqlDelete(t *testing.T) {
	db := MysqlConn()
	defer db.Close()
	var err error
	var post Post
	// where := map[string]interface{}{}
	query := "id > ?"
	args := []interface{}{2}
	dbwhere := db.Where(query, args...)
	err = dbwhere.Delete(&post).Error
	if err != nil {
		log.Printf("TestMysqlDelete-Delete-%v", err.Error())
		return
	}
	log.Printf("TestMysqlDelete-Delete-%v", post)
}

// 更新
func TestMysqlUpdate(t *testing.T) {
	db := MysqlConn()
	defer db.Close()
	var err error
	var post Post
	// where := map[string]interface{}{}
	query := "id = ?"
	args := []interface{}{2}
	data := map[string]interface{}{
		"name": "xiaohuang2",
	}
	dbwhere := db.Model(post).Where(query, args...)
	err = dbwhere.UpdateColumns(data).Error
	if err != nil {
		log.Printf("TestMysqlUpdate-Update-%v", err.Error())
		return
	}
	log.Printf("TestMysqlUpdate-Update-%v", post)
}

// 自增
func TestMysqlUpdateIncr(t *testing.T) {
	db := MysqlConn()
	defer db.Close()
	var err error
	var post Post
	// where := map[string]interface{}{}
	query := "id = ?"
	args := []interface{}{2}
	data := map[string]interface{}{
		"num": gorm.Expr("num - ?", 1),
	}
	dbwhere := db.Model(post).Where(query, args...)
	err = dbwhere.Updates(data).Error
	// err = dbwhere.UpdateColumn("num", gorm.Expr("num - ?", 1)).Error
	if err != nil {
		log.Printf("TestMysqlUpdate-Update-%v", err.Error())
		return
	}
	log.Printf("TestMysqlUpdate-Update-%v", post)
}

// 原始查询-多个
func TestMysqlQueryManyRaw(t *testing.T) {
	db := MysqlConn()
	defer db.Close()
	var err error
	var posts []Post
	rows, err := db.Raw("select * from post where id>=?", 2).Rows() // (*sql.Rows, error)
	defer rows.Close()
	for rows.Next() {
		var post Post
		// name := ""
		// rows.Scan(&name) // 单个变量
		db.ScanRows(rows, &post) // 映射到结构体
		posts = append(posts, post)
	}
	if err != nil {
		log.Printf("TestMysqlQueryRaw-Update-%v", err.Error())
		return
	}
	log.Printf("TestMysqlQueryRaw-Update-%v", posts)
}

// 原始查询-单个
func TestMysqlQueryRaw(t *testing.T) {
	db := MysqlConn()
	defer db.Close()
	var err error
	var post Post
	err = db.Raw("SELECT * FROM post WHERE id = ?", 1).Scan(&post).Error
	if err != nil {
		log.Printf("TestMysqlQueryRaw-Update-%v", err.Error())
		return
	}
	log.Printf("TestMysqlQueryRaw-Update-%v", post)
}
