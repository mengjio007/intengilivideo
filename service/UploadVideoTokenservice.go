package service

import (
	"GiliVideo/serializer"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"path/filepath"
)

// UploadTokenService 获得上传oss token的服务
type UploadVideoTokenService struct {
	Filevideoname string `form:"filevideoame" json:"filevideoname"`
}

// Post 创建token
func (service *UploadVideoTokenService) Post() serializer.Response {
	client, err := oss.New("oss-cn-hongkong.aliyuncs.com", "", "")
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 获取存储空间。
	bucket, err := client.Bucket("gilivideo")
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 获取扩展名
	ext := filepath.Ext(service.Filevideoname)
	if ext != ".mp4"{
		return serializer.Response{
			Status: 50002,
			Data:   nil,
			Msg:    "视频格式只能为MP4" ,
			Error:  "",
		}
	}
	// 带可选参数的签名直传。
	options := []oss.Option{
		oss.ContentType("video/mp4"),

	}

	key := "video/vid" + uuid.Must(uuid.NewRandom()).String() + ext
	// 签名直传。
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}
	// 查看图片
	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			"get": signedGetURL,
		},
	}
}
