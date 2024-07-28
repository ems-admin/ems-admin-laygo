package handler

import (
	"ems-adming-go/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes() *gin.Engine {

	router := gin.Default()
	//	配置静态资源目录
	router.Static("/static", "./web/static")
	//	定义 HTML 模板目录
	router.LoadHTMLGlob("web/templates/*")
	//	定义默认页面
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello World",
		})
	})

	//	登录相关路由
	loginService := &service.LoginService{}
	LoginRoutes(router, loginService)

	return router
}
