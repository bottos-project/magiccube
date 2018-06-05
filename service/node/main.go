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
package main

import (
	"time"
	//"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	pack "github.com/bottos-project/bottos/contract/msgpack"
	"github.com/bottos-project/magiccube/service/common/bean"
	"github.com/bottos-project/magiccube/service/common/data"
	push_sign "github.com/bottos-project/magiccube/service/common/signature/push"
	"github.com/bottos-project/magiccube/service/common/util"
	commonutil "github.com/bottos-project/magiccube/service/common/util"
	datautil "github.com/bottos-project/magiccube/service/data/util"
	"github.com/bottos-project/magiccube/service/node/api"
	"github.com/bottos-project/magiccube/service/node/config"
	"github.com/bottos-project/magiccube/service/node/keystore"
	aes "github.com/bottos-project/magiccube/service/node/keystore/crypto-go/crypto/aes"
	"github.com/bottos-project/magiccube/service/node/keystore/crypto-go/crypto/secp256k1"
	node_proto "github.com/bottos-project/magiccube/service/node/proto"
	slog "github.com/cihub/seelog"
	"github.com/howeyc/gopass"
	"github.com/micro/go-micro"
	"github.com/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	//"log"
	//"sync"
	//"reflect"
	//"github.com/micro/cli"
	//"github.com/urfave/cli"
	//"github.com/hashicorp/consul/version"
	//"golang.org/x/crypto/ssh"
	//"encoding/json"
	//"io/ioutil"
	//"time"
	//"net"
	//"unsafe"
	//"github.com/fsnotify/fsnotify"
	//	"github.com/code/bottos/service/storage/blockchain"
	//"github.com/code/bottos/service/storage/internal/platform/config"
	//"github.com/code/bottos/service/storage/internal/platform/minio"
	//"github.com/code/bottos/service/storage/internal/platform/sqlite"
	//"github.com/code/bottos/service/storacleage/internal/service"
)

//var wg sync.WaitGroup
func exec_shell(s string) {
	//defer wg.Done()
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}

func exec_process(shellPath string) error {
	//shellPath := "/home/xx/test.sh"
	argv := make([]string, 1)
	attr := new(os.ProcAttr)
	newProcess, err := os.StartProcess(shellPath, argv, attr) //run script
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Process PID", newProcess.Pid)
	processState, err := newProcess.Wait() //wait for process done
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("processState PID:", processState.Pid()) //get PID
	fmt.Println("ProcessExit:", processState.Exited())

	return nil
}

func CreateAccount(nodeinfo api.NodeInfos, UserPwd string) error {

	keystore.AccountCreate_Ex("/home/bto", "/home/bto", UserPwd)

	return nil
}

func CreateAccountOld(nodeinfo api.NodeInfos) error {
	//username := nodeinfo.Node[0].BtoUser
	//var out bytes.Buffer
	//url      := nodeinfo.Node[0].IpAddr + ":" + nodeinfo.Node[0].BtoPort
	command := config.SCRIPT_PATH + " createuser " + "\"" + nodeinfo.Node[0].IpAddr + "\" " + "\"" + nodeinfo.Node[0].WalletIP + "\" " + "\"" + nodeinfo.Node[0].BtoPort + "\" " + "\"" + nodeinfo.Node[0].BtoPath + "\" " + "\"" + nodeinfo.Node[0].ProdUser + "\" " + "\"" + nodeinfo.Node[0].BtoUser + "\""
	fmt.Println("command = ", command)
	if "windows" == config.RUN_PLATFORM {
		if buf, err := api.SshCommand(nodeinfo.Node[0].UserName,
			nodeinfo.Node[0].PassWord,
			nodeinfo.Node[0].IpAddr,
			22,
			command); err == nil {
			//fmt.Println(buf)
			if strings.Contains(buf, "error") {
				errinfo := "*ERROR* Fail to create a account !!!"
				fmt.Println(errinfo)
				err = errors.New(errinfo)
				return err
			} else {
				fmt.Println("Create account " + nodeinfo.Node[0].BtoUser + " ok ...")
			}
		}
	} else if "linux" == config.RUN_PLATFORM {
		go exec_shell(command)
		fmt.Println("Create account " + nodeinfo.Node[0].BtoUser + " ok ...")
	}

	return nil
}

