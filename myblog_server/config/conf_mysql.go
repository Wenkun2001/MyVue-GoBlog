package config

import (
	"strconv"
	"time"
)

type Mysql struct {
	Host     string        `yaml:"host"`
	Port     int           `yaml:"port"`
	Config   string        `yaml:"config"`
	DB       string        `yaml:"db"`
	User     string        `yaml:"user"`
	Password string        `yaml:"password"`
	logLevel string        `yaml:"log_level"`
	Idle     int           `yaml:"max_idle_conn"`
	Conn     int           `yaml:"max_open_conn"`
	Lifetime time.Duration `yaml:"Conn_max_lifetime"`
}

func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}
