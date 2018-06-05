/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Data Exchange Client
  Created by Developers Team of Bottos.

  This program is free software: you can distribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with Bottos. If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"os"
	"time"
	//	"github.com/bottos-project/magiccube/service/storage/blockchain"
	baseConfig "github.com/bottos-project/magiccube/config"
	"github.com/bottos-project/magiccube/service/storage/internal/platform/config"
	"github.com/bottos-project/magiccube/service/storage/internal/platform/minio"
	"github.com/bottos-project/magiccube/service/storage/internal/platform/mongodb"
	"github.com/bottos-project/magiccube/service/storage/internal/service"
	"github.com/bottos-project/magiccube/service/storage/proto"
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
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
