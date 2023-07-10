package models

import "backend/models/ctype"

type LoginDataModel struct {
	MODEL
	UserID    uint             `json:"user_id"`
	UserModel UserModel        `gorm:"foreignKey:UserID" json:"-"`
	IP        string           `gorm:"size:20" json:"ip"`
	Nickname  string           `gorm:"size:42" json:"nickname"`
	Token     string           `gorm:"size:256" json:"token"`
	Device    string           `gorm:"size:256" json:"device"`
	Addr      string           `gorm:"size:64" json:"addr"`
	LoginType ctype.SignStatus `gorm:"size:type=smallint(6)" json:"login_type"`
}
