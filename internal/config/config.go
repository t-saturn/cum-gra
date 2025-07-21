package config

import (
	"os"

	"github.com/spf13/viper"
	"github.com/t-saturn/auth-service-server/pkg/logger"
)

// Config contiene la configuración global de la aplicación cargada desde variables de entorno.
type Config struct {
	MONGO_URI      string
	MONGO_DB_NAME  string
	PORT           string
	JWT_SECRET     string
	JWT_EXP_MINUTES string
}

var cfg Config

// LoadConfig carga la configuración de la aplicación desde variables de entorno.
func LoadConfig() {
	viper.AutomaticEnv()

	cfg = Config{
		MONGO_URI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MONGO_DB_NAME:  getEnv("MONGO_DB_NAME", "mongo_db"),
		PORT:           getEnv("PORT", "8000"),
		JWT_SECRET:     getEnv("JWT_SECRET", "my_secret_key"),
		JWT_EXP_MINUTES: getEnv("JWT_EXP_MINUTES", "15"),
	}
}

// GetConfig devuelve la configuración global cargada de la aplicación.
func GetConfig() Config {
	return cfg
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	logger.Log.Infof("Environment variable %s not set, using default value: %s", key, fallback)
	return fallback
}
