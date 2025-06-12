package grpc

import (
	"context"
	"time"

	"github.com/t-saturn/central-user-manager/server/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	start := time.Now()

	var ip string
	if p, ok := peer.FromContext(ctx); ok {
		ip = p.Addr.String()
	}

	resp, err = handler(ctx, req)

	duration := time.Since(start)
	st, _ := status.FromError(err)

	logger.Log.WithFields(map[string]interface{}{
		"method":   info.FullMethod,
		"duration": duration,
		"status":   st.Code().String(),
		"ip":       ip,
	}).Info("gRPC request")

	return resp, err
}
