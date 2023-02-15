package main

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"net"
	"sever-codebase/application/middlewares/jwt"
	"sever-codebase/application/middlewares/log"
	rate_limit "sever-codebase/application/middlewares/rate-limit"
	"sever-codebase/application/models"
	"sever-codebase/application/services"
	"sever-codebase/proto/server-proto/pb"
)

func registerServices(config *models.Config) pb.BaseServerServer {
	return services.NewServer(config)
}

func StartGrpcServer(cfg *models.Config) {
	//init servers
	s := grpc.NewServer(initStreamOptions(), initUnaryOptions())

	//init services
	svc := registerServices(cfg)

	// register grpc server
	pb.RegisterBaseServerServer(s, svc)

	//start server
	//ll.Info("GRPC server start listening", l.Int("GRPC address", cfg.GRPCAddress))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.GrpcPort))
	if err != nil {
		//ll.Fatal("error listening to address", l.Int("address", cfg.Server.GrpcPort), l.Error(err))
		return
	}

	if err = s.Serve(listener); err != nil {
		//ll.Fatal("error serve ", l.Error(err))
		return
	}
}

func initStreamOptions() grpc.ServerOption {
	//create grpc server with jwt and limit rate which are a middleware
	return grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_auth.StreamServerInterceptor(jwt.Authenticate),
		rate_limit.LimitRateStream(nil),
	))
}

func initUnaryOptions() grpc.ServerOption {
	//create http server with jwt, limit rate and log which are a middleware
	return grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		log.LogUnaryServerInterceptor(nil),
		grpc_auth.UnaryServerInterceptor(jwt.Authenticate),
		rate_limit.LimitRateUnary(nil),
	))
}
