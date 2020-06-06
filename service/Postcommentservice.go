package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

type Postcomment struct {
	Content string `json:"content" binding:"required"`
}

func (p *Postcomment)Post(uid uint,vid uint) serializer.Response{
	comment := model.Comment{
		VideoID: vid,
		UserID:  uid,
		Content: p.Content,
	}
	err := model.DB.Create(&comment).Error
	if err != nil{
		return serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "评论失败",
			Error:  "",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildComment(comment),
		Msg:    "评论成功",
		Error:  "",
	}
}