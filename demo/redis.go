package demo

import (
	"fmt"
	"go-demo/dbs"
)

func TestRedis() {
	resp, err := dbs.Conn.Do("set", "name", "mikasa")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
