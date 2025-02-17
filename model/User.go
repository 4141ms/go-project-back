package model

import (
	"fmt"
	"go-project-back/utils/errmsg"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"` 
	Password string `gorm:"type:varchar(20);not null" json:"password"` 
	Role int `gorm:"type:int;not null" json:"role"` 
	// Desc string `gorm:"type:varchar(20);not null" json:"desc"` 
	// Token string `gorm:"type:varchar(20);not null" json:"token"` 
	// Avatar string `gorm:"type:varchar(20);not null" json:"avatar"`   
	// Buttons string `gorm:"type:varchar(20);not null" json:"buttons"`  
	// Routes string `gorm:"type:varchar(20);not null" json:"routes"`   
}

// 查询用户是否存在
func CheckUser(name string)(code int){
	var users User
	result := db.Select("id").Where("username = ?", name).First(&users)
	if result.Error != nil {
		fmt.Println("用户未找到,用户名:",name)
		return errmsg.SUCCESS
	}
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED // 用户名已存在 1001
	}
	
	return errmsg.SUCCESS
}

// 新增用户      这里要传指针
func CreateUser(data *User)int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	} 
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 编辑用户


// 删除用户