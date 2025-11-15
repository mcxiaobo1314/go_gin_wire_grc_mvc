package core

import (
	"fmt"
	"gin_gorm/config"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type hGrpcClient struct {
	AppConfig *config.Config
}

func NewHGrpcClient(AppConfig *config.Config) *hGrpcClient {
	return &hGrpcClient{
		AppConfig: AppConfig,
	}
}

// ------grpc 启动client 客户端--------
func (g *hGrpcClient) Start() {
	if g.AppConfig.Grpc.ClientAddr != "" {
		fmt.Println("grpc client connenct " + g.AppConfig.Grpc.ClientAddr)
		grpcClient, err := grpc.Dial(g.AppConfig.Grpc.ClientAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("grpc client dial not connect: %v", err)
		}
		defer grpcClient.Close()
		config.InitGrpcClient(grpcClient)
	}

}
