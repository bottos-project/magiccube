package blockchain

import (
	"fmt"
	//	"log"
	//	"time"

	//"github.com/bottos-project/magiccube/service/storage/controller"
	"github.com/bottos-project/magiccube/service/storage/internal/service"
	//	"github.com/bottos-project/magiccube/service/storage/util"
)

func tokenAging(timeout int64, stat service.StateRepository, c chan int) {
	fmt.Println("okkkk")
	stat.CallTokenAging(timeout)
	c <- 1
}

//func LoopAging(stat service.SqliteRepository) {
//	channel := make(chan int, 1)
//	go tokenAging(util.DefualtAgingTime, stat, channel)

//	// 周期
//	for {
//		select {
//		case <-channel:
//			log.Println("syncing task is completed.")
//			time.Sleep(5 * time.Second) // TODO: using event listen
//			tokenAging(util.DefualtAgingTime, stat, channel)
//		}
//	}
//}
