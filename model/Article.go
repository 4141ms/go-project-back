package model

import "gorm.io/gorm"

// 如果有外键需要添加外键

type Article struct {
	Category Category `gorm:"foreignKey:Cid"`
	gorm.Model
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	Cid int `gorm:"type:int;not null" json:"cid"`
	Desc string `gorm:"type:varchar(200);not null" json:"desc"`
	Content string `gorm:"type:longtext;not null" json:"content"`
	Img string `gorm:"type:varchar(100);not null" json:"img"`
}