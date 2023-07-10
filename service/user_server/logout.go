package user_server

import (
	"backend/service/redis_server"
	"backend/utils"
	"time"
)

func (UserService) Logout(claims *utils.CustomClaims, token string) error {
	// 计算距离现在的过期时间
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_server.Logout(token, diff)
}
