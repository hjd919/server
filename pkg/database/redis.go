package database

import (
	"log"

	"github.com/go-redis/redis"
	// database driver
)

// RedisConfig Redis config.
type RedisConfig struct {
	Addr     string // for trace
	Password string // pool
	DB       int    // pool
}

// NewRedis new db and retry connection when has error.
func NewRedis(c *RedisConfig) (redisdb *redis.Client) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password, // no password set
		DB:       c.DB,       // use default DB
		// PoolSize: 10,
	})

	pong, err := redisdb.Ping().Result()
	log.Println(pong, err)
	// Output: PONG <nil>
	return
}
