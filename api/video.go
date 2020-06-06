package api

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
	"GiliVideo/service"
	"strconv"

	//"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	//"log"
	"net/http"
	//"os"
)

//展示单个视频
func ShowVideo(c *gin.Context){

	id := c.Param("id")
	service := service.Listvideoservice{}
	if err := c.ShouldBind(&service); err == nil {
		vid,_:=strconv.Atoi(id)
	res :=service.Showvideo(uint(vid))
	c.JSON(http.StatusOK,res)
	}else {
		c.JSON(http.StatusOK, serializer.Response{
			Status: 5000,
			Data:   nil,
			Msg:    "数据绑定失败",
			Error:  "",
		})
	}
}


//最新视频列表
func ShowVideos(c *gin.Context){
	service:=service.Listvideosservice{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Showvideos()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Status: 50002,
			Data:   nil,
			Msg:    "show错误",
			Error:  "",
		})
	}

}

//最热视频
func ShowHotvideos(c *gin.Context){
	ser:=service.HotVideo{}
	if err := c.ShouldBind(&ser); err == nil {
		res := ser.HotVideo()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Status: 50002,
			Data:   nil,
			Msg:    "show错误",
			Error:  "",
		})
	}
}



//上传视频
func UploadVideo(c *gin.Context){
	services := service.Videoform{}
	if err := c.ShouldBind(&services); err==nil {

		if user,ok := c.Get("user"); ok{
			uid := user.(*model.User).ID
			res := services.Upload(uid)
			c.JSON(http.StatusOK,res)
		}else{
			c.JSON(http.StatusOK,serializer.Response{
				Status: 50000,
				Data:   nil,
				Msg:    "上传视频需登录",
				Error:  "",
			})
		}

	}else{
			c.JSON(http.StatusOK,serializer.Response{
				Status: 40001,
				Msg:    "数据绑定错误",
				Error:  err.Error(),
			})
		}

}

////自己上传的视频
//func GetSelfVideo(c *gin.Context){
//	type uservideo struct {
//		uuid string
//
//	}
//}


//删除视频
//func DeleteVideo(c *gin.Context){
//
//}


//模糊搜索视频
func SearchVideo(c *gin.Context){
	param := c.Param("param")
	service:= service.SearchVideoservice{}
	if err := c.ShouldBind(&service);err == nil{
		res:=service.Search(param)
		c.JSON(http.StatusOK,res)

	}else {
		c.JSON(http.StatusOK, serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "数据绑定错误",
			Error:  "",
		})
	}
}

func Poster(c *gin.Context){
	services :=service.GetPosterservice{}
	vid := c.Param("id")
	if err:=c.ShouldBind(services);err!=nil{
		c.JSON(http.StatusOK,serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "佚名",
			Error:  "",
		})
	}else{
		id,_:=strconv.Atoi(vid)
		res := services.Get(uint(id))
		c.JSON(http.StatusOK,res)
	}

}

//查询所有视频

func SelectAll(c *gin.Context){
	service:= service.Allvideos{}
	if err := c.ShouldBind(&service);err == nil{
		res:=service.All()
		c.JSON(http.StatusOK,res)
	}else {
		c.JSON(http.StatusOK, serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "数据绑定错误",
			Error:  "",
		})
	}
}

func AdStar(c *gin.Context){
	ser:=service.Star{}
	if err:=c.ShouldBind(&ser);err == nil{
		res:=ser.AdStar()
		c.JSON(200,res)
	}else{
		c.JSON(200,serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "nook",
			Error:  "",
		})
	}
}

func AdLStar(c *gin.Context){
	ser:=service.Star{}
	if err:=c.ShouldBind(&ser);err == nil{
		res:=ser.AdLstar()
		c.JSON(200,res)
	}else{
		c.JSON(200,serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "nook",
			Error:  "",
		})
	}
}