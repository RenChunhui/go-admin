package model

import (
	"fmt"

	"github.com/renchunhui/go-admin/pkg/config"
	"gorm.io/datatypes"
)

type Category struct {
	ID   uint           `json:"id" gorm:"primarykey"`
	Name string         `json:"name" gorm:"comment:'名称'"`
	Meta datatypes.JSON `json:"meta" gorm:"type:json;comment:'扩展对象'"`
}

func (c Category) TableName() string {
	return fmt.Sprintf("%scategory", config.Database.Prefix)
}
