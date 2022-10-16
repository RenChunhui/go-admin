package model

import (
	"fmt"

	"github.com/renchunhui/go-admin/pkg/config"
)

type Role struct {
	Model

	Name string `json:"name"`
}

func (r Role) TableName() string {
	return fmt.Sprintf("%srole", config.Database.Prefix)
}
