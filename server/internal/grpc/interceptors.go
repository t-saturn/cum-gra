package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/t-saturn/central-user-manager/server/pkg/logger"
)

func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	start := time.Now()

	// Ejecutar el handler
	resp, err = handler(ctx, req)

	// Log luego de la ejecución
	duration := time.Since(start)
	st, _ := status.FromError(err)

	logger.Log.WithFields(map[string]interface{}{
		"method":   info.FullMethod,
		"duration": duration,
		"status":   st.Code().String(),
	}).Info("gRPC request")

	return resp, err
}
