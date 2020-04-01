package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

type Deletevideo struct {
}

func (service *Deletevideo)Save(vid int) serializer.Response{
	err := model.DB.Where("ID = ?", vid).First(&model.Video{}).Error
	if err != nil{
		return serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "数据库错误",
			Error:  "",
		}
	}

	return serializer.Response{
		Status: 0,
		Data:   nil,
		Msg:    "",
		Error:  "",
	}
}