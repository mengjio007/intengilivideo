package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

//分页查询用户服务
type  LiUsers struct{
	Limit  int  `json:"limit"`
	Start  int  `json:"start"`
}
func (list *LiUsers) Users() serializer.Response{
	var users []model.User
	if list.Limit==0{
		list.Limit = 10
	}
	var count uint
	if err := model.DB.Model(model.User{}).Count(&count).Error;err!= nil{
		return serializer.Response{
			Status: 5000,
			Data:   nil,
			Msg:    "查询视频总数失败",
			Error:  "",
		}
	}

	if err := model.DB.Limit(list.Limit).Offset(list.Start).Find(&users).Error;err != nil{
		return serializer.Response{
			Status: 5000,
			Data:   nil,
			Msg:    "limit查询失败",
			Error:  "",
		}
	}
	return serializer.BuildListResponse(serializer.BuildUsers(users),count)
}

//删除用户
type Deluid struct{
	uid uint
}
func(del *Deluid)Del()serializer.Response{
	var User model.User
	if err:=model.DB.First(&User,del.uid).Error;err!=nil{
		return serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "用户不存在",
			Error:  "",
		}
	}
	if err:=model.DB.Delete(&User,del.uid).Error;err!=nil{
		return serializer.Response{
			Status:40001,
			Data:nil,
			Msg:"删除出错",
		}
	}else{
		return serializer.Response{
			Status: 200,
			Data:   nil,
			Msg:    "删除成功",
			Error:  "",
		}
	}
}
//删除用户头像
func(del *Deluid)DelAvatar()serializer.Response{
	var User model.User
	if err:=model.DB.First(&User,del.uid).Error;err!=nil{
		return serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "用户不存在",
			Error:  "",
		}
	}
	if err:=model.DB.Model(&User).Update("avatar","").Error;err!=nil{
		return serializer.Response{
			Status:40001,
			Data:nil,
			Msg:"删除出错",
		}
	}else{
		return serializer.Response{
			Status: 200,
			Data:   nil,
			Msg:    "删除成功",
			Error:  "",
		}
	}
}