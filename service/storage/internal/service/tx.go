package service

//import (
//	"fmt"
//)

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
