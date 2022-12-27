package ginsdk

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (Gin) Src() {
	g := gin.Default()
	if err := g.Run(); err != nil {
		fmt.Println(err)
	}
}
