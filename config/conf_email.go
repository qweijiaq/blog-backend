package config

type Email struct {
	Host             string `json:"host" yaml:"host"`
	Port             int    `json:"port" yaml:"port"`
	User             string `json:"user" yaml:"user"` // 发件人邮箱
	Password         string `json:"password" yaml:"password"`
	DefaultFromEmail string `json:"default_from_email" yaml:"default_from_email"` // 默认发件人
	UseSSL           bool   `json:"use_ssl" yaml:"use_ssl"`                       // 是否使用 SSL
	UseTLS           bool   `json:"use_tls" yaml:"use_tls"`                       // 是否使用 TLS
}
