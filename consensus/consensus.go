package consensus

import (
	"BOT/common/log"
	"fmt"
	"time"
)

type ConsensusService interface {
	Start() error
	Halt() error
}

func Log(message string) {
	logMsg := fmt.Sprintf("[%s] %s", time.Now().Format("08/06/2017 16:06:06"), message)
	fmt.Println(logMsg)
	log.Info(logMsg)
}
