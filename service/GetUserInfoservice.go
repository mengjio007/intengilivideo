package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

type UserInfo struct {
}

func(serv *UserInfo)GetAllInfo(id uint)serializer.Response{
	user:=model.User{}
	model.DB.First(&user,id)
	return serializer.UserResponse(user)
}