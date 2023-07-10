package flag

import (
	"backend/global"
	"backend/models"
	"backend/plugins/log"
)

func Makemigrations() {
	var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollects{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	err = global.DB.Set("gorm:table_options", "EnGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.CategoryModel{},
			&models.MessageModel{},
			&models.UserModel{},
			&models.CommentModel{},
			//&models.ArticleModel{},
			&models.MenuModel{},
			&models.UserCollectsModel{},
			&models.MenuBannerModel{},
			&models.FeedbackModel{},
			&models.LoginDataModel{},
			&models.ChatModel{},
			&log.LogModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success] 生成数据库表结构成功")
}
