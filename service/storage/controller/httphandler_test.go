package controller

import (
	"testing"
	"fmt"
	"github.com/bottos-project/magiccube/service/storage/util"
)


func TestGetInfo(t *testing.T) {
	var rsp *util.Info
	rsp, _= GetInfo()
	fmt.Println(rsp.HeadBlockNum)
	fmt.Println(rsp.ServerVersion)
}
func TestGetBlock(t *testing.T) {
	var rsp *util.Block
	rsp, _= GetBlock("0000000445a9f27898383fd7de32835d5d6a978cc14ce40d9f327b5329de796b")
	fmt.Println(rsp.Producer)
	fmt.Println(rsp.ID)
}
func TestGetBlockNum1(t *testing.T) {
	var rsp *util.Block
	rsp, _= GetBlock("1")
	fmt.Println(rsp.Producer)
	fmt.Println(rsp.ID)
}
func TestGetAccountInfo(t *testing.T) {
	var rsp *util.AccountInfo
	rsp, _= GetAccountInfo()
	fmt.Println(rsp.AccountName)
}
func TestGetTxInfo(t *testing.T) {
	var rsp *util.TxInfo
	rsp, _= GetTxInfo()
	fmt.Println(rsp.TransactionID)
	fmt.Println(rsp.Transaction.Expiration)
}
func TestGetCodeInfo(t *testing.T) {
	fmt.Println("ddd")
}