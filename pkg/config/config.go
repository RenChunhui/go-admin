package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type httpSchema struct {
	Port uint
}

type databaseSchema struct {
	// 数据库名
	Name string
	// 用户名
	Username string
	// 密码
	Password string
	// 主机
	Host string
	// 数据表前缀
	Prefix string
	// 编码
	Charset string
}

type rootSchema struct {
	// 管理员账号
	Username string
	// 管理员密码
	Password string
}

var (
	Http     *httpSchema
	Database *databaseSchema
	Root     *rootSchema
)

func NewConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName(os.Getenv("GIN_MODE"))
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	viper.UnmarshalKey("Http", &Http)
	viper.UnmarshalKey("Database", &Database)
	viper.UnmarshalKey("Root", &Root)

	fmt.Printf("\033[1;32m%s\033[0m", "✔ 配置文件读取成功\n")

	return nil
}
