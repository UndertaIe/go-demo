package swagger

import (
	"fmt"
	"net/http"

	"github.com/UndertaIe/go-demo/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type S struct{}

func (s S) Desc() string {
	return "swagger接口文档示例"
}

func (s S) Name() string {
	return "swagger"
}

func (s S) ReadMe() {
	url := "https://github.com/swaggo/gin-swagger"
	fmt.Println(url)
}

//
// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// go install github.com/swaggo/swag/cmd/swag@latest
// swag init
// go run main.go swagger DemoOfSwagger
func (s S) DemoOfSwagger() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
