package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

// UserRegisterService 用户注册服务
type UserRegisterService struct {
	UserName        string `form:"username" json:"username" binding:"required,min=2,max=20"`
	Email			string `form:"email" json:"email" `
	Password        string `form:"password" json:"password" binding:"required,min=6,max=30"`
	PasswordTwo 	string `form:"passwordtwo" json:"passwordtwo" binding:"required,min=6,max=30"`
	Brithday        string `form:"brithday" json:"brithday" `
	Gender          string `form:"gender" json:"gender" binding:"required"`
}

func (regist *UserRegisterService)Valid() *serializer.Response{
	if regist.PasswordTwo != regist.Password{
		return &serializer.Response{
			Status: 40001,
			Msg:    "两次输入的密码不同，请重新输入",
			Error:  "",
		}
	}
	count := 0
	model.DB.Model(&model.User{}).Where("user_name = ?",regist.UserName).Count(&count)
	if count>0{
		return &serializer.Response{
			Status: 40001,
			Msg:    "用户已被注册",
			Error:  "",
		}
	}
	return nil
}

func(regist *UserRegisterService)Regist() (model.User , *serializer.Response){
	user := model.User{
		UserName:regist.UserName,
		Email:regist.Email,
		Brithday:regist.Brithday,
		Gender:regist.Gender,
		Status:model.Active,
	}

	//密码前后校验
	if err := regist.Valid(); err!=nil{
		return user,err
	}
	if err := user.SetPassword(regist.Password);err!=nil{
		return user,&serializer.Response{
			Status: 40002,
			Msg:    "密码加密失败",
		}
	}
	if err := model.DB.Create(&user).Error; err != nil{
		return	user , &serializer.Response{
			Status: 40002,
			Msg:    "注册失败",
			Error:  err.Error(),
		}
	}

	return user , nil
}
