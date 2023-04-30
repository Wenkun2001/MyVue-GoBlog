package flag

import sys_flag "flag"

type Option struct {
	DB bool
}

// Parse 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("db", true, "初始化数据库")
	return Option{
		DB: *db,
	}
}

func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	return false
}

func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
	}
}
