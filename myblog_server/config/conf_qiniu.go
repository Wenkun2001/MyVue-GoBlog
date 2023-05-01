package config

type QiNiu struct {
	Enable     string  `json:"enable" yaml:"enable"`
	AccessKey  string  `json:"access_key" yaml:"access_key"`
	SecreteKey string  `json:"secrete_key" yaml:"secrete_key"`
	Bucket     string  `json:"bucket" yaml:"bucket"` // 存储桶的名字
	CDN        string  `json:"cdn" yaml:"cdn"`       // 访问图片的地址的前缀
	Zone       string  `json:"zone" yaml:"zone"`     //存储的地区
	size       float64 `json:"size" yaml:"size"`     // 存储的大小限制，单位为MB
}
