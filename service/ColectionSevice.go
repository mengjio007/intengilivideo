package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

type Collect struct{
	Vid uint `json:"vid" binding:"required"`
	Collection bool `json:"collection" binding:"required"`
}

func (col *Collect)CoNums()serializer.Response{
	var count uint
	model.DB.Model(&model.Collection{}).Where("vid = ?",col.Vid).Count(&count)
	return serializer.Response{
		Status: 200,
		Data:   count,
		Msg:    "",
		Error:  "",
	}
}

//func (col *Collect)Uincoll(Uid uint)serializer.Response{
//	var colle model.Collection
//	if err := model.DB.First(&colle).Error;err !=nil{
//		return serializer.Response{
//			Status: 50001,
//			Data:   nil,
//			Msg:    "未收藏",
//			Error:  "",
//		}
//	}else {
//		return serializer.Response{
//			Status: 200,
//			Data:   nil,
//			Msg:    "已收藏",
//			Error:  "",
//		}
//	}
//}