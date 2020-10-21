package models

import "github.com/jinzhu/gorm"

type User struct{
	gorm.Model
	NickName string `gorm:"column:nickname" json:"nickname"`
	FirstName string `gorm:"column:firstname" json:"firstname"`
	LastName string `gorm:"column:lastname" json:"lastname"`
	Age int `gorm:"column:age" json:"age"`
}
