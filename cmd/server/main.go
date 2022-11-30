package main

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"sever-codebase/application/configs"
	"sever-codebase/application/middlewares/jwt"
	"sever-codebase/application/middlewares/log"
	rate_limit "sever-codebase/application/middlewares/rate-limit"
)

func main() {
	env, err := configs.LoadEnv()
	cfg, err := configs.LoadConfig("conf/", env)
	if err != nil {
		return
	}
	fmt.Println("Config: ", cfg)

	//init jwt
	s := grpc.NewServer(initStreamOptions(), initUnaryOptions())
	fmt.Println(s)
	//init service

	// init grpc server

	//init middleware

	//limit resource

	//force shutdown

	//start server

}

func initStreamOptions() grpc.ServerOption {
	return grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_auth.StreamServerInterceptor(jwt.Authenticate),
		rate_limit.LimitRateStream(nil),
	))
}

func initUnaryOptions() grpc.ServerOption {
	return grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		log.LogUnaryServerInterceptor(nil),
		grpc_auth.UnaryServerInterceptor(jwt.Authenticate),
		rate_limit.LimitRateUnary(nil),
	))
}
