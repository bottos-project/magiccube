package service

//import (
//	"fmt"
//)

type ContractInfo struct {
	AddressId int64  `json:"addressid"`
	Name      string `json:"name"`
	Deployer  int64  `json:"deployer"`
	Register  int64  `json:"register"`
}
