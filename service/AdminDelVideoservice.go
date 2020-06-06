package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

//分页查询视频服务
type  ListVideos struct{
	Limit  int  `json:"limit"`
	Start  int  `json:"start"`
}
func (list *ListVideos) Videos() serializer.Response{
	var videos []model.Video
	if list.Limit==0{
		list.Limit = 10
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

type Delvid struct{
	vid uint
}
func(del *Delvid)Del()serializer.Response{
	var Video model.Video
	if err:=model.DB.First(&Video,del.vid).Error;err!=nil{
		return serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "数据不存在",
			Error:  "",
		}
	}
	if err:=model.DB.Delete(&Video,del.vid).Error;err!=nil{
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