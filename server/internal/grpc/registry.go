package grpc

import (
	"google.golang.org/grpc"

	"github.com/t-saturn/central-user-manager/server/internal/repositories"
	"github.com/t-saturn/central-user-manager/server/internal/services"
	userpb "github.com/t-saturn/central-user-manager/server/pb/proto"
)

func RegisterServices(s *grpc.Server) {
	userRepo := repositories.NewUserRepository()
	userSvc := services.NewUserService(userRepo)
	userpb.RegisterUserServiceServer(s, NewUserHandler(userSvc))

	// Futuro: Role, System, etc.
	// roleRepo := repositories.NewRoleRepository()
	// roleSvc := services.NewRoleService(roleRepo)
	// rolepb.RegisterRoleServiceServer(s, NewRoleHandler(roleSvc))
}
