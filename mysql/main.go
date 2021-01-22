package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}
func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if(err!=nil){
		panic(err)
	}
	//创建表
	db.AutoMigrate(&UserInfo{})

	db.Create(&UserInfo{1,"闵续","男","唱跳rap"})
	//db.Create(&UserInfo{2,"asdsa","asdas","sadasdas"})
	//db.Commit()


}
