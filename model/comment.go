package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model `json:"-"`
	VideoID uint
	UserID uint
	Content string
	Time string
}

func (c *Comment)AfterFind()(err error)  {
	c.Time = c.CreatedAt.Format("2006-01-02 15:04:05")

	return nil
}