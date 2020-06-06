package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	gorm.Model
	AdminName string
	AdminPassword string
}

// SetPassword 加密密码
func (admin *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	admin.AdminPassword = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(admin.AdminPassword), []byte(password))
	return err == nil
}