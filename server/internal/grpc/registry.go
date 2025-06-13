package grpc

import (
	"google.golang.org/grpc"

	"github.com/t-saturn/central-user-manager/server/internal/repositories"
	"github.com/t-saturn/central-user-manager/server/internal/services"

	applicationpb "github.com/t-saturn/central-user-manager/server/pb/proto/application"
	userpb "github.com/t-saturn/central-user-manager/server/pb/proto/user"

	applicationhandler "github.com/t-saturn/central-user-manager/server/internal/handlers"
	userhandler "github.com/t-saturn/central-user-manager/server/internal/handlers"
)

func RegisterServices(s *grpc.Server) {
	/* User */
	userRepo := repositories.NewUserRepository()
	userSvc := services.NewUserService(userRepo)
	userpb.RegisterUserServiceServer(s, userhandler.NewUserHandler(userSvc))

	/* Application */
	appRepo := repositories.NewApplicationRepository()
	appSvc := services.NewApplicationService(appRepo)
	applicationpb.RegisterApplicationServiceServer(s, applicationhandler.NewApplicationHandler(appSvc))
}
