package api

import (
	"GiliVideo/model"
	"GiliVideo/serializer"
	"GiliVideo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetComments(c *gin.Context){
	service:=service.Commentslist{}
	vid:=c.Param("id")
	id,_:=strconv.Atoi(vid)
	if err :=c.ShouldBind(&service);err!=nil{
		c.JSON(http.StatusOK,serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "获取评论失败，数据绑定错误",
			Error:  "",
		})
	}else{
		ser :=service.Showcomments(uint(id))
		c.JSON(http.StatusOK,ser)
	}

}

func AddComment(c *gin.Context){
	service :=service.Postcomment{}
	vid :=c.Param("id")
	if err :=c.ShouldBind(&service);err!=nil{
		c.JSON(http.StatusOK,serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "评论失败，数据绑定错误",
			Error:  "",
		})
	}else{
		id, _ := strconv.Atoi(vid)
		if user,ok := c.Get("user"); ok {
			uid := user.(*model.User).ID
			res := service.Post(uid, uint(id))
			c.JSON(http.StatusOK, serializer.Response{
				Status: 200,
				Data:   res,
				Msg:    "",
				Error:  "",
			})
		}else{
			c.JSON(200,serializer.Response{
				Status: 50001,
				Data:   nil,
				Msg:    "错误",
				Error:  "",
			})
		}
	}
}

//func AddGood(){
//
//}
//
//func DeleteGood(){
//
//}
func Getavatar(c *gin.Context ){
	ser :=service.Commentslistss{}
	if err:=c.ShouldBind(ser);err!=nil{
	c.JSON(200,serializer.Response{
		Status: 40001,
		Data:   nil,
		Msg:    "数据错处",
		Error:  "",
	})
	}else{
		uid:=c.Param("id")
		res:=ser.Showcoments(uid)
		c.JSON(200,res)

	}
}