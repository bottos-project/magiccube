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

package api

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	//"github.com/micro/go-micro"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
	"unsafe"

	"github.com/bottos-project/magiccube/service/node/config"
	"github.com/bottos-project/magiccube/service/storage/util"
	//slog "github.com/cihub/seelog"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//NodeInfo struct
type NodeInfo struct {
	NodeName        string
	IpAddr          string
	UserName        string
	PassWord        string
	BtoPort         string
	BtoUser         string
	BtoPath         string
	WalletIP        string
	KeyPath         string
	ProdUser        string
	DbUser          string
	DbPass          string
	StorageSize     string
	StoragePath     string
	ServPath        string
	ServLst         []string
	SeedIp          string
	SlaveIpLst      []string
	StorageCapacity string
}

//NodeInfos struct
type NodeInfos struct {
	Node []NodeInfo
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

type szTongSpoint struct {
	Detail CountryDetails
}

// MongoRepository struct
type MongoRepository struct {
	mgoEndpoint string
}

// MongoContext struct
type MongoContext struct {
	mgoSession *mgo.Session
}

// Ippointxy struct
type Ippointxy struct {
	ID        bson.ObjectId `bson:"_id"`
	Ip        string        `bson:"ip"`
	Pointx    string        `bson:"pointx"`
	Pointy    string        `bson:"pointy"`
	CreatedAt time.Time     `bson:"createdAt"`
}

// StorageDBClusterInfo struct
type StorageDBClusterInfo struct {
	SeedIP  string `json:"nodeIP"` //This is node's self ip, in future, seedip should be added here.
	SlaveIP string `json:"clusterIP"`
	NodeUUID string `json:"uuid"`
}

// NodeCapacityInfo struct
type NodeCapacityInfo struct {
	NodeUUID        string `json:nodeuuid`
	NodeIP          string `json:nodeip`
	StorageCapacity string `json:storagecapacity`
}

func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// ReadFile ReadFile
func ReadFile(filename string) NodeInfos {
	if filename == "" {
		fmt.Println("Error ! parmeter is null")
		return NodeInfos{}
	}
	var ni NodeInfos

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return NodeInfos{}
	}

	str := string(bytes)

	if err := json.Unmarshal([]byte(str), &ni); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return NodeInfos{}
	}

	return ni
}

// WriteFile WriteFile
func WriteFile(filename string, nodeInfos NodeInfos) error {

	jsonData, errs := json.Marshal(nodeInfos)
	if errs != nil {
		log.Println("errs occurs! Marshal failed. filename:", filename)
		return errs
	}
	errs = ioutil.WriteFile(filename, jsonData, os.ModeAppend)
	return errs
}

// MonitorConfigFile MonitorConfigFile
func MonitorConfigFile(configFile string) error {

	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watch.Close()

	//var command string

	err = watch.Add(configFile)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case ev := <-watch.Events:
			{
				if ev.Op&fsnotify.Create == fsnotify.Create {
					log.Println("Monitored file had been created ! : ", ev.Name)
				}
				if ev.Op&fsnotify.Write == fsnotify.Write {
					//log.Println("Monitored file had been modified ! : ", ev.Name);
					nodeInfos := ReadFile(config.CONFIG_FILE)
					if unsafe.Sizeof(nodeInfos) == 0 {
						//if getting empty value from json file , skip it
						continue
					}

					for i := 0; i < len(nodeInfos.Node); i++ {
						//fmt.Println(node_infos.Node[i].IpAddr)
						for j := 0; j < len(nodeInfos.Node[i].ServLst); j++ {
							//command = "echo \""+node_infos.Node[i].PassWord+"\" | sudo -S echo \""+nodeInfos.Node[i].ServLst[j]+"\" > /etc/rc.local"
							//fmt.Println("command = ",command)

							if /*buf*/ _, err := SshCommand(nodeInfos.Node[i].UserName,
								nodeInfos.Node[i].PassWord,
								nodeInfos.Node[i].IpAddr,
								22,
								"ls /"); err == nil {
								//fmt.Println("buf = ",buf)
							}

						}
					}
				}
				if ev.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("Monitored file had been deleted ! : ", ev.Name)
				}
				if ev.Op&fsnotify.Rename == fsnotify.Rename {
					log.Println("Monitored file had been created ! : ", ev.Name)
				}
				if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
					log.Println("Monitored file's authority had been Chmod ! : ", ev.Name)
				}
			}
		case err := <-watch.Errors:
			{
				log.Println("error : ", err)
				return err
			}
		}
	}
	// return nil
}

