package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//数据库单例
var DB *gorm.DB

//数据库初始化
func Database(){
	db,err := gorm.Open("mysql","")
	if err != nil{
		fmt.Println(err)
		panic("DB error")
	}

	//数据库连接池设置
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	db.LogMode(true)


	DB = db

	migration()
}

func migration(){
	DB.AutoMigrate(&Video{},&User{},&Comment{},&Admin{})
}

func init(){
	Database()
}