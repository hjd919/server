package database

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func RedisConn() (redisdb *redis.Client) {
	conf := &RedisConfig{
		Addr:     "39.96.187.72:6379",
		Password: "Hjd123!@#",
		DB:       3,
	}
	redisdb = NewRedis(conf)
	return
}

func TestNewRedis(t *testing.T) {
	redisdb := RedisConn()
	log.Printf("redisdb--%v", redisdb)
}

func TestRedisCurd(t *testing.T) {
	redisdb := RedisConn()
	err := redisdb.HMSet("hello", map[string]interface{}{"hello": 1, "nihao": "hjd"}).Err()
	if err != nil {
		panic(err)
	}

	// val, err := redisdb.Get("key2").Result()
	// if err != nil && err != errors.New("redis: nil") {
	// 	panic(err)
	// }
	// fmt.Println("key2--", val, reflect.TypeOf(val))
}

func TestConnectPool(t *testing.T) {
	client := RedisConn()
	time.Sleep(time.Duration(5) * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 200; j++ {
				client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
				client.Get(fmt.Sprintf("name%d", j)).Result()
			}
			log.Printf("PoolStats, TotalConns: %v", client.PoolStats())
		}()
	}

	wg.Wait()
}
