package model

import "time"

type User struct {
	ID            int64     `gorm:"primary_key;column:id"`
	Name          string    `gorm:"column:name"`
	PictureUrl    string    `gorm:"column:picture_url"`
	StatusMessage string    `gorm:"column:status_message"`
	Datetime      time.Time `gorm:"column:datetime"`
}

func (User) TableName() string {
	return "user"
}
