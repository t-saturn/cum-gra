package config

import (
	"os"

	"github.com/spf13/viper"
	"github.com/t-saturn/auth-service-server/pkg/logger"
)

// PostgreSQL Configuración
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// MongoDB Configuración
type MongoConfig struct {
	URI    string
	DBName string
}

// JWT + Server Configuración
type ServerConfig struct {
	JWTPrivateKeyPath string
	JWTPublicKeyPath  string
	JWTExpMinutes     string
	JWTAlg            string
	JWTKid            string
	JWTIss            string
	JWKSURL           string
	JWKSMaxAge        string
	Port              string
	AppLandingURL     string
}

// App Configuración - NUEVO
type AppConfig struct {
	Version string
}

// Estructura principal que agrupa todas las configuraciones
type Config struct {
	Postgres PostgresConfig
	Mongo    MongoConfig
	Server   ServerConfig
	App      AppConfig // NUEVO
}

var cfg Config

// Carga toda la configuración de variables de entorno
func LoadConfig() {
	viper.AutomaticEnv()

	cfg = Config{
		Postgres: PostgresConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			DBName:   getEnv("DB_NAME", "postgres_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Mongo: MongoConfig{
			URI:    getEnv("MONGO_URI", "mongodb://localhost:27017"),
			DBName: getEnv("MONGO_DB_NAME", "mongo_db"),
		},
		Server: ServerConfig{
			JWTPrivateKeyPath: getEnv("JWT_PRIVATE_KEY_PATH", "./keys/jwtRS256.key"),
			JWTPublicKeyPath:  getEnv("JWT_PUBLIC_KEY_PATH", "./keys/jwtRS256.key.pub"),
			JWTExpMinutes:     getEnv("JWT_EXP_MINUTES", "15"),
			JWTAlg:            getEnv("JWT_ALG", "RS256"),
			JWTKid:            getEnv("JWT_KID", "rrhh-sso"),
			JWTIss:            getEnv("JWT_ISS", "http://localhost:9190"),
			JWKSMaxAge:        getEnv("JWKS_MAX_AGE", "60"),
			JWKSURL:           getEnv("JWKS_URL", "http://localhost:9190/.well-known/jwks.json"),
			Port:              getEnv("SERVER_PORT", "9190"),
			AppLandingURL:     getEnv("APP_LANDING_URL", "http://localhost:9190"),
		},
		// NUEVO: Configuración de la aplicación
		App: AppConfig{
			Version: getEnv("APP_VERSION", "1.0.0"),
		},
	}
}

// Retorna la configuración cargada
func GetConfig() Config {
	return cfg
}

// Utilidad interna para leer variables de entorno con valor por defecto
func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	logger.Log.Infof("Environment variable %s not set, using default value: %s", key, fallback)
	return fallback
}
