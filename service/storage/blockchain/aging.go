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
