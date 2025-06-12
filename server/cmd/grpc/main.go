package main

import (
	"github.com/t-saturn/central-user-manager/server/internal/grpc"
	"github.com/t-saturn/central-user-manager/server/pkg/config"
	"github.com/t-saturn/central-user-manager/server/pkg/logger"
)

func main() {
	logger.InitLogger()

	cfg := config.Load()

	grpc.StartGRPCServer(cfg)
}
