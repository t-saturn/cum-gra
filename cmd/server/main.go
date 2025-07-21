package main

import (
	"github.com/t-saturn/auth-service-server/pkg/logger"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Starting authentication service...")
}
