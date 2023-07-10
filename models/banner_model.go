package models

import (
	"backend/global"
	"backend/models/ctype"
	"os"

	"gorm.io/gorm"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                // 图片路径
	Hash      string          `json:"hash"`                // 图片的 hash 值，用于判断是否重复
	Name      string          `gorm:"size:36" json:"name"` // 图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"`
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		err = os.Remove(b.Path[1:])
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