//Connect about SSH
func Connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

// SshCommand SshCommand
func SshCommand(user, password, host string, port int, command string) (string, error) {
	session, err := Connect(user, password, host, port)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	var buf bytes.Buffer

	//session.Stdout = os.Stdout
	session.Stdout = &buf
	session.Stderr = os.Stderr

	if err = session.Run(command); err != nil {
		//session.Close()
		return "", err
	}

	//session.Close()
	return buf.String(), err
}

// GetAccountInfo GetAccountInfo
func GetAccountInfo(url, username string) (*util.AccountInfo, error) {

	body := strings.NewReader("{\"account_name\":\"" + username + "\"}")
	//req, err := http.NewRequest("POST", "http://"+config.SERV_URL+"/v1/chain/get_account", body)
	req, err := http.NewRequest("POST", "http://"+url+"/v1/chain/get_account", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}

	if resp != nil {
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
		return account, nil
	}

	return nil, err

}

// PathExist PathExist
func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func multiIp2pointxyint(iplist []string) map[string]CityInfo {
	m := make(map[string]CityInfo)
	
	for _, ip := range iplist {
		infosPoint := ip2pointxy(ip)
		if (infosPoint != szTongSpoint{}) {
			m[ip] = infosPoint.Detail.City
		}
	}
	return m
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

//ip2pointxy ip2pointxy
func ip2pointxy(ip string) szTongSpoint {
	var infos szTongS
	var infosPoint szTongSpoint
	url := "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/jsonp")
	resp, err := client.Do(req)

	if err != nil {
		return szTongSpoint{}
	}

	defer resp.Body.Close()
	respBody, ee := ioutil.ReadAll(resp.Body)
	if ee == nil && respBody != nil {
		err = json.Unmarshal(respBody, &infos)

		if err != nil {
			return szTongSpoint{}
		} else if infos.Data.City == "XX" {
			pointx, pointy := useInternationalPolicy(ip)
			infosPoint.Detail.City.Pointx = pointx
			infosPoint.Detail.City.Pointy = pointy

			if pointx == "" || pointy == "" {
				return szTongSpoint{}
			}
			return infosPoint

		}

	} else {
		fmt.Println("Error when Unmarshal infos!")
		return szTongSpoint{}
	}

	url = "http://apis.map.qq.com/jsapi?qt=poi&wd=" + infos.Data.City + "&pn=0&rn=10&rich_source=qipao&rich=web&nj=0&c=1&key=FBOBZ-VODWU-C7SVF-B2BDI-UK3JE-YBFUS&output=jsonp&pf=jsapi&ref=jsapi&cb=qq.maps._svcb3.search_service_0"
	req, _ = http.NewRequest("GET", url, nil)
	resp, err = client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return szTongSpoint{}
	}

	defer resp.Body.Close()
	respBody, ee = ioutil.ReadAll(resp.Body)

	indexS := strings.Index(string(respBody), "(") + 2
	indexE := strings.Index(string(respBody), ")")
	jsonstr := string(respBody)[indexS:indexE]

	if ee == nil && respBody != nil {
		err = json.Unmarshal([]byte(jsonstr), &infosPoint)
		if err != nil {
			return szTongSpoint{}
		}
	} else {
		fmt.Println("Error when Unmarshal infosPoint!")
		return szTongSpoint{}
	}

	return infosPoint
}

func getIpList() []string {

	var iplist []string

	//To be done
	iplist = append(iplist, "115.239.211.112", "221.179.178.112", "123.58.180.8", "140.205.172.21", "133.130.97.172", "14.215.177.38", "54.192.27.6")

	return iplist
}

