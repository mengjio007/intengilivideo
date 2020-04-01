package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

type Commentslist struct {
	Limit int	`json:"limit"`
	Start int	`json:"count"`
}

func(service *Commentslist)Showcomments(vid uint) serializer.Response{
	var comments []model.Comment
	if service.Limit==0{
		service.Limit=10
	}
	var count uint
	if err := model.DB.Model(model.Comment{}).Where("videoid=?",vid).Count(&count).Error;err!= nil {
		return serializer.Response{
			Status: 40001,
			Data:   nil,
			Msg:    "查询评论数错误",
			Error:  "",
		}
	}
	if err := model.DB.Limit(service.Limit).Offset(service.Start).Where("videoid = ?",vid).Find(&comments).Error;err != nil{
		return serializer.Response{
			Status:40002,
			Msg:"评论返回失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildListResponse(serializer.BuildComments(comments),count),
		Msg:    "",
		Error:  "",
	}
}