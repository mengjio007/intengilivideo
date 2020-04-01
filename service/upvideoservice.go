package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

//uploadvideo service
type Videoform struct {
	Title string `form:"title" json:"title" binding:"required"`
	Info  string `form:"info"  json:"info" binding:""`
	Tag  string  `form:"tag" json:"tag" binding:""`
	Avatar string `form:"avatar" json:"avatar" binding:"required"`
	Path string	`form:"path" json:"path" binding:"required"`
}

func (videof *Videoform) Upload(userid uint) serializer.Response{
	 video := model.Video{
		Title : videof.Title ,
		Info : videof.Info ,
		Tag :  videof.Tag,
		Path: videof.Path,
		UserID: userid,
	}
	err := model.DB.Create(&video).Error
	if  err !=nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "视频创建失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status:200,
		Data : serializer.BuildVideo(video),
	}
}
