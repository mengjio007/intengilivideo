package serializer

import (
	"GiliVideo/model"
)

// Video 视频序列化器
type Comment struct {
	ID uint			`json:"id"`
	VideoID uint	`json:"videoid"`
	UserID uint		`json:"userid"`
	Content string	`json:"content"`
	CreatedAt int64  `json:"created_at"`
}

//  序列化评论
func BuildComment(item model.Comment) Comment {
	return Comment{
		ID:          item.ID,
		VideoID:     item.VideoID,
		UserID:      item.UserID,
		Content:     item.Content,
		CreatedAt:   item.CreatedAt.Unix(),
	}
}

//  序列化评论列表
func BuildComments(items []model.Comment) (Comments []Comment) {
	for _, item := range items {
		comment := BuildComment(item)
		Comments = append(Comments, comment)
	}
	return Comments
}