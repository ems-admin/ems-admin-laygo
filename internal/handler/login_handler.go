package handler

import (
	"ems-adming-go/internal/config"
	"ems-adming-go/internal/model"
	"ems-adming-go/internal/service"
	"ems-adming-go/pkg/utils/captchauitl"
	"ems-adming-go/pkg/utils/responseutil"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type LoginHandler struct {
	Service *service.LoginService
}

func (h LoginHandler) RenderLogin(context *gin.Context) {
	//	验证码对象
	stringCaptcha := captchauitl.NewStringCaptcha()
	id, b64s, answer := stringCaptcha.Generate()
	context.HTML(http.StatusOK, "login.html", gin.H{
		"captchaId": id,
		"answer":    answer,
		"b64s":      template.URL(b64s),
	})
}

func (h LoginHandler) GetCaptcha(c *gin.Context) {
	//	验证码对象
	stringCaptcha := captchauitl.NewStringCaptcha()
	id, b64s, answer := stringCaptcha.Generate()
	c.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"answer":    answer,
		"b64s":      template.URL(b64s),
	})
}

func (h LoginHandler) HandleLogin(context *gin.Context) {
	var loginRequest model.Login
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
	//	获取当前用户
	sysUser, err := service.FindByName(loginRequest.Username)
	if err != nil {
		context.JSON(http.StatusOK, responseutil.FailResponse(err.Error()))
	}

	if sysUser != nil {
		err = bcrypt.CompareHashAndPassword([]byte(sysUser.Password), []byte(loginRequest.Password))
		if err != nil {
			context.JSON(http.StatusOK, responseutil.FailResponse("账号或密码错误"))
			return
		}
		tokenString, err := config.GenerateJWT(*sysUser)
		if err != nil {
			context.JSON(http.StatusOK, responseutil.FailResponse(err.Error()))
		}
		result := make(map[string]interface{})
		result["token"] = tokenString
		result["username"] = sysUser.Username
		result["nickName"] = sysUser.NickName
		context.JSON(http.StatusOK, responseutil.SuccessResponse(result))
		return
	} else {
		context.JSON(http.StatusOK, responseutil.FailResponse("账号或密码错误"))
		return
	}
}
