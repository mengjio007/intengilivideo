package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)


//用户登陆服务
type UserLoginService struct{
	Username string `form:"username" json:"username" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required"`
	Status string `form:"status" json:"status"`
}

func(service *UserLoginService) Login() (model.User , *serializer.Response){
	var user model.User
	//校验账号
	if err := model.DB.Where("user_name= ?",service.Username).First(&user).Error; err != nil{
		return user, &serializer.Response{
			Status: 40001,
			Msg:    "账号错误",
		}
	}
	//校验密码
	if user.CheckPassword(service.PassWord)==false{
		return user,&serializer.Response{
			Status: 40002,
			Msg:    "密码错误",
			Error:  "",
		}
	}
	return user,nil
}