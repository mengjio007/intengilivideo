package api

import (
	"GiliVideo/serializer"
	"GiliVideo/service"

	"github.com/gin-gonic/gin"
)

//  上传图片授权
func UploadAvatarToken(c *gin.Context) {
	service := service.UploadAvatarTokenService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Status: 50000,
			Data:   nil,
			Msg:    "",
			Error:  "",
		})
	}
}


//  上传视频授权
func UploadVideoToken(c *gin.Context) {
	service := service.UploadVideoTokenService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Status: 50000,
			Data:   nil,
			Msg:    "",
			Error:  "",
		})
	}
}