// SaveIpPonixToBlockchain SaveIpPonixToBlockchain
func SaveIpPonixToBlockchain(myNodeIpaddr string) map[string]CityInfo {
	iplist := []string{strings.TrimSpace(myNodeIpaddr)}
	ipPointxyMap := multiIp2pointxyint(iplist)

	//To be saved to db
	Repository := NewMongoRepository(config.MONGO_DB_URL)

	for mapKey, mapValue := range ipPointxyMap {
		Repository.CallInsertPointxy(mapKey, mapValue.Pointx, mapValue.Pointy)
	}
	
	return ipPointxyMap
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
		os.Exit(1)
	}
}

// NewMongoRepository NewMongoRepository
func NewMongoRepository(endpoint string) *MongoRepository {
	return &MongoRepository{mgoEndpoint: endpoint}
}

// GetSession GetSession
func GetSession(url string) (*MongoContext, error) {
	if url == "" {
		return nil, errors.New("invalid para url")
	}
	var err error
	mgoSession, err := mgo.Dial(url)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Dial faild" + url)
	}
	return &MongoContext{mgoSession.Clone()}, nil
}

// GetCollection GetCollection
func (c *MongoContext) GetCollection(db string, collection string) *mgo.Collection {
	session := c.mgoSession
	collects := session.DB(config.DB_BOTTOS).C(collection)
	return collects
}

// SetCollectionByDB SetCollectionByDB
func (c *MongoContext) SetCollectionByDB(db string, collection string, s func(*mgo.Collection) error) error {
	session := c.mgoSession
	defer session.Close()
	collects := session.DB(db).C(collection)
	//fmt.Println("SetCollectionByDB: collects:", collects)
	return s(collects)
}

// CallInsertPointxy CallInsertPointxy
func (r *MongoRepository) CallInsertPointxy(ip string, pointx string, pointy string) (uint32, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("Get session faild" + r.mgoEndpoint)
	}

	var mesgs *Ippointxy
	session.GetCollection(config.DB_BOTTOS, config.TABLE_POINTXY).Find(bson.M{"$or": []bson.M{{"ip": ip}}}).One(&mesgs)

	if mesgs == nil {
		record := &Ippointxy{
			ID:     bson.NewObjectId(),
			Ip:     ip,
			Pointx: pointx,
			Pointy: pointy}

		insert := func(c *mgo.Collection) error {
			return c.Insert(record)
		}

		err = session.SetCollectionByDB(config.DB_BOTTOS, config.TABLE_POINTXY, insert)
		fmt.Println(err)
		return 1, err

	}
	selector := bson.M{"ip": ip}
	data := bson.M{"$set": bson.M{"pointx": pointx, "pointy": pointy}}

	/*changeInfo*/
	changeInfo, err := session.mgoSession.DB(config.DB_BOTTOS).C(config.TABLE_POINTXY).UpdateAll(selector, data)
	if err != nil {
		fmt.Println("update failed!", changeInfo)
	}
	//fmt.Printf("%+v\n", changeInfo)
	return 1, err

}

// InsertRecord InsertRecord
func (r *MongoRepository) InsertRecord(tablename string, keyname string, key string, valueName string, value interface{}) (uint32, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("Get session faild" + r.mgoEndpoint)
	}

	var mesgs *Ippointxy
	session.GetCollection(config.DB_BOTTOS, tablename).Find(bson.M{"$or": []bson.M{{keyname: key}}}).One(&mesgs)
	fmt.Println("InsertRecord: tablename: ", tablename, ", mesgs:", mesgs)
	if mesgs == nil {

		insert := func(c *mgo.Collection) error {
			return c.Insert(value)
		}

		err = session.SetCollectionByDB(config.DB_BOTTOS, tablename, insert)
		fmt.Println(err)
		return 1, err

	}
	selector := bson.M{keyname: key}
	data := bson.M{"$set": bson.M{valueName: value}}

	/*changeInfo*/
	changeInfo, err := session.mgoSession.DB(config.DB_BOTTOS).C(config.TABLE_POINTXY).UpdateAll(selector, data)
	if err != nil {
		fmt.Println("update failed!", changeInfo)
	}
	//fmt.Printf("%+v\n", changeInfo)
	return 1, err

}
