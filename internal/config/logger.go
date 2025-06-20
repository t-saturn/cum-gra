package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()

	// Crear carpeta de logs si no existe
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, os.ModePerm)
	}

	// Archivo con fecha como nombre
	fileName := time.Now().Format("2006-01-02") + ".log"
	logFilePath := filepath.Join(logDir, fileName)
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Logger.Out = os.Stdout
		Logger.Warn("No se pudo abrir archivo de log, usando stdout")
	} else {
		Logger.SetOutput(file)
	}

	// Formato JSON para mejor integrabilidad (puede cambiarse a TextFormatter si prefieres)
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     false,
	})

	// Nivel de log mínimo
	Logger.SetLevel(logrus.DebugLevel)

	// Incluir nombre de archivo y línea
	Logger.SetReportCaller(true)
}
