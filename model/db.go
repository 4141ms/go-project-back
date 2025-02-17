package model

import (
	"fmt"
	"go-project-back/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDb(){
	// db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
	// 	utils.DbUser,
	// 	utils.DbPassword,
	// 	utils.DbHost,
	// 	utils.DbPort,
	// 	utils.DbName,
	// ))
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
	utils.DbUser,
	utils.DbPassword,
	utils.DbHost,
	utils.DbPort,
	utils.DbName,
	)
  	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("连接数据库失败，请检查参数", err)
	}


	db.AutoMigrate(&User{},&Category{},&Article{})

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10*time.Second)


}