package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

type Deletevideo struct {
	Vid int `json:"vid" binding:"required"`
}

func (service *Deletevideo)Delete() serializer.Response{
	var video model.Video
	err := model.DB.First(&video,uint(service.Vid)).Error
	if err != nil{
		return serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "数据库错误",
			Error:  "",
		}
	}else {
		model.DB.Unscoped().Delete(&video)
		return serializer.Response{
			Status: 200,
			Data:   nil,
			Msg:    "删除成功",
			Error:  "",
		}
	}
}