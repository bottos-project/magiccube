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
	"log"

	"strconv"
	"time"

	"github.com/bottos-project/magiccube/service/storage/controller"
	"github.com/bottos-project/magiccube/service/storage/internal/service"
)

func InsertTransaction(tx interface{}) error {
	fmt.Println("insert transaction")
	return nil
}
func GetSyncedBlockCount(stat service.StateRepository) uint64 {

	fmt.Println("find blocknumber")
	var syncedNumber uint64
	syncedNumber, err := stat.CallGetSyncBlockCount()
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
		return 0
	}
	return syncedNumber
}

func Sync(syncedNumber uint64, latestBlock uint64, c chan int) {
	if syncedNumber == 0 {
		syncedNumber++
	}
	for i := syncedNumber; i <= latestBlock; i++ {
		num := strconv.FormatUint(i, 10)
		fmt.Println(num, latestBlock)
		mBlock, err := controller.GetBlock(num)
		if err != nil {
			log.Fatal(err)
			break
		}
		if err := InsertTransaction(mBlock); err != nil {
			log.Fatal(err)
			break
		}
	}

	c <- 1
}

func StartSync(stat service.StateRepository) {

	sync := make(chan int, 1)
	go Sync(GetSyncedBlockCount(stat), GetLatestBlockNumber(), sync)

	// 周期同步
	for {
		select {
		case <-sync:
			log.Println("syncing task is completed.")
			time.Sleep(7 * time.Second) // TODO: using event listen
			Sync(GetSyncedBlockCount(stat), GetLatestBlockNumber(), sync)
		}
	}
}
