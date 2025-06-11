package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/server/internal/adapters/primary/grpc/proto/user"
	"github.com/t-saturn/central-user-manager/server/internal/adapters/secondary/persistence/postgres/repositories"
	"github.com/t-saturn/central-user-manager/server/internal/core/domain/entities"
	"github.com/t-saturn/central-user-manager/server/internal/core/services"
	"github.com/t-saturn/central-user-manager/server/pkg/utils"
)

type UserGrpcHandler struct {
	user.UnimplementedUserServiceServer
	service *services.UserService
}

func NewUserGrpcHandler() *UserGrpcHandler {
	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)
	return &UserGrpcHandler{service: service}
}

func (h *UserGrpcHandler) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	userEntity := &entities.User{
		Name:     req.Name,
		LastName: req.Last_name,
		UserName: req.User_name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := h.service.CreateUser(userEntity); err != nil {
		return nil, err
	}

	return &user.CreateUserResponse{
		Id:      uuid.New().String(),
		Message: "Usuario creado correctamente",
	}, nil
}