func GenerateKeyStone(nodeinfo api.NodeInfos) error {
	//var out bytes.Buffer
	command := config.SCRIPT_PATH + " createkey " + "\"" + nodeinfo.Node[0].BtoPath + "\" " + "\"" + nodeinfo.Node[0].KeyPath + "\" "
	fmt.Println("command = ", command)
	if "windows" == config.RUN_PLATFORM {
		if _, err := api.SshCommand(nodeinfo.Node[0].UserName,
			nodeinfo.Node[0].PassWord,
			nodeinfo.Node[0].IpAddr,
			22,
			command); err != nil {
			errinfo := "*ERROR* Failed to generate the keystone file : " + nodeinfo.Node[0].KeyPath
			fmt.Println(errinfo)
			err = errors.New(errinfo)
		}
	} else if "linux" == config.RUN_PLATFORM {
		go exec_shell(command)
		//fmt.Println("buf = ",buf)
		fmt.Println("Generate keyStone file " + nodeinfo.Node[0].KeyPath + " ok ...")
	}
	return nil
}

func CheckKeyStore(nodeinfo api.NodeInfos, UserPwd string) error {
	username := nodeinfo.Node[0].BtoUser
	url := nodeinfo.Node[0].IpAddr + ":" + nodeinfo.Node[0].BtoPort
	if accountinfo, err := api.GetAccountInfo(url, username); err != nil || accountinfo.AccountName != username {
		//user doesn't exist
		fmt.Println("*WARN* Account doesn't exist , create it ...")
		if err = CreateAccount(nodeinfo, UserPwd); err != nil {
			return err
		}
		//if err = GenerateKeyStore(nodeinfo);err != nil {return err}
	} else {
		//check if keystone file exists, it need add public key check in the feture
		fmt.Println("Check account exists ok ...")
		command := config.SCRIPT_PATH + " chkfile " + "\"" + nodeinfo.Node[0].KeyPath + "\""
		//fmt.Println("command = ",command)
		if "windows" == config.RUN_PLATFORM {
			if buf, err := api.SshCommand(nodeinfo.Node[0].UserName,
				nodeinfo.Node[0].PassWord,
				nodeinfo.Node[0].IpAddr,
				22,
				command); err != nil || len(buf) != 0 {
				errinfo := "*ERROR* Failed to search the keystone file : " + nodeinfo.Node[0].KeyPath
				fmt.Println(errinfo)
				err = errors.New(errinfo)
				return err
			} else {
				fmt.Println("Check keystone file ok ...")
			}

		} else if "linux" == config.RUN_PLATFORM {
			filename := "bto.keystore"
			filepath := "/home/bto/" + filename

			if api.PathExist(filepath) == false {
				errinfo := "*ERROR* Failed to search the keystone file : " + filepath
				log.Println(errinfo)
				err = errors.New(errinfo)
				return err
			} else {
				log.Println("Check keystone file ok ...Now decrypt file:", filepath)
				key, Account := aes.KeyDecrypt(filepath, UserPwd)
				log.Println("DECRYPT KEYSTORE DONE! Account:", Account, ", key:", key)
			}
		}
	}

	return nil
}

func InitServer(nodeinfo api.NodeInfos) error {
	// read configurated servers from config file and set it

	for i := 0; i < len(nodeinfo.Node[0].ServLst); i++ {
		command := config.SCRIPT_PATH + " setserv " + "\"" + nodeinfo.Node[0].ServPath + "\" " + "\"" + nodeinfo.Node[0].ServLst[i] + "\""
		if "windows" == config.RUN_PLATFORM {
			//command := config.SCRIPT_PATH+" setserv "+"\""+nodeinfo.Node[0].ServPath+"\" "+"\""+nodeinfo.Node[0].ServLst[i]+"\""
			//fmt.Println("command = ",command)
			if buf, err := api.SshCommand(nodeinfo.Node[0].UserName,
				nodeinfo.Node[0].PassWord,
				nodeinfo.Node[0].IpAddr,
				22,
				command); err == nil {
				fmt.Println("buf = ", buf)
			}
		} else if "linux" == config.RUN_PLATFORM {
			//wg.Add(1)
			fmt.Println("Start service", nodeinfo.Node[0].ServLst[i], "...")
			go exec_shell(command)
			time.Sleep(1 * time.Second)
		}
	}
	//wg.Wait()

	fmt.Println("Init server ok ...")
	return nil
}

func InitDatabase(nodeinfo api.NodeInfos) error {
	if "windows" == config.RUN_PLATFORM {
		//
	} else if "linux" == config.RUN_PLATFORM {
		//
	}
	return nil
}

