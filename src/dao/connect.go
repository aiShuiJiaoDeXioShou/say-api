package dao

import (
	"log"
	"time"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

//在引用这个文件之前，初始化数据库连接
func init() {
	if DB == nil {
		DB = Connect()
	}
}

//连接数据库
func Connect() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/say?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		panic(err)
	}
	dbconfig, err := db.DB()
	if err != nil {
		log.Println(err)
	}
	dbconfig.SetMaxIdleConns(10)
	dbconfig.SetMaxOpenConns(100)
	dbconfig.SetConnMaxLifetime(time.Hour)

	return db
}
