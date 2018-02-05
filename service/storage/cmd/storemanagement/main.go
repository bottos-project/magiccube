package main

import (
	"os"
	"time"
	"github.com/code/bottos/service/storemanagement/proto"
	"github.com/code/bottos/service/storemanagement/internal/platform/config"
	//	"github.com/code/bottos/service/storemanagement/internal/platform/redis"
	"github.com/code/bottos/service/storemanagement/internal/service"
	//	"github.com/code/bottos/service/storemanagement/proto"
	"github.com/code/bottos/service/storemanagement/internal/platform/minio"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	//	"github.com/micro/go-web"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	svc := grpc.NewService(
		micro.Name(config.ServiceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version(config.Version),
	)
	svc.Init()

	minioStorageRepository := minio.NewMinioRepository("https://127.0.0.1:9000", "aaaaa", "bbbb")
	storemanagement.RegisterStoragemanagementHandler(svc.Server(), service.NewStorageService(minioStorageRepository))

	if err := svc.Run(); err != nil {
		panic(err)
	}
}
