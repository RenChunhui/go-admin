package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/renchunhui/go-admin/internal/model"
	"github.com/renchunhui/go-admin/pkg/db"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser() User {
	return User{}
}

// 创建用户
func (u User) Create(c *gin.Context) {
	err := c.ShouldBind(&u)

	if err != nil {
		log.Fatal(err)
	}

	db.GetInstance().Create(&model.User{Username: u.Username, Password: u.Password})
}

// 获取用户
func (u User) Get(c *gin.Context) {

}

// 更新用户
func (u User) Update(c *gin.Context) {}

// 删除用户
func (u User) Delete(c *gin.Context) {}
