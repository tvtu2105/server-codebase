package services

import (
	"context"
	"sever-codebase/proto/server-proto/pb"
)

func (s *services) HelloWorld(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}
