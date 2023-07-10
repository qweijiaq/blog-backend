package models

import "time"

// 记录用户什么时候收藏了什么文章
type UserCollectsModel struct {
	UserID    uint      `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID string    `gorm:"size:32"`
	CreatedAt time.Time
}
