package common

import (
	"backend/global"
	"backend/models"
	"fmt"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
	Likes []string // 模糊匹配的字段
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc"
	}

	DB = DB.Where(model)

	for index, column := range option.Likes {
		if index == 0 {
			DB.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			continue
		}
		DB.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
	}

	// 分页
	count = DB.Where(model).Find(&list).RowsAffected
	// 这里的 query 会受上面查询的影响，需要手动复位
	query := DB.Where(model)
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
