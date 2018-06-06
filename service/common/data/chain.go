/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Service Layer
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

package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bottos-project/magiccube/config"
	"github.com/bottos-project/magiccube/service/common/bean"
	user_proto "github.com/bottos-project/magiccube/service/user/proto"
	log "github.com/cihub/seelog"
)

const (
	// BASE_URL BASE_URL
	BASE_URL = config.BASE_RPC
	// TX_PARAMS TX_PARAMS
	TX_PARAMS = "service=bottos&method=CoreApi.PushTrx&request=%s"
)

// BlockHeader get block header
func BlockHeader() (*user_proto.BlockHeader, error) {
	params := `service=bottos&method=CoreApi.QueryChainInfo&request={}`
	resp, err := http.Post(BASE_URL, "application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		log.Error(resp.Status)
		return nil, errors.New(string(body))
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var commonRet = &bean.CoreBaseReturn{}
	err = json.Unmarshal(body, commonRet)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if commonRet.Errcode != 0 {
		return nil, errors.New(string(body))
	}

	resultBuf, err := json.Marshal(commonRet.Result)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var blockHeader = &user_proto.BlockHeader{}
	err = json.Unmarshal(resultBuf, blockHeader)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return blockHeader, nil
}

// PushTransaction push transaction
func PushTransaction(i interface{}) (*bean.CoreCommonReturn, error) {
	var params = ""
	switch i.(type) {
	case string:
		params = fmt.Sprintf(TX_PARAMS, i.(string))
	default:
		r, err := json.Marshal(i)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		params = fmt.Sprintf(TX_PARAMS, string(r))
	}
	log.Info(params)
	resp, err := http.Post(BASE_URL, "application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	log.Info("body:", string(body))
	var commonRet bean.CoreCommonReturn
	err = json.Unmarshal(body, &commonRet)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if commonRet.Errcode == 0 {
		return &commonRet, nil
	}
	return nil, errors.New(string(body))
}

// AccountInfo get account info
func AccountInfo(account string) (*user_proto.AccountInfoData, error) {
	params := `service=bottos&method=CoreApi.QueryAccount&request={"account_name":"%s"}`
	resp, err := http.Post(BASE_URL, "application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf(params, string(account))))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Info(string(body))
	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var commonRet = &bean.CoreBaseReturn{}
	err = json.Unmarshal(body, commonRet)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if commonRet.Errcode != 0 {
		return nil, errors.New(string(body))
	}

	resultBuf, err := json.Marshal(commonRet.Result)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var accountInfo = &user_proto.AccountInfoData{}
	err = json.Unmarshal(resultBuf, accountInfo)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return accountInfo, nil
}
