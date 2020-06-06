package router

import (
	"GiliVideo/api"
	"GiliVideo/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine{
	r := gin.New()

	//所有中间键

	//gin自带的宕机恢复
	r.Use(gin.Recovery())
	//gin日志
	r.Use(gin.Logger())
	//session中间键
	r.Use(middleware.Session("mengjio007@ggvideo"))
	//跨域，用ngnix时，请关闭
	//r.Use(middleware.Cors())
	//获取登陆当前用户
	r.Use(middleware.CurrentUser())
	//获取当前管理员
	r.Use(middleware.AdminCurrentUser())


	//路由
	v1 := r.Group("/api/v1")
	{
		//搜索视频 参数为模糊查询所需参数
		v1.GET("videos/:param",api.SearchVideo)
		//展示最新上传的视频
		v1.GET("videos",api.ShowVideos)
		//展示播放量最高的几个视频
		v1.GET("hotvideos",api.ShowHotvideos)
		//播放视频 参数为视频id
		v1.GET("video/Av:id",api.ShowVideo)
		//展示点赞数最高的几个视频

		//用户登陆
		v1.POST("login",api.UserLogin)
		//用户注册
		v1.POST("regist",api.UserRegist)

		//给视频点赞
		v1.POST("star",api.AdStar)
		//给视频点踩
		v1.POST("lstar",api.AdLStar)
		//获取视频发布者信息：参数id为获取视频时的视频发布者id
   		v1.GET("poster/:id",api.Poster)
		//获取视频评论
		v1.POST("comments/Av:id",api.GetComments)
		//获取所有视频
		v1.GET("allvideos",api.SelectAll)
        v1.GET("getusera/:id",api.Getavatar)
		authed :=v1.Group("/")
		authed.Use(middleware.AuthRequired())//登陆保护中间键使用session
		{
			//用户登出，清楚redis/session
			authed.POST("user/logout",api.Logout)
			//用户删除视频
			authed.POST("deletevideo",api.DeleteVideo)
			authed.POST("modiinfo",api.ModiInfo)
			//获取上传视频封面的token
			authed.POST("avatartoken",api.UploadAvatarToken)
			//获取上传视频token
			authed.POST("videotoken",api.UploadVideoToken)
			//上传视频详情
			authed.POST("video",api.UploadVideo)
			//获取用户信息
			authed.GET("userinfo",api.InfoofUser)
			//获取用户上传的视频
			authed.POST("uservideo",api.Uservideo)
			//删除用户上传的视频
			authed.POST("user/AV:id")
			//获取上传头像token
			authed.POST("authavatartoken",api.UpAvatartoken)
			//上传头像
			authed.POST("upavatar",api.Upavatar)
			//修改密码
			authed.POST("modifypswd",api.ModifyPassWd)
			//
			//视频操作
			video := authed.Group("video")
			{
				//给视频发表评论
				video.POST("comment/Av:id",api.AddComment)

			}
		}
		Admin :=v1.Group("/admin")
		{
			Admin.POST("login",api.AdLogin)
			Admin.POST("regist",api.AdRegist)


			Ear := Admin.Group("/ear")
			Ear.Use(middleware.AdminAuthRequired())//管理员登陆保护
			{
				//获取
				//获取所有视频
				Ear.POST("allvideos",api.LiVideo)
				//获取所有用户
				Ear.POST("allusers",api.LiUser)
				//获取所有评论
				Ear.POST("allcomments",api.LiComment)
				//获取所有头像信息
				//Ear.POST("allavatars")


				//删除/封禁
				//删除某视频 传视频id列表
				Ear.POST("delvideos")
				//删除某用户 传用户id列表
				Ear.POST("delusers")
				//删除某评论 传评论id列表
				Ear.POST("delcomments")
				//删除某头像 传用户id列表
				Ear.POST("delavatars")

			}
		}

	}
	return  r
}