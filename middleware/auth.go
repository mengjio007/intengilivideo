package middleware

import (
	"GiliVideo/model"
	"GiliVideo/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.Response{
			Status: 401,
			Msg:    "需要登录",
		})
		c.Abort()
	}
}

// AdminCurrentUser 获取登录管理员
func AdminCurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		Adminid := session.Get("admin_id")
		if Adminid != nil {
			admin, err := model.GetUser(Adminid)
			if err == nil {
				c.Set("admin", &admin)
			}
		}
		c.Next()
	}
}
// AdminAuthRequired 需要登录
func AdminAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if admin, _ := c.Get("admin"); admin != nil {
			if _, ok := admin.(*model.Admin); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.Response{
			Status: 401,
			Msg:    "需要登录",
		})
		c.Abort()
	}
}