func daemon(nochdir, noclose int) int {
	var ret, ret2 uintptr
	var err syscall.Errno

	darwin := runtime.GOOS == "darwin"

	// already a daemon
	if syscall.Getppid() == 1 {
		return 0
	}

	// fork off the parent process
	ret, ret2, err = syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		return -1
	}

	// failure
	if ret2 < 0 {
		os.Exit(-1)
	}

	// handle exception for darwin
	if darwin && ret2 == 1 {
		ret = 0
	}

	// if we got a good PID, then we call exit the parent process.
	if ret > 0 {
		os.Exit(0)
	}

	/* Change the file mode mask */
	_ = syscall.Umask(0)

	// create a new SID for the child process
	s_ret, s_errno := syscall.Setsid()
	if s_errno != nil {
		log.Printf("Error: syscall.Setsid errno: %d", s_errno)
	}
	if s_ret < 0 {
		return -1
	}

	if nochdir == 0 {
		os.Chdir("/")
	}

	if noclose == 0 {
		f, e := os.OpenFile("/dev/null", os.O_RDWR, 0)
		if e == nil {
			fd := f.Fd()
			syscall.Dup2(int(fd), int(os.Stdin.Fd()))
			syscall.Dup2(int(fd), int(os.Stdout.Fd()))
			syscall.Dup2(int(fd), int(os.Stderr.Fd()))
		}
	}

	return 0
}

func slog_init() {

	defer slog.Flush()
	logger, err := slog.LoggerFromConfigAsFile("./log.xml")
	if err != nil {
		slog.Critical("err parsing config log file", err)
		os.Exit(1)
		return
	}
	slog.ReplaceLogger(logger)

	slog.Trace("Hello , World !!!")
	slog.Debug("Hello , World !!!")

}

func Sign(msg, seckey []byte) ([]byte, error) {
	sign, err := secp256k1.Sign(msg, seckey)
	return sign[:len(sign)-1], err
}

type nodeTrxInfo struct{}

func (u *nodeTrxInfo) Register(ctx context.Context, req *node_proto.RegisterRequests, rsp *node_proto.RegisterResponses) error {
	slog.Info("req:", req)
	block_header, err := data.BlockHeader()
	if err != nil {
		rsp.Code = 1003
		rsp.Msg = err.Error()
		return nil
	}
	//Register account
	rsp.Code = 1004
	req.NodeUUIDInfo.UserName = keystore.GetAccount()
	req.NodeUUIDInfo.PubKey = keystore.GetPubKey()

	account_buf, err := pack.Marshal(req.NodeUUIDInfo)
	if err != nil {
		rsp.Msg = err.Error()
		return nil
	}
	tx_account_sign := &push_sign.TransactionSign{
		Version:     1,
		CursorNum:   block_header.HeadBlockNum,
		CursorLabel: block_header.CursorLabel,
		Lifetime:    block_header.HeadBlockTime + 20,
		Sender:      req.NodeUUIDInfo.UserName,
		Contract:    "nodemng",
		Method:      "nodeinforeg",
		Param:       account_buf,
		SigAlg:      1,
	}

	msg, err := proto.Marshal(tx_account_sign)
	if err != nil {
		rsp.Msg = err.Error()
		return nil
	}

	key := keystore.GetPubKey()

	if len(key) <= 0 {
		key = "81407d25285450184d29247b5f06408a763f3057cba6db467ff999710aeecf8e"
	}

	seckey, err := hex.DecodeString(key)

	if err != nil {
		rsp.Msg = err.Error()
		return nil
	}

	signature, err := Sign(util.Sha256(msg), seckey)
	if err != nil {
		rsp.Msg = err.Error()
		return nil
	}

	tx_account := &bean.TxBean{
		Version:     1,
		CursorNum:   block_header.HeadBlockNum,
		CursorLabel: block_header.CursorLabel,
		Lifetime:    block_header.HeadBlockTime + 20,
		Sender:      req.NodeUUIDInfo.UserName,
		Contract:    "nodemng",
		Method:      "nodeinforeg",
		Param:       hex.EncodeToString(account_buf),
		SigAlg:      1,
		Signature:   hex.EncodeToString(signature),
	}

	ret, err := data.PushTransaction(tx_account)
	if err != nil {
		rsp.Code = 1005
		rsp.Msg = err.Error()
		return nil
	}

	slog.Info("ret-account:", ret.Result.TrxHash)

	return nil
}

