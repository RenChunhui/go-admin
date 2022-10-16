package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/renchunhui/go-admin/pkg/db"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (u User) Create(c *gin.Context) {
	err := c.ShouldBind(&u)

	if err != nil {
		log.Fatal(err)
	}

	db.GetInstance().Create(&u)
}
func (u User) Get(c *gin.Context) {

}
func (u User) Update(c *gin.Context) {}
func (u User) Delete(c *gin.Context) {}
