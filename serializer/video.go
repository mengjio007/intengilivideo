package serializer

import "GiliVideo/model"

// Video 视频序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"Info"`
	Uid       uint   `json:"uid"`
	Tag       string  `json:"tag"`
	Root      string  `json:"root"`
	Url       string `json:"url"`
	Avatar    string `json:"avatar"`
	View      uint64 `json:"view"`
	Star      uint64 `json:"star"`
	Lstar     uint64 `json:"lstar"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		Uid:       item.UserID,
		Tag:       item.Tag,
		Root:      item.Root,
		Url:       item.VideoURL(),
		Avatar:    item.AvatarURL(),
		View:      item.View(),
		Star:      item.Star(),
		Lstar:      item.Lstar(),
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildVideos 序列化视频列表
func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}