package model

import "gorm.io/datatypes"

type User struct {
	Model

	Username string         `json:"username" gorm:"type:varchar(50);comment:'用户名'"`
	Password string         `json:"password" gorm:"type:text;comment:'密码'"`
	Email    string         `json:"email" gorm:"type:varchar(50);comment:'邮箱'"`
	Phone    string         `json:"phone" gorm:"type:varchar(11);comment:'电话'"`
	Nickname string         `json:"nickname" gorm:"type:varchar(50);comment:'昵称'"`
	Avatar   string         `json:"avatar" gorm:"comment:'头像'"`
	RoleId   uint8          `json:"role" gorm:"comment:'角色 ID'"`
	Meta     datatypes.JSON `json:"meta" gorm:"type:json;comment:'扩展对象'"`
}

func (u User) TableName() string {
	return "user"
}
