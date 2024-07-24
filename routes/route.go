package routes

import (
	"ems-adming-go/api/login"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	auth := r.Group("/auth")
	{
		auth.POST("/login", login.HandleLogin)
	}
}
