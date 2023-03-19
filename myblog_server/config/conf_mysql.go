package config

type Mysql struct {
	host      string `yaml:"host"`
	port      int    `yaml:"port"`
	db        string `yaml:"db"`
	user      string `yaml:"user"`
	password  string `yaml:"password"`
	log_level string `yaml:"log_level"`
}
