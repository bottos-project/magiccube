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
)

const (
	BASE_URL = "http://139.217.206.43:8080/rpc"
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
