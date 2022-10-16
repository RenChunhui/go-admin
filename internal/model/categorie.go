package model

import "gorm.io/datatypes"

type Categorie struct {
	ID   uint           `json:"id" gorm:"primarykey"`
	Name string         `json:"name" gorm:"comment:'名称'"`
	Meta datatypes.JSON `json:"meta" gorm:"type:json;comment:'扩展对象'"`
}

func (c Categorie) TableName() string {
	return "categorie"
}
