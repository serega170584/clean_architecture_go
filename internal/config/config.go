package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type Config struct {
	App *AppConfig `json:"application"`
}

func New() *Config {
	jsonBytes, err := os.ReadFile("config/config.json")
	if err != nil {
		panic(err.Error())
	}
	var config *Config
	err = json.Unmarshal([]byte(jsonBytes), &config)
	if err != nil {
		panic(err.Error())
	}

	return config
}
