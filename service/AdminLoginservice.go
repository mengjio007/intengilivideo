package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

//管理员登陆服务
type AdminLoginService struct{
	Adminname string `form:"adminname" json:"adminname" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required"`
}

func(service *AdminLoginService) Login() (model.Admin , *serializer.Response){
	var admin model.Admin
	//校验账号
	if err := model.DB.Where("admin_name= ?",service.Adminname).First(&admin).Error; err != nil{
		return admin, &serializer.Response{
			Status: 40001,
			Msg:    "账号错误",
		}
	}
	//校验密码
	if admin.CheckPassword(service.PassWord)==false{
		return admin,&serializer.Response{
			Status: 40002,
			Msg:    "密码错误",
			Error:  "",
		}
	}
	return admin,nil
}