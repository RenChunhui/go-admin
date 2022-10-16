package model

type Tag struct {
	ID   uint   `json:"id" gorm:"primarykey"`
	Name string `json:"name" gorm:"comment:'名称'"`
}
