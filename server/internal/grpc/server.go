package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/t-saturn/central-user-manager/server/pkg/config"
	"github.com/t-saturn/central-user-manager/server/pkg/database"
	"github.com/t-saturn/central-user-manager/server/pkg/logger"

	"github.com/t-saturn/central-user-manager/server/internal/models"
)

func StartGRPCServer(cfg *config.Config) {
	database.Init(cfg)
	_ = database.DB.AutoMigrate(&models.User{})

	addr := fmt.Sprintf(":%s", cfg.GRPCPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Log.Fatalf("Failed to start listener in %s: %v", addr, err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(LoggingInterceptor),
	)

	RegisterServices(s)

	logger.Log.Infof("User management gRPC server started in %s", addr)

	if err := s.Serve(lis); err != nil {
		logger.Log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
