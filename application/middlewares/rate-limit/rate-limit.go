package rate_limit

import (
	"context"
	"google.golang.org/grpc"
)

func LimitRateUnary(ll interface{}) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//check number of requests of client ip

		//pass to server

		return handler(ctx, req)
	}
}

func LimitRateStream(ll interface{}) grpc.StreamServerInterceptor {

	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		//check number of requests of client ip

		//pass to server

		return handler(srv, ss)
	}
}
