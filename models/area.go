package models

type Area struct {
	BaseModel
	Name        string `gorm:"column:name;size:47" json:"name"`
	Description string `gorm:"column:description;size:200" json:"description"`
	Sort        int    `gorm:"column:sort;default:0" json:"sort"`
}
