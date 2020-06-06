package api

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
	"GiliVideo/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户注册
func UserRegist(c *gin.Context){
	service :=service.UserRegisterService{}

	if err := c.ShouldBindJSON(&service);err == nil {
		if user,err :=service.Regist();err!=nil{
			response:=serializer.UserResponse(user)
			c.JSON(http.StatusOK,response)
		}else {
			response := serializer.UserResponse(user)
			c.JSON(http.StatusOK,response)
		}
	}else{
		c.JSON(http.StatusOK,serializer.Response{
			Status: 50002,
			Msg:    "注册错误",
		})
	}

}

//用户登陆
func UserLogin(c *gin.Context){
	service := service.UserLoginService{}

	if err := c.ShouldBindJSON(&service);err == nil{
		if user,err := service.Login();err!=nil{
			c.JSON(http.StatusOK,err)
		}else {
			// 设置Session
			s := sessions.Default(c)
			s.Set("user_id",user.ID)
			s.Save()

			response := serializer.UserResponse(user)
			c.JSON(http.StatusOK,response)
		}
	}else{
		c.JSON(http.StatusOK,err)
	}
}

//修改密码
func ModifyPassWd(c *gin.Context){
 service:=service.UserModify{}
if err:=c.ShouldBind(&service);err!=nil{
	c.JSON(200,serializer.Response{
		Status: 50001,
		Data:   nil,
		Msg:    "数据绑定出错",
		Error:  "",
	})
}
	if res1:=service.Vaild();res1!=nil{
		c.JSON(200,res1)
	}else{
		res :=service.ChangePasswd()
		c.JSON(200,res)
	}

}
func ModiInfo(c *gin.Context){
	service:=service.UserMemodify{}
	if err := c.ShouldBind(&service);err!=nil{
		c.JSON(http.StatusOK,serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "数据绑定错误",
			Error:  "",
		})
	}else{
		user,_:=c.Get("user")
		uid := user.(*model.User).ID
		res:=service.INfo(uid)
		c.JSON(http.StatusOK,res)
	}
}
//登出，删除session
func Logout(c *gin.Context){
	s:=sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(http.StatusOK,serializer.Response{
		Status: 200,
		Data:   nil,
		Msg:    "登出",
		Error:  "",
	})
}

//获得上传token
func UpAvatartoken(c *gin.Context){
	service := service.UploadUserTokenService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Status: 50000,
			Data:   nil,
			Msg:    "",
			Error:  "",
		})
	}
}

//上传头像
func Upavatar(c *gin.Context){
	service:=service.AvatarUser{}
	if err:=c.ShouldBind(&service);err ==nil{
		if user, ok := c.Get("user");ok{
			uid := user.(*model.User).ID
			res := service.Upavatar(uid)
			c.JSON(http.StatusOK,res)
		}else {
			c.JSON(200, serializer.Response{
				Status: 50000,
				Data:   nil,
				Msg:    "获取用户失败",
				Error:  "",
			})
		}
	}else{
		c.JSON(200,serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "数据绑定失败",
			Error:  "",
		})
	}
}

//获得用户信息
func InfoofUser(c *gin.Context){
	service:=service.UserInfo{}
	if err:=c.ShouldBind(&service);err==nil{
		if user,ok := c.Get("user"); ok{
			uid := user.(*model.User).ID
		res:=service.GetAllInfo(uid)
		c.JSON(http.StatusOK,res)
	}else{
		c.JSON(http.StatusOK,serializer.Response{
			Status: 50000,
			Data:   nil,
			Msg:    "获取用户失败",
			Error:  "",
		})
	}}else{
		c.JSON(http.StatusOK,serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "数据绑定失败",
			Error:  "",
		})
	}
}

//获得用户上传的视频
func Uservideo(c *gin.Context){
	service:=service.ListUservideoservice{}
	if err:=c.ShouldBind(&service);err!=nil{
		c.JSON(http.StatusOK,serializer.Response{
			Status: 40001,
			Data:   nil,
			Msg:    "数据绑定错误",
			Error:  "",
		})
	}else{
		if user,ok := c.Get("user"); ok{
			uid := user.(*model.User).ID
			res :=service.List(uid)
			c.JSON(http.StatusOK,res)
		}else{
			c.JSON(http.StatusOK,serializer.Response{
				Status:50001,
				Msg:"登陆获取失败",
			})
		}
	}
}

func DeleteVideo( c  *gin.Context){
	service:=service.Deletevideo{}
	if err:=c.ShouldBind(&service);err!=nil{
		c.JSON(http.StatusOK,serializer.Response{
			Status: 40001,
			Data:   nil,
			Msg:    "数据绑定错误",
			Error:  "",
		})
	}else{
		res :=service.Delete()
		c.JSON(http.StatusOK,res)
	}
		}