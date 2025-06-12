package main

import (
	"fmt"
	"log"
	"net"

	"github.com/t-saturn/central-user-manager/server/pkg/config"
	"github.com/t-saturn/central-user-manager/server/pkg/database"
	"google.golang.org/grpc"
)

func main() {
	// Cargar configuración
	cfg := config.Load()

	// Inicializar base de datos
	database.Init(cfg)

	// Iniciar listener gRPC
	addr := fmt.Sprintf(":%s", cfg.GRPCPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to start listener in %s: %v", addr, err)
	}

	// Crear servidor gRPC
	s := grpc.NewServer()

	log.Printf("User management gRPC server started in %s", addr)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
