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
	Username string
	Password string
	Name     string
	Host     string
	Charset  string
}

type rootSchema struct {
	Username string
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
