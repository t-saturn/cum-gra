// Package logger proporciona una configuración global del sistema de logs para la aplicación.
package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

// Log es el logger global utilizado por toda la aplicación, configurado con salida a consola y archivo.
var Log *logrus.Logger

// InitLogger inicializa el logger global con formato, nivel de log y salida a archivo + consola.
func InitLogger() {
	Log = logrus.New()

	// Crear carpeta logs si no existe
	logDir := "logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		logrus.Fatalf("No se pudo crear el directorio de logs: %v", err)
	}

	// Nombre del archivo: logs/YYYY-MM-DD.log
	logFile := filepath.Join(logDir, time.Now().Format("2006-01-02")+".log")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		Log.SetOutput(os.Stdout)
		Log.Warn("No se pudo abrir el archivo de log, se usará solo salida estándar")
	} else {
		// Salida múltiple: archivo + consola
		multiWriter := io.MultiWriter(os.Stdout, file)
		Log.SetOutput(multiWriter)
	}

	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Log.SetLevel(logrus.DebugLevel)
}
