package model

import (
	"fmt"

	"github.com/renchunhui/go-admin/pkg/config"
	"gorm.io/datatypes"
)

type Tag struct {
	ID   uint           `json:"id" gorm:"primarykey"`
	Name string         `json:"name" gorm:"comment:'名称'"`
	Meta datatypes.JSON `json:"meta" gorm:"type:json;comment:'扩展对象'"`
}

func (t Tag) TableName() string {
	return fmt.Sprintf("%stag", config.Database.Prefix)
}
