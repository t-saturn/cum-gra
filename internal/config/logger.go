package config

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()

	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, os.ModePerm)
	}

	logFileName := time.Now().Format("2006-01-02") + ".log"
	logPath := filepath.Join(logDir, logFileName)

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Logger.Out = os.Stdout
		Logger.Warn("Could not write log file, using stdout only")
	} else {
		multiWriter := io.MultiWriter(os.Stdout, file)
		Logger.SetOutput(multiWriter)
	}

	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     false,
	})

	Logger.SetLevel(logrus.DebugLevel)

	Logger.SetReportCaller(true)
}
