package middleware

import (
	"fmt"
    "github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv12/global"
	"github.com/liuhongdi/digv12/pkg/result"
)

func PermissionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 请求的path
		p := c.Request.URL.Path
		// 请求的方法
		m := c.Request.Method

		//role:="superAdmin"
		role:="user"
		//role:="guest"

		fmt.Println("role:"+role)
		fmt.Println("path:"+p)
		fmt.Println("method:"+m)

		// 检查用户权限
		isPass, err := global.Enforcer.Enforce(role, p, m)
		if err != nil {
			resultRes := result.NewResult(c)
			resultRes.Error(2005,err.Error())
			return
		}
		if isPass {
			c.Next()
		} else {
			resultRes := result.NewResult(c)
			resultRes.Error(2006,"无访问权限")
			return
		}
	}
}
