package service

import (
	"GiliVideo/cache"
	"GiliVideo/model"
	"GiliVideo/serializer"
	"strconv"
)

//展示视频
type Listvideoservice struct{
}

//单个视频
func (list *Listvideoservice) Showvideo( id uint) serializer.Response{
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	//处理视频被观看的一系问题
	video.AddView()

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}


type Listvideosservice struct{
	Limit int	`json:"limit"`
	Start int	`json:"count"`
}
//逆序新鲜多个视频
func (list *Listvideosservice) Showvideos() serializer.Response{
	var videos []model.Video
	if list.Limit==0{
		list.Limit = 8
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

	if err := model.DB.Order("id DESC").Limit(list.Limit).Offset(list.Start).Find(&videos).Error;err != nil{
		return serializer.Response{
			Status: 5000,
			Data:   nil,
			Msg:    "limit查询失败",
			Error:  "",
		}
	}
	return serializer.BuildListResponse(serializer.BuildVideos(videos),count)
}

//火热视频

type HotVideo struct {
}

func(res *HotVideo)HotVideo()serializer.Response{
	var videos []model.Video
	var video1 model.Video
	var video2 model.Video
	var video3 model.Video
	var video4 model.Video
	var video5 model.Video
	var video6 model.Video
	var video7 model.Video
	var video8 model.Video
	nums := model.HotVideo()
	var avids []uint
	for i:=0;i<8;i++{
		id:=nums[i]
		avid,_:=strconv.Atoi(id)
		avids =append(avids,uint(avid))
	}

	//model.DB.Where("id in (?)",avids).Find(&videos)
	model.DB.Where("id= ?",avids[0]).First(&video1)
	model.DB.Where("id= ?",avids[1]).First(&video2)
	model.DB.Where("id= ?",avids[2]).First(&video3)
	model.DB.Where("id= ?",avids[3]).First(&video4)
	model.DB.Where("id= ?",avids[4]).First(&video5)
	model.DB.Where("id= ?",avids[5]).First(&video6)
	model.DB.Where("id= ?",avids[6]).First(&video7)
	model.DB.Where("id= ?",avids[7]).First(&video8)
	videos = append(videos,video1)
	videos = append(videos,video2)
	videos = append(videos,video3)
	videos = append(videos,video4)
	videos = append(videos,video5)
	videos = append(videos,video6)
	videos = append(videos,video7)
	videos = append(videos,video8)
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildVideos(videos),
		Msg:    "",
		Error:  "",
	}
}

type Star struct{
	ID uint `json:"id" binding:"required"`
}

func (s *Star)AdStar()serializer.Response{
	//增加视频点赞数
	cache.RedisClient.Incr(cache.VideoStar(s.ID))
	return serializer.Response{
		Status:200,
		Msg:"ok",
	}
}

func(s *Star)AdLstar()serializer.Response{
	//增加视频点踩数
	cache.RedisClient.Incr(cache.VideoLStar(s.ID))
	return serializer.Response{
		Status:200,
		Msg:"ok",
	}
}

//type MoreStart struct {
//
//}
//func(star *MoreStart)GetMStar(){
//
//}