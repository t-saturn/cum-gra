package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/t-saturn/central-user-manager/server/internal/models"
	userpb "github.com/t-saturn/central-user-manager/server/pb/proto"
	"github.com/t-saturn/central-user-manager/server/pkg/config"
	"github.com/t-saturn/central-user-manager/server/pkg/database"
	"github.com/t-saturn/central-user-manager/server/pkg/logger"

	handler "github.com/t-saturn/central-user-manager/server/internal/grpc"
	"github.com/t-saturn/central-user-manager/server/internal/repositories"
	"github.com/t-saturn/central-user-manager/server/internal/services"
)

func main() {
	// Inicializar logger
	logger.InitLogger()

	// Cargar configuración
	cfg := config.Load()

	// Inicializar base de datos
	database.Init(cfg)

	// Migrar automáticamente el modelo User
	_ = database.DB.AutoMigrate(&models.User{})

	// Iniciar listener gRPC
	addr := fmt.Sprintf(":%s", cfg.GRPCPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Log.Fatalf("Failed to start listener in %s: %v", addr, err)
	}

	// Crear servidor gRPC
	s := grpc.NewServer()

	repo := repositories.NewUserRepository()
	svc := services.NewUserService(repo)
	userpb.RegisterUserServiceServer(s, handler.NewUserHandler(svc))

	logger.Log.Infof("User management gRPC server started in %s", addr)

	if err := s.Serve(lis); err != nil {
		logger.Log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
