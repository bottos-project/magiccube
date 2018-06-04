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
package controller
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/bottos-project/magiccube/service/storage/util"
	baseConfig "github.com/bottos-project/magiccube/config"
)

var (
	serverurl= baseConfig.BASE_CHAIN_URL
)

//https://github.com/ethereum/wiki/wiki/JSON-RPC

func SetServer(newServer string) {
	serverurl = newServer
}
func GetInfo()(*util.Info,error){
	resp, err := http.Get("http://"+serverurl+"/v1/chain/get_info")
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jResp := new(util.Info)

	err = json.Unmarshal(body, jResp)
	if err != nil {
		return nil, err
	}

	return jResp,nil
}
func GetBlock(num_or_id string)(*util.Block,error){
	body := strings.NewReader(`{"block_num_or_id":`+num_or_id+`}`)
	req, err := http.NewRequest("POST", "http://"+serverurl+"/v1/chain/get_block", body)
	if err != nil {
		return nil,err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	block := new(util.Block)

	err = json.Unmarshal(respBody, block)
	if err != nil {
		return nil, err
	}

	return block,nil
}
func GetAccountInfo()(*util.AccountInfo,error){

	body := strings.NewReader(`{"account_name":"testa"}`)
	req, err := http.NewRequest("POST", "http://"+serverurl+"/v1/chain/get_account", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	account := new(util.AccountInfo)

	err = json.Unmarshal(respBody, account)
	if err != nil {
		return nil, err
	}

	return account,nil
}
func GetTxInfo()(*util.TxInfo,error){
	body := strings.NewReader(`{"transaction_id":"06ffce7503d82a4e19bd7cdfb9c507c5c3c40fda3bd316ee35f344d42807db6e"}`)
	req, err := http.NewRequest("POST", "http://"+serverurl+"/v1/account_history/get_transaction", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tx := new(util.TxInfo)

	err = json.Unmarshal(respBody, tx)
	if err != nil {
		return nil, err
	}

	return tx,nil

}


func GetCodeInfo()(string, error){
	body := strings.NewReader(`{"account_name":"currency"}`)
	req, err := http.NewRequest("POST", "http://"+serverurl+"/v1/chain/get_code", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	return "",nil
}
