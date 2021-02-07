package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisPool *redis.Pool

func InitRedisPool(addr, key string) {
	RedisPool = GenRedisPool(addr, key)
}

func GenRedisPool(addr, key string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     100,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if key == "" {
				return c, nil
			}
			if _, err := c.Do("AUTH", key); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
