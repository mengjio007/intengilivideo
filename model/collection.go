package model


type Collection struct {
	ID uint `gorm:"primary_key"`
	UserId uint
	VideoId  uint
}
