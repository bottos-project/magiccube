package data

import (
	log "github.com/cihub/seelog"
	"strings"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/bottos-project/bottos/service/common/bean"
	user_proto "github.com/bottos-project/bottos/service/user/proto"
	"errors"
	"fmt"
)

const (
	BASE_URL = "http://139.217.206.43:8080/rpc"
	TX_PARAMS = "service=core&method=CoreApi.PushTrx&request=%s"
)

func BlockHeader() (*user_proto.BlockHeader, error) {
	params := `service=core&method=CoreApi.QueryChainInfo&request={}`
	resp, err := http.Post(BASE_URL, "application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if (resp.StatusCode != 200) {
		return nil, errors.New(string(body))
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var common_ret = &bean.CoreCommonReturn{}
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

func PushTransaction (i interface{}) (interface{}, error) {
	r, err := json.Marshal(i)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info(fmt.Sprintf(TX_PARAMS, string(r)))
	resp, err := http.Post(BASE_URL, "application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf(TX_PARAMS, string(r))))
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
	var common_ret = &bean.CoreCommonReturn{}
	err = json.Unmarshal(body, common_ret)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if common_ret.Errcode == 0 {
		return common_ret, nil
	}

	return nil, errors.New(string(body))
}

func AccountInfo(account string) (*user_proto.AccountInfoData, error) {
	params := `service=core&method=CoreApi.QueryAccount&request={"account_name":"%s"}`
	resp, err := http.Post(BASE_URL, "application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf(params, string(account))))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if (resp.StatusCode != 200) {
		return nil, errors.New(string(body))
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var common_ret = &bean.CoreCommonReturn{}
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
