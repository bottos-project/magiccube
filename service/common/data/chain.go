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
	log "github.com/cihub/seelog"
	"strings"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/bottos-project/magiccube/service/common/bean"
	user_proto "github.com/bottos-project/magiccube/service/user/proto"
	"errors"
	"fmt"
	"github.com/bottos-project/magiccube/config"
)

const (
	BASE_URL  = config.BASE_RPC
	TX_PARAMS = "service=bottos&method=CoreApi.PushTrx&request=%s"
)

// get block header
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

	if (resp.StatusCode != 200) {
		log.Error(resp.Status)
		return nil, errors.New(string(body))
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var common_ret = &bean.CoreBaseReturn{}
	err = json.Unmarshal(body, common_ret)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if common_ret.Errcode != 0 {
		return nil, errors.New(string(body))
	}

	result_buf, err := json.Marshal(common_ret.Result)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var block_header = &user_proto.BlockHeader{}
	err = json.Unmarshal(result_buf, block_header)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return block_header, nil
}

// push transaction
func PushTransaction (i interface{}) (*bean.CoreCommonReturn, error) {
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

	if (resp.StatusCode != 200) {
		return nil, errors.New(string(body))
	}
	log.Info("body:", string(body))
	var common_ret bean.CoreCommonReturn
	err = json.Unmarshal(body, &common_ret)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if common_ret.Errcode == 0 {
		return &common_ret, nil
	}
	return nil, errors.New(string(body))
}

// get account info
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
	if (resp.StatusCode != 200) {
		return nil, errors.New(string(body))
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var common_ret = &bean.CoreBaseReturn{}
	err = json.Unmarshal(body, common_ret)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if common_ret.Errcode != 0 {
		return nil, errors.New(string(body))
	}

	result_buf, err := json.Marshal(common_ret.Result)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var account_info = &user_proto.AccountInfoData{}
	err = json.Unmarshal(result_buf, account_info)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return account_info, nil
}
