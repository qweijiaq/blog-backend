package models

type TagModel struct {
	MODEL `json:","`
	Title string `gorm:"size:30" json:"title"`
}
