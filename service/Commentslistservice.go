package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
	"strconv"
)

type Commentslist struct {
	Limit int	`json:"limit" binding:"required"`
}

func(service *Commentslist)Showcomments(vid uint) serializer.Response{
	var comments []model.Comment
	var count uint
	if err := model.DB.Model(model.Comment{}).Where("video_id=?",vid).Count(&count).Error;err!= nil {
		return serializer.Response{
			Status: 40001,
			Data:   nil,
			Msg:    "查询评论数错误",
			Error:  "",
		}
	}
	if err := model.DB.Order("id DESC").Limit(8).Offset(service.Limit*8-8).Where("video_id = ?",vid).Find(&comments).Error;err != nil{
		return serializer.Response{
			Status:40002,
			Msg:"评论返回失败",
		}
	}
	return serializer.BuildListResponse(serializer.BuildComments(comments),count)
}

type Commentslistss struct {

}

func(service *Commentslistss)Showcoments(uid string)serializer.Response{
	user:=model.User{}
	vid, _ := strconv.Atoi(uid)
	model.DB.Where("id = ?",vid).First(&user)
	return serializer.UserResponse(user)
}