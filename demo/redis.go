package demo

import (
	"fmt"
	"go-demo/dbs"

	"github.com/UndertaIe/passwd/pkg/cache"
)

func TestRedis() {
	resp, err := dbs.Conn.Do("set", "name", "mikasa")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func cc() map[string]interface{} {
	var m map[string]interface{}
	return m
}

func TestPkgCache() {
	c, _ := cache.NewCache(cache.RedisT, nil)
	err := c.Set("lang", "Go", cache.FOREVER)
	fmt.Println(err)

	var s *interface{}
	err = c.Get("lang", s)
	fmt.Println(s, err)

}
