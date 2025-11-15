package core

import (
	"fmt"
	"gin_gorm/config"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HGrpcServer struct {
	AppConfig *config.Config
}

func NewHGrpcServer(AppConfig *config.Config) *HGrpcServer {
	return &HGrpcServer{
		AppConfig: AppConfig,
	}
}

// ------grpc 启动server 服务--------
func (g *HGrpcServer) Start() {
	if g.AppConfig.Grpc.ServerPort != "" {
		listen, err := net.Listen(g.AppConfig.Grpc.Protocol, g.AppConfig.Grpc.ServerPort)
		if err != nil {
			log.Fatalf("grpc server failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		grpcServer = config.InitGrpcServer(grpcServer)

		go func() {
			fmt.Println("grpc server starting " + g.AppConfig.Grpc.Protocol + " port " + g.AppConfig.Grpc.ServerPort)
			if err := grpcServer.Serve(listen); err != nil {
				log.Fatalf("grpc server starting listen: %s\n", err)
			}
		}()
	}
}
