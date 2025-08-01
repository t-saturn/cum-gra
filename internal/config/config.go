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
	JWTSecret     string
	JWTExpMinutes string
	ServerPort    string
}

// App Configuración - NUEVO
type AppConfig struct {
	Version string
}

// Estructura principal que agrupa todas las configuraciones
type Config struct {
	Postgres         PostgresConfig
	Mongo            MongoConfig
	Server           ServerConfig
	App              AppConfig         // NUEVO
	ExternalServices map[string]string // NUEVO
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
			JWTSecret:     getEnv("JWT_SECRET", "my_secret_key"),
			JWTExpMinutes: getEnv("JWT_EXP_MINUTES", "15"),
			ServerPort:    getEnv("SERVER_PORT", "8000"),
		},
		// NUEVO: Configuración de la aplicación
		App: AppConfig{
			Version: getEnv("APP_VERSION", "1.0.0"),
		},
		// NUEVO: Servicios externos
		ExternalServices: loadExternalServices(),
	}
}

// Retorna la configuración cargada
func GetConfig() Config {
	return cfg
}

// NUEVA: Carga los servicios externos desde variables de entorno
func loadExternalServices() map[string]string {
	services := make(map[string]string)

	// Servicios por defecto
	defaultServices := map[string]string{
		"auth-api":     "http://auth-api:8080/health",
		"user-service": "http://user-service:8080/health",
	}

	// Cargar servicios por defecto
	for name, url := range defaultServices {
		services[name] = url
	}

	// Sobrescribir con variables de entorno si existen
	if authAPI := os.Getenv("AUTH_API_HEALTH_URL"); authAPI != "" {
		services["auth-api"] = authAPI
	}
	if userService := os.Getenv("USER_SERVICE_HEALTH_URL"); userService != "" {
		services["user-service"] = userService
	}

	// Patrón genérico para agregar más servicios
	// EXTERNAL_SERVICE_[NOMBRE]_HEALTH_URL
	// for _, env := range os.Environ() {
	// 	// Puedes implementar lógica más compleja aquí si necesitas
	// 	// detectar automáticamente servicios desde variables de entorno
	// }

	return services
}

// Utilidad interna para leer variables de entorno con valor por defecto
func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	logger.Log.Infof("Environment variable %s not set, using default value: %s", key, fallback)
	return fallback
}
