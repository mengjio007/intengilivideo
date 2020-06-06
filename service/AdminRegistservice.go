package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)
//管理员手工登陆
type AdminRe struct{
	Adminname string `json:"adminname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Passwordtwo string `json:"passwordtwo" binding:"required"`
}

func(ser *AdminRe)Regist() *serializer.Response{
	if ser.Passwordtwo != ser.Password{
		return &serializer.Response{
			Status: 40001,
			Msg:    "两次输入的密码不同，请重新输入",
			Error:  "",
		}
	}
	count := 0
	model.DB.Model(&model.Admin{}).Where("admin_name = ?",ser.Adminname).Count(&count)
	if count>0{
		return &serializer.Response{
			Status: 40001,
			Msg:    "用户已被注册",
			Error:  "",
		}
	}
	return nil
}
func(regist *AdminRe)Keep() (model.Admin , *serializer.Response){
	admin := model.Admin{
		AdminName:regist.Adminname,
	}

	//密码前后校验
	if err := regist.Regist(); err!=nil{
		return admin,err
	}
	if err := admin.SetPassword(regist.Password);err!=nil{
		return admin,&serializer.Response{
			Status: 40002,
			Msg:    "密码加密失败",
		}
	}
	if err := model.DB.Create(&admin).Error; err != nil{
		return	admin , &serializer.Response{
			Status: 40002,
			Msg:    "注册失败",
			Error:  err.Error(),
		}
	}

	return admin , nil
}