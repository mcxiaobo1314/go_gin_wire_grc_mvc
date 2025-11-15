package config

import (
	"gin_gorm/grpc_client"
	"gin_gorm/grpc_proto"
	"gin_gorm/grpc_server"

	"google.golang.org/grpc"
)

//protoc version: 3.20.1
//生成grpc代码命令：
//protoc --go_out=./grpc_proto --go-grpc_out=./grpc_proto ./grpc_proto/hello.proto

func InitGrpcServer(grpcServer *grpc.Server) *grpc.Server {
	grpc_proto.RegisterHelloServiceServer(grpcServer, &grpc_server.HelloGrpcServer{})
	return grpcServer
}

func InitGrpcClient(grpcClient *grpc.ClientConn) *grpc.ClientConn {
	grpc_client.Conn(grpcClient)
	return grpcClient
}
