package models

type CategoryModel struct {
	MODEL `json:","`
	Title string `gorm:"size:30" json:"title"`
}
