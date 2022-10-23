package migration

import (
	"github.com/renchunhui/go-admin/internal/model"
	"github.com/renchunhui/go-admin/pkg/config"
	"github.com/renchunhui/go-admin/pkg/db"
)

func New() {
	hasUser := db.GetInstance().Migrator().HasTable(&model.User{})

	// User table
	if !hasUser {
		db.GetInstance().AutoMigrate(&model.User{})

		user := model.User{Username: config.Root.Username, Password: config.Root.Password}
		db.GetInstance().Create(&user)
	}

	// Role table
	hasRole := db.GetInstance().Migrator().HasTable(&model.Role{})

	if !hasRole {
		db.GetInstance().AutoMigrate(&model.Role{Name: "admin"})

		role := model.Role{}
		db.GetInstance().Create(&role)
	}

	// Tag
	hasTag := db.GetInstance().Migrator().HasTable(&model.Tag{})

	if !hasTag {
		db.GetInstance().AutoMigrate(&model.Tag{})
	}

	// Category
	hasCategory := db.GetInstance().Migrator().HasTable(&model.Category{})

	if !hasCategory {
		db.GetInstance().AutoMigrate(&model.Category{})
	}
}
