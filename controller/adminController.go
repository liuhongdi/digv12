package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv12/pkg/result"
)

type AdminController struct{}

func NewAdminController() AdminController {
	return AdminController{}
}
//admin的管理页面，只有superAdmin角色的用户可以访问
func (a *AdminController) Admin(c *gin.Context) {
	resultRes := result.NewResult(c)
	resultRes.Success("this is admin's home page");
	return
}