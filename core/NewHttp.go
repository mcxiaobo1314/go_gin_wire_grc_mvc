package core

import (
	"context"
	"fmt"
	"gin_gorm/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type HServer struct {
	R           *gin.Engine
	AppConfig   *config.Config
	HGrpcServer *HGrpcServer
	hGrpcClient *hGrpcClient
}

func NewHServer(R *gin.Engine, AppConfig *config.Config, HGrpcServer *HGrpcServer, hGrpcClient *hGrpcClient) *HServer {
	return &HServer{
		R:           R,
		AppConfig:   AppConfig,
		HGrpcServer: HGrpcServer,
		hGrpcClient: hGrpcClient,
	}
}

// ------http 启动server 服务--------
func (s *HServer) Start() {

	srv := &http.Server{
		Addr:    s.AppConfig.App.Port,
		Handler: s.R,
	}

	go func() {
		fmt.Println("http server starting port " + s.AppConfig.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server starting listen error: %s\n", err)
		}
	}()

	s.HGrpcServer.Start()
	s.hGrpcClient.Start()
	close(srv)
}

// 关闭http server 服务
func close(srv *http.Server) {
	//建立1个缓冲区的信号通道
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutdown http server ....")
	log.Println("shutdown grpc server ....")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown ", err)
	}
	log.Printf("server exiting")
}
