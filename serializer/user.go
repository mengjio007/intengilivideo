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
	Email	 string `json:"email"`
	Brithday   string `json:"brithday"`
	Gender	   string `json:"gender"`
}


func UserBuild(item model.User) User{
	return User{
		ID:item.ID,
		UserName:item.UserName,
		Nickname:item.Nickname,
		Status:item.Status,
		Email:item.Email,
		Brithday:item.Brithday,
		Gender:item.Gender,
		Avatar:item.AvatarUrl(),
		CreatedAt:item.CreatedAt.Unix(),
	}
}


//序列化用户列表
func BuildUsers(items []model.User) (Users []User) {
	for _, item := range items {
		user := UserBuild(item)
		Users = append(Users, user)
	}
	return Users
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

//管理员序列化器
type Admin struct {
	ID uint	`json:"id"`
	AdminName string	`json:"adminname"`
}

func AdminBuild(item model.Admin) Admin{
	return Admin{
		ID:        item.ID,
		AdminName: item.AdminName,
	}
}

//序列化管理员响应
func AdminResponse(item model.Admin) Response{
	return Response{
		Status: 200,
		Data:   AdminBuild(item),
		Msg:    "",
		Error:  "",
	}
}


//评论用户结构体
type ComUSer struct {
	ID uint	`json:"id"` //评论id
	UserName string	`json:"username"`
	Avatar	string	`json:"avatar"`
	VideoID uint	`json:"videoid"`
	UserID uint		`json:"userid"`
	Content string	`json:"content"`
	CreatedAt int64	`json:"created_at"`
}