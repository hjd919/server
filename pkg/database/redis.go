package database

// http://redisdoc.com/

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisConfig mysql config.
type RedisConfig struct {
	Addr     string // for trace
	Password string // write data source name.
	MaxOpen  int    // pool
	MaxIdle  int    // pool
	DB       int    // pool
	Debug    bool
}

// NewRedis new db and retry connection when has error.
func NewRedisPool(conf *RedisConfig) (pool *redis.Pool) {
	pool = &redis.Pool{
		MaxIdle:     conf.MaxIdle,
		MaxActive:   conf.MaxOpen,
		IdleTimeout: 240 * time.Second,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", conf.Addr, redis.DialPassword(conf.Password), redis.DialDatabase(conf.DB))
			if err != nil {
				return nil, err
			}
			return
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return
}

func RedisPoolGet() (pool *redis.Pool) {
	c := &RedisConfig{
		Addr:     "39.96.187.72:6379",
		Password: "Hjd123!@#",
		DB:       0,
		MaxOpen:  500,
		MaxIdle:  2,
		Debug:    true,
	}
	pool = NewRedisPool(c)
	return
}
