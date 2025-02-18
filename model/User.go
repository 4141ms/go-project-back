package model

import (
	"encoding/base64"
	"fmt"
	"go-project-back/utils/errmsg"
	"log"

	"golang.org/x/crypto/scrypt"
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
	// 在存入数据库之前将密码进行加密 (使用下面的钩子函数)
	// data.Password = ScryptPw(data.Password)
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
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int)int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

// 密码加密
// 钩子函数
func(u *User)BeforeSave(tx *gorm.DB) (err error){
	u.Password = ScryptPw(u.Password)

	return
}
// 使用 Scrypt 进行密码加密
func ScryptPw(password string) string{
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12,32,4,6,66,22,222,11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}

	fpw := base64.StdEncoding.EncodeToString(HashPw)

	return fpw

}