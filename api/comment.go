package api

import (
	"GiliVideo/serializer"
	service2 "GiliVideo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetComments(){

}

func AddComment(c *gin.Context){
	service :=service2.Postcoment{}
	if err :=c.ShouldBind(service);err!=nil{
		c.JSON(http.StatusOK,serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "服务获取失败",
			Error:  "",
		})
	}
	c.JSON(http.StatusOK,serializer.Response{
		Status: 200,
		Data:   nil,
		Msg:    "",
		Error:  "",
	})
}

//func AddGood(){
//
//}
//
//func DeleteGood(){
//
//}
