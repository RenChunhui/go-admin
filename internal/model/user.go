package model

import (
	"fmt"

	"github.com/renchunhui/go-admin/pkg/config"
	"github.com/renchunhui/go-admin/pkg/db"
	"gorm.io/datatypes"
)

type User struct {
	Model

	Username string         `json:"username" gorm:"type:varchar(50);comment:用户名"`
	Password string         `json:"password" gorm:"type:text;comment:密码"`
	Email    string         `json:"email" gorm:"type:varchar(50);comment:邮箱"`
	Meta     datatypes.JSON `json:"meta" gorm:"type:json;comment:扩展对象"`
}

func (u User) TableName() string {
	return fmt.Sprintf("%suser", config.Database.Prefix)
}

func (u User) Create() {
	db.GetInstance().Create(&u)
}

func (u User) Read() {}
