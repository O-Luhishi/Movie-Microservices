package config

import "time"

type Conf struct {
	Server serverConf
	Db     dbConf
	Debug  bool
}

type dbConf struct {
	Port     int    `env:"DB_PORT,required"`
	Host     string `env:"DB_HOST,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DbName   string `env:"DB_NAME,required"`
}

type serverConf struct {
	Port         string
	TimeoutRead  time.Duration
	TimeoutWrite time.Duration
	TimeoutIdle  time.Duration
	Certificate  string
	Key          string
}

func NewAppConfig() *Conf {
	var c = Conf{
		Server: serverConf{
			Port:         ":8080",
			TimeoutRead:  5,
			TimeoutWrite: 180,
			TimeoutIdle:  15,
			Certificate:  "localhost.crt",
			Key:          "localhost.key",
		},
		Debug: true,
	}
	return &c
}
