package grpc_server

import (
	"context"
	"gin_gorm/grpc_proto"
	"log"
)

type HelloGrpcServer struct {
	grpc_proto.UnimplementedHelloServiceServer
}

func (s *HelloGrpcServer) SayHello(ctx context.Context, req *grpc_proto.HelloRequest) (*grpc_proto.HelloResponse, error) {
	log.Printf("Received: %v", req.GetName())
	return &grpc_proto.HelloResponse{Message: "Hello " + req.GetName()}, nil
}
