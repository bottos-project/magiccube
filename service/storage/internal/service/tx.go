package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/bottos-project/magiccube/service/storage/proto"
	"github.com/bottos-project/magiccube/service/storage/util"
)

type Transaction struct {
	Txid          int64  `json:"txid"`
	Blockid       int64  `json:"blockid"`
	FirstParty    string `json:"firstparty"`
	SecondParty   string `json:"secondparty"`
	Witness       string `json:"witness"`
	Time          int64  `jsons:"time"`
	Signature     string `json:"signature"`
	Price         string `json:"price"`
	AssetId       string `json:"assetid"`
	RequirementId string `json:"requirementid"`
	status        string `json:"status"`
}

func (c *StorageService) GetTx(ctx context.Context, request *storage.Request, response *storage.Response) error {

	if request == nil {
		//	response.Message = "para error"
		return nil //errors.BadRequest("", "Missing storage request")
	}
	fmt.Print(request)
	url := "http:"
	//url, err := c.storageRepo.GetTx(request.txid, request.account)
	//if err != nil
	{
		//	response.Message = "get url failed"
		return nil //errors.InternalServerError("", "Failed get put url: %s", err.Error())

	}
	fmt.Print(url)
	//todo
	//c.dbRepo.CallGetTx("")
	return nil
}

func (c *StorageService) InsertTx(ctx context.Context, request *storage.Request, response *storage.Response) error {

	if request == nil {
		//	response.Message = "para error"
		return nil //errors.BadRequest("", "Missing storage request")
	}
	fmt.Print(request)
	url := "http:"
	//url, err := c.storageRepo.GetTx(request.txid, request.account)
	//if err != nil
	{
		//	response.Message = "get url failed"
		return nil //errors.InternalServerError("", "Failed get put url: %s", err.Error())

	}
	fmt.Print(url)
	//response.Message = "OK"
	return nil
}

func (c *StorageService) GetRecentTxList(ctx context.Context, request *storage.RecentTxListRequest, response *storage.RecentTxListResponse) error {
	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetRecentTxList")
	txs, err := c.mgoRepo.CallGetRecentTxList()
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}
	if txs == nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed CallGetRecentTxList")
	}
	response.RecentTxList = []*storage.RecentTx{}
	for _, tx := range txs {
		dbTag := &storage.RecentTx{tx.TransactionID,
			tx.From,
			tx.To,
			tx.Price,
			tx.Type,
			tx.Date,
			tx.BlockId}
		response.RecentTxList = append(response.RecentTxList, dbTag)
	}
	response.Code = 1
	return nil
}
func (c *StorageService) GetUserTxList(ctx context.Context, request *storage.UserRequest, response *storage.UserTxListResponse) error {
	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetRecentTxList")
	txs, err := c.mgoRepo.CallGetUserTxList(request.Username)
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}
	if txs == nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed CallGetRecentTxList")
	}
	response.RecentTxList = []*storage.RecentTx{}
	for _, tx := range txs {
		dbTag := &storage.RecentTx{tx.TransactionID,
			tx.From,
			tx.To,
			tx.Price,
			tx.Type,
			tx.Date,
			tx.BlockId}
		response.RecentTxList = append(response.RecentTxList, dbTag)
	}
	response.Code = 1
	return nil
}
func (c *StorageService) GetSumTxAmount(ctx context.Context, request *storage.AllRequest, response *storage.SumTxAmountResponse) error {
	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetSumTxAmount")
	sum, err := c.mgoRepo.CallGetSumTxAmount()
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed CallGetSumTxAmount")

	}
	response.AmountSum = sum
	response.Code = 1
	return nil
}
func (c *StorageService) GetAllTxNum(ctx context.Context, request *storage.AllRequest, response *storage.AllTxNumResponse) error {
	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetTxNum")
	sum, err := c.mgoRepo.CallGetAllTxNum()
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed CallGetAllTxNum")

	}
	response.TxNum = sum
	response.Code = 1
	return nil
}
func (c *StorageService) GetTxNumByDay(ctx context.Context, request *storage.AllRequest, response *storage.DayTxNumResponse) error {
	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetTxNumByDay")
	begin, end := util.YesterdayDur()

	txnum, err := c.mgoRepo.CallGetTxNumByDay(begin, end)
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}
	response.DayTxNum = txnum
	response.Code = 1
	return nil
}
func (c *StorageService) GetTxNumByWeek(ctx context.Context, request *storage.AllRequest, response *storage.WeekTxNumResponse) error {
	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	response.WeekTxNum = make([]uint64, 1, 7)
	days := util.WeekDur()
	for _, day := range days {
		txnum, err := c.mgoRepo.CallGetTxNumByDay(day.Begin, day.End)
		if err != nil {
			response.Code = 0
			fmt.Println(err)
			return errors.New("Failed CallGetTxNumByDay")
		}
		response.WeekTxNum = append(response.WeekTxNum, txnum)
	}
	response.Code = 1
	return nil
}

func (c *StorageService) GetRecentTransferList(ctx context.Context, request *storage.AllRequest, response *storage.TransferListResponse) error {
	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetRecentTransfersList")
	txs, err := c.mgoRepo.CallGetRecentTransfersList()
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}
	if txs == nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed CallGetRecentTxList")
	}
	response.TransferList = []*storage.Transfer{}
	for _, tx := range txs {
		dbTag := &storage.Transfer{tx.TransactionID,
			tx.TxTime,
			tx.Price,
			tx.From,
			tx.To,
			tx.BlockNum}
		response.TransferList = append(response.TransferList, dbTag)
	}
	response.Code = 1
	return nil
}
