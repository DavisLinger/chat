package sql

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// 初始化MYSQL连接
func InitDb() {
	var err error
	DB, err = gorm.Open("mysql", "root:root@(127.0.0.1)/chat?charset=utf8&parseTime=True&loc=Local")
	DB.LogMode(true)
	if err != nil {
		fmt.Println(err)
	}
}