func main() {
	var input string
	var input1 []byte
	var err error

	fmt.Println("\nPlease input your password for generating keystore: ")
	input1, err = gopass.GetPasswd()

	if err != nil || len(input1) <= 0 {
		fmt.Println("Input error! Failed to start node.")
		return
	}

	input = string(input1)

	if input == "\n" || len(input) <= 0 {
		fmt.Println("Input error! Failed to start node.")
		return
	}

	UserPwd := input

	slog_init()

	service := micro.NewService(
		micro.Name("go.micro.srv.node"),
		micro.Version("2.0.0"),
	)

	service.Init(
	/*
		micro.Action(func(c *cli.Context) {
			env := c.StringFlag("environment")
			if len(env) > 0 {
				fmt.Println("Environment set to", env)
			}
		}),*/
	)

	//node_proto.RegisterUserHandler(service.Server(), new(nodeTrxInfo))

	if api.PathExist(config.CONFIG_FILE) == false {
		fmt.Println("*ERROR* configuration file :", config.CONFIG_FILE, " doesn't exist !!!")
		return
	}

	//read configuration from json file
	var nodeinfos api.NodeInfos
	nodeinfos = api.ReadFile(config.CONFIG_FILE)

	//check if exists keystone file
	if CheckKeyStore(nodeinfos, UserPwd) != nil {
		return
	}

	//log.Println("now call Save_ip_ponix_to_blockchain")
	api.Save_ip_ponix_to_blockchain()

	//set server according to the json file
	if InitServer(nodeinfos) != nil {
		return
	}

	//init mangodb
	if InitDatabase(nodeinfos) != nil {
		return
	}

	SetNodeDBClusterInfo()

	fmt.Println("Starting Server ...")

	daemon(0, 1)

	//go monitorConfigFile(config.Conf_File)
	//user_proto.RegisterUserHandler(service.Server(), new(User))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func PushNodeClusterTrx(value interface{}, pri_key string) {

	block_header, err := data.BlockHeader()
	if err != nil {
		return
	}

	account_buf, err := pack.Marshal(value)
	if err != nil {
		return
	}
	tx_account_sign := &push_sign.TransactionSign{
		Version:     1,
		CursorNum:   block_header.HeadBlockNum,
		CursorLabel: block_header.CursorLabel,
		Lifetime:    block_header.HeadBlockTime + 20,
		Sender:      "bottos",
		Contract:    "nodemng",
		Method:      "nodeinforeg2",
		Param:       account_buf,
		SigAlg:      1,
	}

	msg, err := proto.Marshal(tx_account_sign)
	if err != nil {
		return
	}
	//配对的pubkey   0401787e34de40f3aeb4c28259637e8c9e84b5a58f57b3c23f010f4dc7230dffced4976238196bd32cd90569d66f747525b194ca83146965df092d2585b975d0d3
	seckey, err := hex.DecodeString(pri_key) //("81407d25285450184d29247b5f06408a763f3057cba6db467ff999710aeecf8e")
	if err != nil {
		return
	}

	signature, err := Sign(commonutil.Sha256(msg), seckey)
	if err != nil {
		return
	}

	tx_account := &bean.TxBean{
		Version:     1,
		CursorNum:   block_header.HeadBlockNum,
		CursorLabel: block_header.CursorLabel,
		Lifetime:    block_header.HeadBlockTime + 20,
		Sender:      "bottos",
		Contract:    "bottos",
		Method:      "newaccount",
		Param:       hex.EncodeToString(account_buf),
		SigAlg:      1,
		Signature:   hex.EncodeToString(signature),
	}

	ret, err := data.PushTransaction(tx_account)
	if err != nil {
		return
	}

	log.Info("ret-account:", ret.Result.TrxHash)
}

func SetNodeDBClusterInfo() {
	node_infos := api.ReadFile(config.CONFIG_FILE)

	var dbclusterinfo api.StorageDBClusterInfo

	node_uuid := keystore.GetUUID()
	dbclusterinfo.Nodetype = node_uuid
	dbclusterinfo.Nodedbinfo.NodeId = node_uuid
	dbclusterinfo.Nodedbinfo.NodeIP = node_infos.Node[0].IpAddr
	dbclusterinfo.Nodedbinfo.NodePort = node_infos.Node[0].BtoPort
	dbclusterinfo.Nodedbinfo.NodeAddress = node_infos.Node[0].IpAddr
	dbclusterinfo.Nodedbinfo.SeedIP = node_infos.Node[0].SeedIp
	dbclusterinfo.Nodedbinfo.SlaveIP = node_infos.Node[0].SlaveIpLst

	log.Println("Repository: ", config.MONGO_DB_URL)
	Repository := api.NewMongoRepository(config.MONGO_DB_URL)
	log.Println("Repository: ", Repository)

	PushNodeClusterTrx(dbclusterinfo, keystore.GetPriKey())
}

type NodeApi interface {
	GetNodeDBInfoList(nodedbinfo *datautil.NodeDBInfo)
}
