package redis

import "github.com/gomodule/redigo/redis"

var RedisConn redis.Conn

func NewRedisConn() (redis.Conn, error) {
	conn, err := redis.Dial("tcp", ":6379")
	return conn, err
}

func init() {
	var err error
	RedisConn, err = NewRedisConn()
	if err != nil {
		panic("init mysql error")
	}
}
