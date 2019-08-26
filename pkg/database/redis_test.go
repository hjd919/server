package database

import (
	// database driver

	"fmt"
	"log"
	"testing"

	"github.com/gomodule/redigo/redis"
)

// 链接
func TestRedisConn(t *testing.T) {
	pool := RedisPoolGet()
	redis := pool.Get()
	defer redis.Close()
}

type RedisData struct {
	Title  string `redis:"title"`
	Author string `redis:"author"`
	Body   string `redis:"body"`
}

// string
func TestStringSET(t *testing.T) {
	pool := RedisPoolGet()
	c := pool.Get()
	defer c.Close()
	// del key
	if _, err := c.Do("del", "stringkey"); err != nil {
		log.Println(err)
		return
	}
	// set key
	res, _ := redis.String(c.Do("set", "stringkey", "333"))
	log.Printf("%#v", res)
	// get int
	intval, _ := redis.Int(c.Do("get", "stringkey"))
	log.Printf("%#v", intval)
	// get str
	strval, _ := redis.String(c.Do("get", "stringkey"))
	log.Printf("%#v", strval)
}
func TestStringGET(t *testing.T) {
	pool := RedisPoolGet()
	c := pool.Get()
	defer c.Close()
	// get int
	intval, _ := redis.Int(c.Do("get", "stringkey"))
	log.Printf("%#v", intval)
}

// setex expire
func TestStringSETEX(t *testing.T) {
	pool := RedisPoolGet()
	c := pool.Get()
	defer c.Close()
	resBool, _ := redis.Bool(c.Do("exists", "stringkey"))
	if resBool {
		log.Println("存在", resBool)
		// del key
		if _, err := c.Do("del", "stringkey"); err != nil {
			log.Println(err)
			return
		}
	} else {
		log.Println("不存在", resBool)
	}
	// setex key time val
	resStr, _ := redis.String(c.Do("setex", "stringkey", 10, "hello"))
	log.Printf("%#v", resStr)
	// get int
	intval, _ := redis.Int(c.Do("get", "stringkey"))
	log.Printf("%#v", intval)
	// get str
	strval, _ := redis.String(c.Do("get", "stringkey"))
	log.Printf("%#v", strval)
}

func TestStringExist(t *testing.T) {
	pool := RedisPoolGet()
	c := pool.Get()
	defer c.Close()
	resBool, _ := redis.Bool(c.Do("exists", "stringkey"))
	if resBool {
		log.Println("存在", resBool)
		// del key
		if _, err := c.Do("del", "stringkey"); err != nil {
			log.Println(err)
			return
		}
	} else {
		log.Println("不存在", resBool)
	}
}
func TestStringIncr(t *testing.T) {
	pool := RedisPoolGet()
	c := pool.Get()
	defer c.Close()
	resInt, _ := redis.Int(c.Do("incr", "stringkey"))
	log.Println("incr", resInt)
}

// list
// sort
// zsort
// hash
func TestHMSET(t *testing.T) {
	pool := RedisPoolGet()
	c := pool.Get()
	defer c.Close()
	var p1, p2 RedisData
	// data is struct
	p1.Title = "Example"
	p1.Author = "Gary"
	p1.Body = "Hello"
	if _, err := c.Do("hmset", redis.Args{}.Add("id1").AddFlat(&p1)...); err != nil {
		log.Println(err)
		return
	}
	// data is map
	m := map[string]string{
		"title":  "Example3",
		"author": "Steve",
		"body":   "Map",
	}
	if _, err := c.Do("hmset", redis.Args{}.Add("id2").AddFlat(m)...); err != nil {
		log.Println(err)
		return
	}
	for _, id := range []string{"id1", "id2"} {
		v, err := redis.Values(c.Do("hgetall", id))
		if err != nil {
			log.Println(err)
			return
		}

		if err := redis.ScanStruct(v, &p2); err != nil {
			log.Println(err)
			return
		}

		log.Printf("%+v\n", p2)
	}
}

// 单个结构体
func TestScanSlice(t *testing.T) {
	pool := RedisPoolGet()
	c := pool.Get()
	defer c.Close()
	c.Send("HMSET", "album:1", "title", "Red", "rating", 5)
	c.Send("HMSET", "album:2", "title", "Earthbound", "rating", 1)
	c.Send("HMSET", "album:3", "title", "Beat", "rating", 4)
	c.Send("LPUSH", "albums", "1")
	c.Send("LPUSH", "albums", "2")
	c.Send("LPUSH", "albums", "3")
	values, err := redis.Values(c.Do("SORT", "albums",
		"BY", "album:*->rating",
		"GET", "album:*->title",
		"GET", "album:*->rating"))
	if err != nil {
		fmt.Println(err)
		return
	}

	var albums []struct {
		Title  string
		Rating int
	}
	if err := redis.ScanSlice(values, &albums); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", albums)
}

// 单个值
func TestScan(t *testing.T) {
	pool := RedisPoolGet()
	c := pool.Get()
	defer c.Close()

	c.Send("HMSET", "album:1", "title", "Red", "rating", 5)
	c.Send("HMSET", "album:2", "title", "Earthbound", "rating", 1)
	c.Send("HMSET", "album:3", "title", "Beat")
	c.Send("LPUSH", "albums", "1")
	c.Send("LPUSH", "albums", "2")
	c.Send("LPUSH", "albums", "3")
	values, err := redis.Values(c.Do("SORT", "albums",
		"BY", "album:*->rating",
		"GET", "album:*->title",
		"GET", "album:*->rating"))
	if err != nil {
		fmt.Println(err)
		return
	}

	for len(values) > 0 {
		var title string
		rating := -1 // initialize to illegal value to detect nil.
		values, err = redis.Scan(values, &title, &rating)
		if err != nil {
			fmt.Println(err)
			return
		}
		if rating == -1 {
			fmt.Println(title, "not-rated")
		} else {
			fmt.Println(title, rating)
		}
	}
}
