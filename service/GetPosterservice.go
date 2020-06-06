package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

type GetPosterservice struct{
}

func (service *GetPosterservice)Get(id uint) serializer.Response{
	var user  model.User
	err := model.DB.Where("id = ?",id).First(&user).Error
	if err != nil{
		return serializer.Response{
			Status: 50000,
			Data:"佚名",
		}
	}
	return serializer.UserResponse(user)
}