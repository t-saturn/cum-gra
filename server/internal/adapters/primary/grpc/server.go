package grpc

import (
	"log"
	"net"

	"github.com/t-saturn/central-user-manager/server/internal/adapters/primary/grpc/handlers"
	userpb "github.com/t-saturn/central-user-manager/server/internal/adapters/primary/grpc/proto/user"
	"google.golang.org/grpc"
)

func StartGrpcServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Error al iniciar el listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	userpb.RegisterUserServiceServer(grpcServer, handlers.NewUserGrpcHandler())

	log.Println("🚀 Servidor gRPC corriendo en puerto 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Error al levantar el servidor gRPC: %v", err)
	}
}
