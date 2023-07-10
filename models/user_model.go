package models

import (
	"backend/models/ctype"
)

// AuthModel 用户表
type UserModel struct {
	MODEL
	NickName   string           `gorm:"size:36" json:"nick_name,select(c)"`
	UserName   string           `gorm:"size:36" json:"user_name"`
	Password   string           `gorm:"size:128" json:"-"`
	Avatar     string           `gorm:"size:256" json:"avatar,select(c)"`
	Email      string           `gorm:"size:256" json:"email"`
	Tel        string           `gorm:"size:18" json:"tel"`
	Addr       string           `gorm:"size:64" json:"addr,select(c)"`
	Token      string           `gorm:"size:64" json:"token"`
	IP         string           `gorm:"size:20" json:"ip,select(c)"`
	Role       ctype.Role       `gorm:"size:4;default:1" json:"role"`
	SignStatus ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`
}

// func ParseRole(role Role) string {
// 	switch role {
// 	case PermissionAdmin:
// 		return "管理员"
// 	case PermissionUser:
// 		return "普通用户"
// 	case PermissionVisitor:
// 		return "游客"
// 	case PermissionDisabledUser:
// 		return "被禁言用户"
// 	}
// 	return "其他"
// }
