package main

import (
	"ems-adming-go/internal/utils/captchauitl"
	"ems-adming-go/routes"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes(router)

	router.Static("/static", "./web/static")
	router.LoadHTMLGlob("web/templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello World",
		})
	})

	stringCaptcha := captchauitl.NewStringCaptcha()

	router.GET("/login.html", func(c *gin.Context) {
		id, b64s, answer := stringCaptcha.Generate()
		c.HTML(http.StatusOK, "login.html", gin.H{
			"captchaId": id,
			"answer":    answer,
			"b64s":      template.URL(b64s),
		})
	})

	router.GET("/captcha", func(c *gin.Context) {
		id, b64s, answer := stringCaptcha.Generate()
		c.JSON(http.StatusOK, gin.H{
			"captchaId": id,
			"answer":    answer,
			"b64s":      template.URL(b64s),
		})
	})

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
