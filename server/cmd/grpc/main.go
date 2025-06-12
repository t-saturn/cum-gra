package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	"github.com/t-saturn/central-user-manager/server/internal/grpc"
	userpb "github.com/t-saturn/central-user-manager/server/pb/user"
	"github.com/t-saturn/central-user-manager/server/pkg/database"
	"github.com/t-saturn/central-user-manager/server/pkg/logger"
	"google.golang.org/grpc"
)

func main() {
	// Cargar env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando .env")
	}

	logger.InitLogger()
	logger.Log.Info("Iniciando servidor gRPC...")

	database.InitDatabase()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log.Fatalf("Error al escuchar: %v", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, grpc.NewUserHandler())

	logger.Log.Info("Servidor gRPC escuchando en puerto 50051")
	if err := s.Serve(lis); err != nil {
		logger.Log.Fatalf("Error al servir: %v", err)
	}
}
