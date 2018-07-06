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

package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
)

//Sha256 sha256
func Sha256(msg []byte) []byte {
	sha := sha256.New()
	sha.Write([]byte(hex.EncodeToString(msg)))
	return sha.Sum(nil)
}

func HexIptoDec(slaveIP string) string {
	j := 0
	node1 := string([]rune(slaveIP)[8*j:8*j+2])

	ip1, _ := strconv.ParseInt(node1, 16, 32)
	node2 := string([]rune(slaveIP)[8*j+2:8*j+4])
	ip2, _ := strconv.ParseInt(node2, 16, 32)
	node3 := string([]rune(slaveIP)[8*j+4:8*j+6])
	ip3, _ := strconv.ParseInt(node3, 16, 32)
	node4 := string([]rune(slaveIP)[8*j+6:8*j+8])
	ip4, _ := strconv.ParseInt(node4, 16, 32)
	node := strconv.FormatInt(ip1, 10) + "." + strconv.FormatInt(ip2, 10) + "." + strconv.FormatInt(ip3, 10) + "." + strconv.FormatInt(ip4, 10)
	return node
}

//Country struct
type Country struct {
	Country   string
	Province  string
	City      string
	Latitude  float32
	Longitude float32
}

//szTongS struct
type szTongS struct {
	Code uint64
	Data Country
}

// CityInfo struct
type CityInfo struct {
	Pointx string
	Pointy string
}

// CountryDetails struct
type CountryDetails struct {
	City CityInfo
}
type SzTongSpoint struct {
	Detail CountryDetails
}

//ip2pointxy ip2pointxy
func Ip2pointxy(ip string) SzTongSpoint {
	fmt.Println(ip)
	var infos szTongS
	var infosPoint SzTongSpoint

	url := "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/jsonp")
	resp, err := client.Do(req)

	if err != nil {
		return SzTongSpoint{}
	}

	defer resp.Body.Close()
	respBody, ee := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
	if ee == nil && respBody != nil {

		err = json.Unmarshal(respBody, &infos)
		if err != nil {
			return SzTongSpoint{}
		} else if infos.Data.City == "XX" {
			pointx, pointy := useInternationalPolicy(ip)
			fmt.Println("pointx", pointx)
			infosPoint.Detail.City.Pointx = pointx
			infosPoint.Detail.City.Pointy = pointy

			if pointx == "" || pointy == "" {
				return SzTongSpoint{}
			}
			return infosPoint

		}

	} else {
		fmt.Println("Error when Unmarshal infos!")
		return SzTongSpoint{}
	}

	url = "http://apis.map.qq.com/jsapi?qt=poi&wd=" + infos.Data.City + "&pn=0&rn=10&rich_source=qipao&rich=web&nj=0&c=1&key=FBOBZ-VODWU-C7SVF-B2BDI-UK3JE-YBFUS&output=jsonp&pf=jsapi&ref=jsapi&cb=qq.maps._svcb3.search_service_0"
	req, _ = http.NewRequest("GET", url, nil)
	resp, err = client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return SzTongSpoint{}
	}

	defer resp.Body.Close()
	respBody, ee = ioutil.ReadAll(resp.Body)

	indexS := strings.Index(string(respBody), "(") + 2
	indexE := strings.Index(string(respBody), ")")
	jsonstr := string(respBody)[indexS:indexE]

	if ee == nil && respBody != nil {
		err = json.Unmarshal([]byte(jsonstr), &infosPoint)
		if err != nil {
			return SzTongSpoint{}
		}
	} else {
		fmt.Println("Error when Unmarshal infosPoint!")
		return SzTongSpoint{}
	}

	return infosPoint
}

func useInternationalPolicy(ipaddr string) (string, string) {
	client := &http.Client{}
	url := "http://api.ipinfodb.com/v3/ip-city/?key=8ae944829cc080834bb7ee22638f1c474dd64db171dcc6e567ab7d312f365926&ip=" + ipaddr
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return "", ""
	}

	respBody, ee := ioutil.ReadAll(resp.Body)
	if ee == nil && respBody != nil {
		s := strings.Split(string(respBody), ";")

		//fmt.Println(s[len(s)-3:len(s)-1])
		pointx, pointy := s[len(s)-3], s[len(s)-2]
		return pointx, pointy
	}

	return "", ""

}
