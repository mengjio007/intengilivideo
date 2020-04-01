package model

import (
	"GiliVideo/cache"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"strconv"
)
type Video struct {
	gorm.Model	`json:"-"`
	UserID uint	`json:"-"`
	Title string	`json:"title"`
	Info string `json:"info"`
	Tag string `json:"tag" `
	Path string		`json:"path"`
	Avatar string	`json:"avatar"`
	Comment []Comment	`json:"comment"`
	Time string		`gorm:"-" json:"time"`
}

const (
	//tag
	Selfcontrol string = "自制"
	//tag
	Reprint string = "转载"
	////tag
	//Funny string = "搞笑"
	////tag
	//Anmation string = "动漫"
	////tag
	//Game string = "游戏"
)

func (v *Video) AfterFind() {
	v.Time = v.CreatedAt.Format("2006-01-02 15:04:05")
}

// AvatarURL 封面地址
func (video *Video) AvatarURL() string {
	client, _ := oss.New("oss-cn-hongkong.aliyuncs.com", "LTAI4FxAEupvUBzdNhZucg1G", "wbtQOguOztGK5iS4QvM7JvkB1JPbCX")
	bucket, _ := client.Bucket("gilivideo/avatar/*")
	signedGetURL, _ := bucket.SignURL(video.Avatar, oss.HTTPGet, 600)
	return signedGetURL
}

// VideoURL 视频地址
func (video *Video) VideoURL() string {
	client, _ := oss.New("oss-cn-hongkong.aliyuncs.com", "LTAI4FxAEupvUBzdNhZucg1G", "wbtQOguOztGK5iS4QvM7JvkB1JPbCX")
	bucket, _ := client.Bucket("gilivideo/video/*")
	signedGetURL, _ := bucket.SignURL(video.Path, oss.HTTPGet, 600)
	return signedGetURL
}

// View 点击数
func (video *Video) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.VideoViewKey(video.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}


// AddView 视频游览
func (video *Video) AddView() {
	// 增加视频点击数
	cache.RedisClient.Incr(cache.VideoViewKey(video.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(video.ID)))
}
