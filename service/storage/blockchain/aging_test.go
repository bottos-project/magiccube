package blockchain

import (
	"fmt"
	"testing"
)

func TestAging(t *testing.T) {
	c := make(chan int, 1)
	go Sync(0, 12, c)
	fmt.Println("sync")
}
