package grpc

import (
	context "context"

	pb "github.com/t-saturn/central-user-manager/server/pb/user"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Ping(ctx context.Context, in *pb.Empty) (*pb.Pong, error) {
	return &pb.Pong{Message: "pong"}, nil
}
