package config

import (
	"log"
	"os"
)

type Config struct {
	AppPort    string `json:"app_port"`
	PostgreDSN string `json:"postgre_dsn"`
}

func LoadConfig() (config *Config) {
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "1323"
	}

	postgreDSN := os.Getenv("POSTGRE_DSN")
	if postgreDSN == "" {
		log.Fatal("PostgreDSN not set")
	}

	return &Config{
		AppPort:    appPort,
		PostgreDSN: postgreDSN,
	}
}
