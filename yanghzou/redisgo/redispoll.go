package redisgo

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"yanghzou/utill"
)

var pool *redis.Pool
//redis连接池,
func init() {
	pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   1,
		IdleTimeout: utill.ReadTimeout,
		Dial: func() (redis.Conn, error) {
			url := fmt.Sprintf("%s:%d", utill.RedisHost, utill.RedisPort)
			return redis.Dial("tcp",url,redis.DialDatabase(1),redis.DialPassword(utill.RedisAuth))
		},
	}
}

func GetPoolRecords()  {
	c := pool.Get()
	defer c.Close()
	res,err := redis.String(c.Do("get","name"))
	if err != nil {
		fmt.Println("get result right:",err)
		return
	}
	fmt.Println("get successful",res)
	pool.Close()
}
