package service

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
)

type SearchVideoservice struct{
}

func(sevice *SearchVideoservice) Search(param string) serializer.Response{
	videos := []model.Video{}
	 err := model.DB.Where("title LIKE ?","%"+param+"%").Find(&videos).Error
	 if err != nil{
		return serializer.Response{
			Status: 50000,
			Data:   nil,
			Msg:    "数据库错误",
			Error:  "",
			}
	 }
	 return serializer.Response{
		 Status: 200,
		 Data:   serializer.BuildVideos(videos),
		 Msg:    "",
		 Error:  "",
	 }

}