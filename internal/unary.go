package internal

import (
	"context"

	"google.golang.org/grpc"
)

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logger("[request log] method: %v, req: %v", info.FullMethod, req)

	m, err := handler(ctx, req)
	if err != nil {
		logger("RPC failed with error %v", err)
	}
	return m, err
}
