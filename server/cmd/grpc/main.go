package main

import (
	"fmt"
	"net"

	"github.com/t-saturn/central-user-manager/server/pkg/config"
	"github.com/t-saturn/central-user-manager/server/pkg/database"
	"github.com/t-saturn/central-user-manager/server/pkg/logger"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	logger.InitLogger()

	database.Init(cfg)

	addr := fmt.Sprintf(":%s", cfg.GRPCPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Log.Fatalf("Failed to start listener in %s: %v", addr, err)
	}

	s := grpc.NewServer()

	logger.Log.Infof("User management gRPC server started in %s", addr)

	if err := s.Serve(lis); err != nil {
		logger.Log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
