package main

import (
	"os"
	"time"
	//	"github.com/bottos-project/bottos/service/storage/blockchain"
	"github.com/bottos-project/bottos/service/storage/internal/platform/config"
	"github.com/bottos-project/bottos/service/storage/internal/platform/minio"
	"github.com/bottos-project/bottos/service/storage/internal/platform/mongodb"
	"github.com/bottos-project/bottos/service/storage/internal/service"
	"github.com/bottos-project/bottos/service/storage/proto"
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
	baseConfig "github.com/bottos-project/bottos/config"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	svc := micro.NewService(
		micro.Name(config.ServiceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version(config.Version),
	)
	svc.Init()
	minioStorageRepository := minio.NewMinioRepository(baseConfig.BASE_MINIO_ADDR, baseConfig.BASE_MINIO_ACCESS_KEY, baseConfig.BASE_MINIO_SECRET_KEY)
	mgoRepository := mongodb.NewMongoRepository(baseConfig.BASE_MONGODB_ADDR)

	repo := service.NewStorageService(minioStorageRepository, mgoRepository)

	storage.RegisterStorageHandler(svc.Server(), repo)
	//blockchain.StartSync(stateRepository)
	//blockchain.LoopAging(repo)
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
