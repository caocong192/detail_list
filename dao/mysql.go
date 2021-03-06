package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dns := "caocong:cc1764..@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dns)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}