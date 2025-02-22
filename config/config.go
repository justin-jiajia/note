package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type Config struct {
	DBType      string
	DBPath      string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	Environment string
	ServerPort  string
	AllowOrigin string
}

func NewConfig() *Config {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal("Failed to load config file: ", err)
	}

	return &Config{
		DBType:      cfg.Section("database").Key("type").String(),
		DBPath:      cfg.Section("database").Key("path").String(),
		DBHost:      cfg.Section("database").Key("host").String(),
		DBPort:      cfg.Section("database").Key("port").String(),
		DBUser:      cfg.Section("database").Key("user").String(),
		DBPassword:  cfg.Section("database").Key("password").String(),
		DBName:      cfg.Section("database").Key("name").String(),
		Environment: cfg.Section("server").Key("environment").MustString("development"),
		ServerPort:  cfg.Section("server").Key("port").MustString("8080"),
		AllowOrigin: cfg.Section("server").Key("allow_origin").MustString("*"),
	}
}
