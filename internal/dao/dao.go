package dao

import (
	"github.com/hjd919/server/pkg/database"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/jinzhu/gorm"
)

// // Dao Dao interface
// type Dao interface {
// 	Close()
// 	// Ping(ctx context.Context) (err error)
// }

// Dao Dao.
type Dao struct {
	db *gorm.DB
	// redis       *redis.Pool
	// redisExpire int32
	// mc       *memcache.Memcache
	// mcExpire int32
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
		// rc struct {
		// 	Demo       *redis.Config
		// 	DemoExpire xtime.Duration
		// }
		// mc struct {
		// 	Demo       *memcache.Config
		// 	DemoExpire xtime.Duration
		// }
	)
	checkErr(paladin.Get("mysql.toml").UnmarshalTOML(&dc))
	// checkErr(paladin.Get("redis.toml").UnmarshalTOML(&rc))
	// checkErr(paladin.Get("memcache.toml").UnmarshalTOML(&mc))
	return &Dao{
		// mysql
		db: database.NewMySQL(dc.Demo),
		// redis
		// redis:       redis.NewPool(rc.Demo),
		// redisExpire: int32(time.Duration(rc.DemoExpire) / time.Second),
		// memcache
		// mc:       memcache.New(mc.Demo),
		// mcExpire: int32(time.Duration(mc.DemoExpire) / time.Second),
	}
}

// Close close the resource.
func (d *Dao) Close() {
	// d.mc.Close()
	// d.redis.Close()
	d.db.Close()
}

// Ping ping the resource.
// func (d *Dao) Ping(ctx context.Context) (err error) {
// 	if err = d.pingMC(ctx); err != nil {
// 		return
// 	}
// 	if err = d.pingRedis(ctx); err != nil {
// 		return
// 	}
// 	return d.db.Ping(ctx)
// }

// func (d *Dao) pingMC(ctx context.Context) (err error) {
// 	if err = d.mc.Set(ctx, &memcache.Item{Key: "ping", Value: []byte("pong"), Expiration: 0}); err != nil {
// 		log.Error("conn.Set(PING) error(%v)", err)
// 	}
// 	return
// }

// func (d *Dao) pingRedis(ctx context.Context) (err error) {
// 	conn := d.redis.Get(ctx)
// 	defer conn.Close()
// 	if _, err = conn.Do("SET", "ping", "pong"); err != nil {
// 		log.Error("conn.Set(PING) error(%v)", err)
// 	}
// 	return
// }
