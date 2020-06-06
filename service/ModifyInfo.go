package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)


//更改密码
type UserModify struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
	PassWordtwo string `json:"passwordtwo" `
	PassWordthree string `json:"passwordthree"`
}

func(usermodi *UserModify)Vaild() *serializer.Response {
	if usermodi.PassWordtwo != usermodi.PassWordthree{
		return &serializer.Response{
			Status: 40001,
			Msg:    "两次输入的密码不同，请重新输入",
			Error:  "",
		}
	}
	return nil
}

//修改密码
func(usermodi *UserModify)ChangePasswd() serializer.Response{
	var user model.User
	model.DB.Where("user_name = ?",usermodi.UserName).First(&user)
	//校验密码
	if user.CheckPassword(usermodi.PassWord)==false{
		return serializer.Response{
			Status: 40002,
			Msg:    "密码错误",
			Error:  "",
		}
	}
	if err:=user.SetPassword(usermodi.PassWordtwo);err!=nil{
		return serializer.Response{
			Status: 40002,
			Msg:    "密码加密失败",
		}
	}
	if err:=model.DB.Model(&user).Where("user_name = ?", usermodi.UserName).Update("password_digest", user.PasswordDigest).Error;err!=nil{
				return serializer.Response{
					Status:40001,
					Msg:"更改错误",
				}
			}else{
			return serializer.Response{
					Status: 200,
					Data:   nil,
					Msg:    "密码更改成功",
					Error:  "",
				}
			}
	}



//更改其他信息
type UserMemodify struct {
	Brithday string `form:"brithday" json:"brithday" `
	Gender   string `form:"gender" json:"gender"`
	Email    string `form:"email" json:"email" `
}

func (Info *UserMemodify) INfo(userid uint) serializer.Response {
	var user model.User
	model.DB.First(&user,userid)
	if err:=model.DB.Model(&user).Updates(map[string]interface{}{"brithday": Info.Brithday, "gender": Info.Gender, "email": Info.Email}).Error;err!=nil{
		return serializer.Response{
			Status: 40001,
			Data:   nil,
			Msg:    "更新失败",
			Error:  "",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.UserResponse(user),
		Msg:    "更新成功",
		Error:  "",
	}
}

//用户拥有的视频
//type Uservideoservice struct {
//}
//func(uvs *Uservideoservice)Get(uid uint)serializer.Response{
//	var videos []model.Video
//	if err:=model.DB.Where("user_id = ?",uid).Find(&videos).Error;err!=nil{
//		return serializer.Response{
//			Status: 50001,
//			Data:   nil,
//			Msg:    "数据库出错",
//			Error:  "",
//		}
//	}else{
//		return serializer.Response{
//			Status: 200,
//			Data:   serializer.BuildVideos(videos),
//			Msg:    "",
//			Error:  "",
//		}
//	}
//}

//更改用户头像
type AvatarUser struct {
	Avatar string `json:"avatar" binding:"required"`
}

func(ser *AvatarUser)Upavatar(uid uint)serializer.Response {
err :=	model.DB.Model(&model.User{}).Where("id = ?", uid).Update("avatar",ser.Avatar).Error
if err!=nil{
	return serializer.Response{
		Status: 50000,
		Data:   nil,
		Msg:    "错误",
		Error:  "",
	}
}else{
	return serializer.Response{
		Status: 200,
		Msg:    "更新头像成功！",
		Error:  "",
	}
}
}

