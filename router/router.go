package router

import (
	"GiliVideo/api"
	"GiliVideo/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine{
	r := gin.Default()

	//中间键
	r.Use(middleware.Session("mengjio007@ggvideo"))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	//路由
	v1 := r.Group("/api/v1")
	{

		v1.GET("videos/:param",api.SearchVideo)
		//视频展示
		v1.GET("videos",api.ShowVideos)
		v1.GET("video/Av:id",api.ShowVideo)


		v1.GET("video/comments/Av:id",api.GetComments)

		//用户登陆
		v1.POST("login",api.UserLogin)
		//用户注册
		v1.POST("regist",api.UserRegist)
		v1.POST("video",api.UploadVideo)
		v1.DELETE("logout",api.Logout)
		v1.POST("avatartoken",api.UploadAvatarToken)
		v1.POST("videotoken",api.UploadVideoToken)



		authed :=v1.Group("/")
		authed.Use(middleware.AuthRequired())//登陆保护中间键使用session
		{
			//用户
			authed.GET("user/space",api.GetSpace)
			authed.DELETE("user/logout",api.Logout)


			//视频操作
			video := authed.Group("video")
			{

				video.GET("/videos",api.GetSelfVideo)
				video.DELETE("/delete",api.DeleteVideo)
				video.POST("/comment",api.AddComment)
			}
		}


	}
	return  r
}