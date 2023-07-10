package user_server

import (
	"backend/global"
	"backend/models"
	"backend/models/ctype"
	"backend/utils"
	"errors"
)

const Avatar = "/uploads/avatar/default.jpeg"

func (UserService) CreateUser(userName string, nickName string, password string, role ctype.Role, email string, ip string) error {
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		return errors.New("用户名已存在")
	}
	// 加密
	hsahPwd := utils.HashPwd(password)

	addr := utils.GetAddr(ip)

	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hsahPwd,
		Email:      email,
		Role:       role,
		Avatar:     Avatar,
		IP:         ip,
		Addr:       addr,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
