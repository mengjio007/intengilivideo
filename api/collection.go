package api

import (
	"GiliVideo/serializer"
	service2 "GiliVideo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Collect(c *gin.Context){
	service := service2.Collect{}
	if err:=c.ShouldBind(&service);err!=nil{
		c.JSON(http.StatusOK,serializer.Response{
			Status: 50000,
			Data:   nil,
			Msg:    "服务出错",
			Error:  "",
		})
	}else{
		res:=service.CoNums()
		c.JSON(200,serializer.Response{
			Status: 200,
			Data:   res,
			Msg:    "",
			Error:  "",
		})
	}
}
