package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1 // QQ登录
	SignGitee SignStatus = 2 // Gitee 登录
	SignEmail SignStatus = 3 // 邮箱登录
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ登录"
	case SignGitee:
		str = "Gitee 登录"
	case SignEmail:
		str = "邮箱登录"
	default:
		str = "其他"
	}
	return str
}
