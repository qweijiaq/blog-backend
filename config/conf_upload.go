package config

type Upload struct {
	Size int    `json:"size" yaml:"size"` // 图片上传大小的上限
	Path string `json:"path" yaml:"path"` // 图片上传的目录
}
