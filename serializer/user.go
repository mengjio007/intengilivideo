package serializer

import "GiliVideo/model"

//用户序列化器
type User struct {
	ID uint	`json:"id"`
	UserName string	`json:"username"`
	Nickname string	`json:"nickname`
	Status	string	`json:"status"`
	Avatar	string	`json:"avatar"`
	CreatedAt int64	`json:"created_at"`
}


func UserBuild(item model.User) User{
	return User{
		ID:item.ID,
		UserName:item.UserName,
		Nickname:item.Nickname,
		Status:item.Status,
		Avatar:item.Avatar,
		CreatedAt:item.CreatedAt.Unix(),
	}
}


//序列化用户响应
func UserResponse(item model.User) Response{
	return Response{
		Status: 200,
		Data:   UserBuild(item),
		Msg:    "",
		Error:  "",
	}
}