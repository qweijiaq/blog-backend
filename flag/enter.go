package flag

import (
	sys_flag "flag"
	"github.com/fatih/structs"
)

type Option struct {
	DB   bool
	User string // -u superuser 创建超级用户 | -u user 创建普通用户
	ES   string // -es create | -es delete
}

// Parse 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户(string 为 user 或 superuser)")
	es := sys_flag.String("es", "", "es操作(string 为 create 或 delete)")
	// 解析命令行参数写入注册的 flag 里
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
		ES:   *es,
	}
}

// 是否停止 Web 项目
func IsWebStop(option Option) (f bool) {
	maps := structs.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}
	}
	return f
}

// 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}
	if option.User == "superuser" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	//sys_flag.Usage()
	if option.ES == "create" {
		EsCreateIndex()
	}
}
