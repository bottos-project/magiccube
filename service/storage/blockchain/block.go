package blockchain

import (

	"github.com/code/bottos/service/storage/controller"
)

func GetLatestBlockNumber() uint64 {
	blockInfo,err := controller.GetInfo()
	if err != nil {
		return 0
	}
	latestBlock := blockInfo.HeadBlockNum
	return latestBlock
}