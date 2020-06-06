package api

import (
	"GiliVideo/serializer"
	service2 "GiliVideo/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

//展示视频
func LiVideo(c *gin.Context){
	service:=service2.ListVideos{}
	if err := c.ShouldBind(&service);err!=nil{
		c.JSON(http.StatusOK,serializer.Response{
			Status:50001,
			Msg:"数据绑定错误",
		})
	}else{
		res:=service.Videos()
		c.JSON(http.StatusOK,res)
	}
}

//展示用户
func LiUser(c *gin.Context){
	service:=service2.LiUsers{}
	if err := c.ShouldBind(&service);err!=nil{
		c.JSON(http.StatusOK,serializer.Response{
			Status:50001,
			Msg:"数据绑定错误",
		})
	}else{
		res:=service.Users()
		c.JSON(http.StatusOK,res)
	}
}

//展示评论
func LiComment(c *gin.Context){
	service:=service2.Licomment{}
	if err := c.ShouldBind(&service);err!=nil{
		c.JSON(http.StatusOK,serializer.Response{
			Status:50001,
			Msg:"数据绑定错误",
		})
	}else{
		res:=service.Comments()
		c.JSON(http.StatusOK,res)
	}
}

//管理员登陆
func AdLogin(c *gin.Context){
	service:=service2.AdminLoginService{}
	if err := c.ShouldBind(&service);err !=nil{
		c.JSON(http.StatusOK,serializer.Response{
			Status:50000,
			Msg:"json结构体绑定错误",
			})
	}else{
		if admin,err:=service.Login();err!=nil{
			c.JSON(http.StatusOK,err)
		}else{
			// 设置Session
			s := sessions.Default(c)
			s.Set("admin_id",admin.ID)
			s.Save()

			response := serializer.AdminResponse(admin)
			c.JSON(http.StatusOK,response)
		}
	}
}
//管理员注册
func AdRegist(c *gin.Context){
	service:=service2.AdminRe{}
	if err := c.ShouldBind(&service);err !=nil{
		c.JSON(http.StatusOK,serializer.Response{
			Status:50000,
			Msg:"json结构体绑定错误",
		})
	}else{
		if res ,err:=service.Keep();err!=nil{
			c.JSON(http.StatusOK,err)
		}else{
			c.JSON(http.StatusOK,serializer.AdminResponse(res))
		}
	}
}