package model

import (
	"fmt"
	"go-project-back/utils/errmsg"

	"gorm.io/gorm"
)

type Category struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`   
}

// 查询分类是否存在
func CheckCategory(name string)(code int){
	var cate Category
	result := db.Select("id").Where("name = ?", name).First(&cate)
	if result.Error != nil {
		fmt.Println("用户未找到,用户名:",name)
		return errmsg.SUCCESS
	}
	if cate.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED // 用户名已存在 1001
	}
	
	return errmsg.SUCCESS
}

// todo 查询分类下所有文章


// 新增分类      这里要传指针
func CreateCategory(data *Category)int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	} 
	return errmsg.SUCCESS
}

// 查询分类列表
func GetCategory(pageSize int, pageNum int) []Category {
	var cates []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cates
}

// 编辑分类
func EditCategory(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除分类
func DeleteCategory(id int)int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}