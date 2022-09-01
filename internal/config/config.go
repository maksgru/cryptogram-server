package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	LogLevel  string `yaml:"log_level"`
	APIServer struct {
		BindAddr string `yaml:"bind_addr"`
	} `yaml:"api_server"`
	PostgresConfig struct {
		Host     string `yaml:"host"`
		Dbname   string `yaml:"dbname"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"postgres_config"`
}

func GetConfig() *Config {
	var instance Config
	var configFile, err = os.ReadFile("configs/apiserver.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(configFile, &instance)
	if err != nil {
		log.Fatal(err)
	}

	return &instance
}
