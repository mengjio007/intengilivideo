package api

import (
	"GiliVideo/serializer"
	"GiliVideo/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
//
//func GetSpace(){
//
//}
//

//func ModifyInfo(){
//
//}
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