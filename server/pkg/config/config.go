package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
	GRPCPort string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró .env, cargando variables del entorno del sistema")
	}

	cfg := &Config{
		DBHost:   os.Getenv("db_host"),
		DBPort:   os.Getenv("db_port"),
		DBUser:   os.Getenv("db_user"),
		DBPass:   os.Getenv("db_pass"),
		DBName:   os.Getenv("db_name"),
		GRPCPort: os.Getenv("grpc_port"),
	}

	return cfg
}
