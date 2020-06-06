package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

type ListUservideoservice struct{
	Limit int	`json:"limit"`
	Start int	`json:"count"`
}

func(listuvideo *ListUservideoservice)List(uid uint) serializer.Response{
	var videos []model.Video
	if listuvideo.Limit==0{
		listuvideo.Limit =5
	}
	var count uint
	if err := model.DB.Model(model.Video{}).Where("user_id = ?",uid).Count(&count).Error;err!= nil{
		return serializer.Response{
			Status: 5000,
			Data:   nil,
			Msg:    "查询视频总数失败",
			Error:  "",
		}
	}

	if err := model.DB.Order("id DESC").Limit(listuvideo.Limit).Offset(listuvideo.Start).Where("user_id= ?",uid).Find(&videos).Error;err != nil{
		return serializer.Response{
			Status: 5000,
			Data:   nil,
			Msg:    "limit查询失败",
			Error:  "",
		}
	}
	return serializer.BuildListResponse(serializer.BuildVideos(videos),count)
}