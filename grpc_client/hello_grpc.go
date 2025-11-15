package grpc_client

import (
	"context"
	"gin_gorm/grpc_proto"
	"log"

	"google.golang.org/grpc"
)

func Conn(grpcClient *grpc.ClientConn) {
	helloClient := grpc_proto.NewHelloServiceClient(grpcClient)
	var rsp *grpc_proto.HelloResponse
	rsp, _ = helloClient.SayHello(context.Background(), &grpc_proto.HelloRequest{Name: "test"})
	log.Println(rsp)
}
