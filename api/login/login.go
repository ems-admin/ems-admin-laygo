package login

import (
	"ems-adming-go/internal/utils/captchauitl"
	"ems-adming-go/internal/utils/responseutil"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Uuid     string `json:"uuid" binding:"required"`
}

func HandleLogin(context *gin.Context) {
	var loginRequest Request
	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		context.JSON(http.StatusOK, responseutil.FailResponse("必填字段不能为空"))
		return
	}
	//	校验验证码
	stringCaptcha := captchauitl.NewStringCaptcha()
	if !stringCaptcha.Verify(loginRequest.Uuid, loginRequest.Code) {
		context.JSON(http.StatusOK, responseutil.FailResponse("验证码错误"))
		return
	}
	if loginRequest.Username == "admin" && loginRequest.Password == "123456" {
		context.JSON(http.StatusOK, responseutil.SuccessResponse("登录成功"))
		return
	} else {
		context.JSON(http.StatusOK, responseutil.FailResponse("账号或密码错误"))
		return
	}
}
