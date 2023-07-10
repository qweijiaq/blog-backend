package models

import "backend/models/ctype"

// 菜单表
type MenuModel struct {
	MODEL
	Title        string        `gorm:"size:32" json:"title"`
	Path         string        `gorm:"size:32" json:"path"`
	Slogan       string        `gorm:"size:64" json:"slogan"`
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`
	AbstractTime int           `json:"abstract_time"`
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;JoinReferences:BannerID" json:"banners"`
	BannerTime   int           `gorm:"size:32" json:"banner_time"`
	Sort         int           `gorm:"size:10" json:"sort"`
}
