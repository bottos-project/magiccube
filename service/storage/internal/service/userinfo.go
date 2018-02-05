package service

//import (
//	"fmt"
//)

type UserInfo struct {
	AccountName string `json:"accountname"`
	PublicKey   int64  `json:"publickey"`
	UserType    string `json:"usertype"`
	RoleType    string `json:"roletype"`
	Privacy     string `json:"Privacy"`
}
