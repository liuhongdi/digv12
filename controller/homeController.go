package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv12/pkg/result"
)

type HomeController struct{}

func NewHomeController() HomeController {
	return HomeController{}
}
//首页，任何人可访问，不登录也可访问
func (a *HomeController) Home(c *gin.Context) {
	resultRes := result.NewResult(c)
	resultRes.Success("this is home page");
	return
}
