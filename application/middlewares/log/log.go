package log

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

func LogUnaryServerInterceptor(ll interface{}) grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		logRequest(ctx, req, info, ll)
		start := time.Now()
		defer logResponse(req, info, start, ll, resp, err)
		return handler(ctx, req)
	}
}

func logResponse(req interface{}, info *grpc.UnaryServerInfo, start time.Time, ll interface{}, resp interface{}, err error) {

}

func logRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, ll interface{}) {

}
