package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

//展示视频
type Listvideoservice struct{
}

//单个视频
func (list *Listvideoservice) Showvideo( id string) serializer.Response{
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	//处理视频被观看的一系问题
	video.AddView()

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}


type Listvideosservice struct{
	Limit int	`json:"limit"`
	Start int	`json:"count"`
}
//多个视频
func (list *Listvideosservice) Showvideos() serializer.Response{
	var videos []model.Video
	if list.Limit==0{
		list.Limit = 16
	}
	var count uint
	if err := model.DB.Model(model.Video{}).Count(&count).Error;err!= nil{
		return serializer.Response{
			Status: 5000,
			Data:   nil,
			Msg:    "查询视频总数失败",
			Error:  "",
		}
	}


	if err := model.DB.Limit(list.Limit).Offset(list.Start).Find(&videos).Error;err != nil{
		return serializer.Response{
			Status: 5000,
			Data:   nil,
			Msg:    "limit查询失败",
			Error:  "",
		}
	}
	return serializer.BuildListResponse(serializer.BuildVideos(videos),count)
}