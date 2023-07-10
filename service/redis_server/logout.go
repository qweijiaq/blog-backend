package redis_server

import (
	"backend/global"
	"backend/utils"
	"time"
)

const prefix = "logout_"

// 针对注销的操作
func Logout(token string, diff time.Duration) error {
	return global.Redis.Set(prefix+token, "", diff).Err()
}

func CheckLogut(token string) bool {
	keys := global.Redis.Keys(prefix + "*").Val()
	if utils.InList(prefix+token, keys) {
		return true
	}
	return false
}
