package model

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName       string
	AccountNb	   uint
	PasswordDigest string
	Nickname       string
	Email		   string
	Brithday	   string
	Gender		   string			//1为男2为女
	Status         string
	Avatar         string   `gorm:"size:1000"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 加密密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

func (user *User) AvatarUrl() string{
	client, _ := oss.New("oss-cn-hongkong.aliyuncs.com", "LTAI4FxAEupvUBzdNhZucg1G", "wbtQOguOztGK5iS4QvM7JvkB1JPbCX")
	bucket, _ := client.Bucket("gilivideo")
	signedGetURL, _ := bucket.SignURL(user.Avatar, oss.HTTPGet, 600)
	return signedGetURL
}

