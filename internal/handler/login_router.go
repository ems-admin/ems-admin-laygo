package handler

import (
	"ems-adming-go/internal/service"
	"github.com/gin-gonic/gin"
)

func LoginRoutes(router *gin.Engine, loginService *service.LoginService) {
	loginHandler := &LoginHandler{loginService}
	router.GET("/login.html", loginHandler.RenderLogin)
	router.GET("/captcha", loginHandler.GetCaptcha)
	router.POST("/auth/login", loginHandler.HandleLogin)
}
