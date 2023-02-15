package services

import (
	"context"
	"sever-codebase/application/models"
	"sever-codebase/proto/server-proto/pb"
)

type services struct {
	Kafka *models.KafkaConfig
	JWT   *models.JWT
	Cfg   *models.Config
}

func NewServer(cfg *models.Config) pb.BaseServerServer {
	return &services{Cfg: cfg}
}

type Server interface {
	HelloWorld(ctx context.Context, request *pb.Request) (*pb.Response, error)
}
