package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Host string `json:"host"`
	DB   struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	}
	Log struct {
		Level string `json:"level"`
	}
}

var AppConfig *Config

func Init() {
	envName := os.Getenv("env")
	fileName := fmt.Sprintf("env/config.%s.json", envName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open config file: %v\n", err)
		panic(err)
	}
	defer file.Close()

	AppConfig = &Config{}
	if err := json.NewDecoder(file).Decode(AppConfig); err != nil {
		fmt.Printf("Failed to decode config JSON: %v\n", err)
		panic(err)
	}
	fmt.Printf("Loaded config: %+v\n", AppConfig)
}
