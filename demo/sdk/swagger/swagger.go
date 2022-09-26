package swagger

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// docs "github.com/go-project-name/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Sw struct{}

func (s Sw) Desc() string {
	return "swagger接口文档示例"
}

func (s Sw) Name() string {
	return "swagger"
}

func (s Sw) ReadMe() {
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

func (s Sw) DemoOfSwagger() {
	r := gin.Default()
	// docs.SwaggerInfo.BasePath = "/api/v1"
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
