package db

import (
	"fmt"

	"github.com/renchunhui/go-admin/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// 获取实例
func Get() *gorm.DB {
	return db
}

// 连接
func Open() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Name,
		config.Database.Charset,
	)
	mydb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = mydb

	if err != nil {
		return err
	}

	fmt.Printf("\033[1;32m%s\033[0m", "✔ 数据库连接成功\n")

	return nil
}

// 断开
func Close() {

}
