package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

type Allvideos struct {

}

func(all *Allvideos)All()serializer.Response{
	var videos []model.Video
	var count uint
	if err := model.DB.Model(model.Video{}).Count(&count).Error;err!= nil{
		return serializer.Response{
			Status: 50000,
			Data:   nil,
			Msg:    "查询视频总数失败",
			Error:  "",
		}
	}
	if err:=model.DB.Find(&videos).Error;err!=nil{
		return serializer.Response{
			Status:50001,
			Msg:"查询视频失败",
		}
	}
	return serializer.BuildListResponse(serializer.BuildVideos(videos),count)
}