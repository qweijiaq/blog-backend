package config

type Logger struct {
	Level string `yaml:"level"`
	Prefix string `yaml:"prefix"`
	Director string `yaml:"director"`
	Show_line bool `yaml:"show_line"`			// 是否显示行号
	Log_in_console bool `yaml:"log_in_console"` // 是否显示打印的路径
}