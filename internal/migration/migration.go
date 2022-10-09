package migration

import (
	"github.com/renchunhui/go-admin/internal/model"
	"github.com/renchunhui/go-admin/pkg/config"
	"github.com/renchunhui/go-admin/pkg/db"
)

func New() {
	hasUser := db.Get().Migrator().HasTable(&model.User{})

	if !hasUser {
		db.Get().AutoMigrate(&model.User{Username: config.Root.Username, Password: config.Root.Password})

		user := model.User{}
		db.Get().Create(&user)
	}
}
