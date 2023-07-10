package flag

import (
	"backend/global"
	"backend/models/ctype"
	"backend/service/user_server"
	"fmt"
)

func CreateUser(permission string) {
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)

	fmt.Printf("请输入用户名: ")
	fmt.Scan(&userName)
	fmt.Printf("请输入昵称: ")
	fmt.Scan(&nickName)
	fmt.Printf("请输入密码: ")
	fmt.Scan(&password)
	fmt.Printf("请重新输入密码: ")
	fmt.Scan(&rePassword)
	fmt.Printf("请输入邮箱: ")
	fmt.Scanln(&email)

	// 检验密码
	if password != rePassword {
		global.Log.Error("两次密码不一致，请重新输入")
		return
	}

	role := ctype.PermissionUser
	if permission == "superuser" {
		role = ctype.PermissionAdmin
	}

	err := user_server.UserService{}.CreateUser(userName, nickName, password, role, email, "127.0.0.1")
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("创建用户%s并入库成功", userName)
}
