package handlers

import (
	"context"

	"github.com/t-saturn/central-user-manager/server/internal/services"
	userpb "github.com/t-saturn/central-user-manager/server/pb/proto/user"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	Service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) FnCreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user, err := h.Service.CreateUser(req.GetName(), req.GetEmail())
	if err != nil {
		return nil, err
	}
	return &userpb.CreateUserResponse{Id: user.ID.String()}, nil
}
