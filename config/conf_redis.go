package config

import "fmt"

type Redis struct {
	IP       string `yaml:"ip" json:"ip"`
	Port     int    `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
	PoolSize int    `yaml:"pool_size" json:"pool_size"`
}

func (r Redis) Addr() string {
	return fmt.Sprintf("%s:%d", r.IP, r.Port)
}
