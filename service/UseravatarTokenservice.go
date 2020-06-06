package service


// UploauserdavatarTokenService 获得上传用户头像oss token的服务
import (
	"GiliVideo/serializer"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"path/filepath"
)

type UploadUserTokenService struct {
	FileAvatarname string `form:"fileavatarame" json:"fileavatarname"`
}

// Post 创建token
func (service *UploadUserTokenService) Post() serializer.Response {
	client, err := oss.New("oss-cn-hongkong.aliyuncs.com", "LTAI4FxAEupvUBzdNhZucg1G", "wbtQOguOztGK5iS4QvM7JvkB1JPbCX")
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
	ext := filepath.Ext(service.FileAvatarname)
	if ext != ".jpg"{
		return serializer.Response{
			Status: 50002,
			Data:   nil,
			Msg:    "头像只能为jpg" ,
			Error:  "",
		}
	}
	// 带可选参数的签名直传。
	options := []oss.Option{
		oss.ContentType("image/jpeg"),

	}

	key := "UserAvatar/vid" + uuid.Must(uuid.NewRandom()).String() + ext
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
