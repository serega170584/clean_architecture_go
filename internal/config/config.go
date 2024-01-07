package config

import (
	"encoding/json"
	"log"
	"os"
)

type AppConfig struct {
	Name       string `json:"name"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	AdminLogin string `json:"admin_login"`
	AdminToken string `json:"admin_token"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	App *AppConfig `json:"app"`
	DB  *DBConfig  `json:"db"`
}

func New() *Config {
	jsonBytes, err := os.ReadFile("config/config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	var config *Config
	err = json.Unmarshal([]byte(jsonBytes), &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	return config
}
