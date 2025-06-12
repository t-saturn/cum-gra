package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logFile, err := os.OpenFile("logs/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Warn("Could not open log file, using only stdout")
		Log.SetOutput(os.Stdout)
		return
	}

	multi := io.MultiWriter(os.Stdout, logFile)
	Log.SetOutput(multi)
	Log.SetLevel(logrus.DebugLevel)
}
