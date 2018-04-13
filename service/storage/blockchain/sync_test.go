package blockchain

import (
	"fmt"
	"testing"
)

func TestSync(t *testing.T) {
	sync := make(chan int, 1)
	go Sync(0, 12, sync)
	fmt.Println("sync")
}
