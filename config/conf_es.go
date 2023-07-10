package config

import "fmt"

type ES struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Config   string `json:"config"` // 高级配置，如 charset
	DB       string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
	LogLevel string `json:"log_level"` // 日志等级，debug 为输出全部 sql, dev, release
}

func (es ES) URL() string {
	return fmt.Sprintf("%s:%d", es.Host, es.Port)
}
