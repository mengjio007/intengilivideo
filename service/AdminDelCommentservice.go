package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

//分页查询评论服务
type  Licomment struct{
	Limit  int  `json:"limit"`
	Start  int  `json:"start"`
}
func (list *Licomment) Comments() serializer.Response{
	var comments []model.Comment
	if list.Limit==0{
		list.Limit = 10
	}
	var count uint
	if err := model.DB.Model(model.Comment{}).Count(&count).Error;err!= nil{
		return serializer.Response{
			Status: 5000,
			Data:   nil,
			Msg:    "查询视频总数失败",
			Error:  "",
		}
	}

	if err := model.DB.Limit(list.Limit).Offset(list.Start).Find(&comments).Error;err != nil{
		return serializer.Response{
			Status: 5000,
			Data:   nil,
			Msg:    "limit查询失败",
			Error:  "",
		}
	}
	return serializer.BuildListResponse(serializer.BuildComments(comments),count)
}

//删除评论
type DelMsg struct{
	cid uint
}
func(del *DelMsg)Del()serializer.Response{
	var Comment model.Comment
	if err:=model.DB.First(&Comment,del.cid).Error;err!=nil{
		return serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "数据不存在",
			Error:  "",
		}
	}
	if err:=model.DB.Delete(&Comment,del.cid).Error;err!=nil{
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