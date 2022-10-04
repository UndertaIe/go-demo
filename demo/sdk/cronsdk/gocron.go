package cronsdk

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
)

var channel = make(chan int, 10)

var x = 0

var get = func(ch chan int) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := <-ch
		fmt.Fprintln(c.Writer, i)
	}
}
var update = func(ch chan int) func() {
	return func() {
		ch <- x
		x += 1
	}
}

func (c C) Bug() {
	router := gin.Default()
	router.GET("/get", get(channel))
	fmt.Println(1)
	router.GET("/ok", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})

	gocron.Every(5).Second().Do(update(channel))
	stop := <-gocron.Start()
	router.Run(":11111")
	fmt.Println(stop)
	fmt.Println(1)
} // gocron and http server will block each other in main goroutine

// run gocron in other goroutine
// https://stackoverflow.com/questions/52719015/using-gin-gonic-and-some-scheduler-in-golang
func (c C) Fix() {

	router := gin.Default()
	router.GET("/get", get(channel))
	fmt.Println(1)
	router.GET("/ok", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})

	go func() {
		gocron.Every(3).Second().Do(update(channel))
		stop := <-gocron.Start()
		fmt.Println(stop)
	}()
	router.Run(":11111")
	fmt.Println(1)
}
