package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primaryKey" json:"id,select($any)"`
	CreatedAt time.Time `json:"created_at,select($any)"`
	UpdatedAt time.Time `json:"-"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type ESIDListRequest struct {
	IDList []string `json:"id_list" binding:"required"`
}
