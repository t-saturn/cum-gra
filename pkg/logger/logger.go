package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()

	// Crear carpeta logs si no existe
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, os.ModePerm)
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
