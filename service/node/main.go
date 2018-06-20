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
	aes "github.com/bottos-project/crypto-go/crypto/aes"
	"github.com/bottos-project/crypto-go/crypto/secp256k1"
	"github.com/bottos-project/gopass"
	"github.com/bottos-project/magiccube/service/common/bean"
	"github.com/bottos-project/magiccube/service/common/data"
	push_sign "github.com/bottos-project/magiccube/service/common/signature/push"
	"github.com/bottos-project/magiccube/service/common/util"
	commonutil "github.com/bottos-project/magiccube/service/common/util"
	datautil "github.com/bottos-project/magiccube/service/data/util"
	"github.com/bottos-project/magiccube/service/node/api"
	"github.com/bottos-project/magiccube/service/node/config"
	"github.com/bottos-project/magiccube/service/node/keystore"
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	"github.com/protobuf/proto"
	"math/big"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	global_config "github.com/bottos-project/magiccube/config"
)

//var wg sync.WaitGroup
func execShell(s string) {
	//defer wg.Done()
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}
	fmt.Printf("%s", out.String())
}

func execProcess(shellPath string) error {
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

//CreateAccount function
func CreateAccount(nodeinfo api.NodeInfos, UserPwd string) error {

	keystore.AccountCreateEx("/home/bto", "/home/bto", UserPwd)

	return nil
}

//CreateAccountOld function
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
			}

			fmt.Println("Create account " + nodeinfo.Node[0].BtoUser + " ok ...")

		}
	} else if "linux" == config.RUN_PLATFORM {
		go execShell(command)
		fmt.Println("Create account " + nodeinfo.Node[0].BtoUser + " ok ...")
	}

	return nil
}

//GenerateKeyStone function
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
		go execShell(command)
		//fmt.Println("buf = ",buf)
		fmt.Println("Generate keyStone file " + nodeinfo.Node[0].KeyPath + " ok ...")
	}
	return nil
}

//CheckKeyStore function
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
			}

			fmt.Println("Check keystone file ok ...")

		} else if "linux" == config.RUN_PLATFORM {
			filename := "bto.keystore"
			filepath := "/home/bto/" + filename

			if api.PathExist(filepath) == false {
				errinfo := "*ERROR* Failed to search the keystone file : " + filepath
				log.Info(errinfo)
				err = errors.New(errinfo)
				return err
			}

			log.Info("Check keystone file ok ...Now decrypt file:", filepath)
			key, Account := aes.KeyDecrypt(filepath, UserPwd)
			log.Info("DECRYPT KEYSTORE DONE! Account:", Account, ", key:", key)

		}
	}

	return nil
}

//InitServer function
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
			go execShell(command)
			time.Sleep(1 * time.Second)
		}
	}
	//wg.Wait()

	fmt.Println("Init server ok ...")
	return nil
}

//InitDatabase function
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
	sret, serrno := syscall.Setsid()
	if serrno != nil {
		log.Error("Error: syscall.Setsid errno: %d", serrno)
	}
	if sret < 0 {
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

func sloginit() {

	defer log.Flush()
	logger, err := log.LoggerFromConfigAsFile("./config/node-log.xml")
	if err != nil {
		log.Critical("err parsing config log file", err)
		os.Exit(1)
		return
	}
	log.ReplaceLogger(logger)
	//log.Trace("11 Hello , World !!!")
	//log.Debug("22 Hello , World !")
}

//Sign function
func Sign(msg, seckey []byte) ([]byte, error) {
	sign, err := secp256k1.Sign(msg, seckey)
	return sign[:len(sign)-1], err
}

//NodeUUIDInfo struct
type NodeUUIDInfo struct {
	UserName string
	PubKey   string
}

//Register function
func Register() error {
	//slog.Info("req:", req);
	blockheader, err := data.BlockHeader()
	if err != nil {
		return nil
	}

	var nodeUUIDinfo NodeUUIDInfo

	//Register account
	nodeUUIDinfo.UserName = keystore.GetAccount()
	nodeUUIDinfo.PubKey = keystore.GetPubKey()

	accountbuf, err := pack.Marshal(nodeUUIDinfo)
	if err != nil {
		return nil
	}
	
    txAccountSign := &push_sign.TransactionSign{
		Version:     1,
		CursorNum:   blockheader.HeadBlockNum,
		CursorLabel: blockheader.CursorLabel,
		Lifetime:    blockheader.HeadBlockTime + 20,
		Sender:      nodeUUIDinfo.UserName,
		Contract:    "nodemng",
		Method:      "nodeinforeg",
		Param:       accountbuf,
		SigAlg:      1,
	}

	msg, err := proto.Marshal(txAccountSign)
	if err != nil {
		return nil
	}

	key := keystore.GetPriKey()

	if len(key) <= 0 {
		return nil
	}

	seckey, err := hex.DecodeString(key)

	if err != nil {
		return nil
	}

    //Add chainID Flag
    chainID,_:=hex.DecodeString(global_config.CHAIN_ID)
    msg = bytes.Join([][]byte{msg, chainID}, []byte{})
               
	signature, err := Sign(util.Sha256(msg), seckey)
	if err != nil {
		return nil
	}

	txAccount := &bean.TxBean{
		Version:     1,
		CursorNum:   blockheader.HeadBlockNum,
		CursorLabel: blockheader.CursorLabel,
		Lifetime:    blockheader.HeadBlockTime + 20,
		Sender:      nodeUUIDinfo.UserName,
		Contract:    "nodemng",
		Method:      "nodeinforeg",
		Param:       hex.EncodeToString(accountbuf),
		SigAlg:      1,
		Signature:   hex.EncodeToString(signature),
	}

	ret, err := data.PushTransaction(txAccount)
	if err != nil {
		return nil
	}

	log.Info("ret-account:", ret.Result.TrxHash)

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

	sloginit()
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
	api.SaveIpPonixToBlockchain()

	//set server according to the json file
	if InitServer(nodeinfos) != nil {
		return
	}

	//init mangodb
	if InitDatabase(nodeinfos) != nil {
		return
	}
    
	SetNodeDBClusterInfo(nodeinfos)

	fmt.Println("Starting Server ...")

	daemon(0, 1)

	//go monitorConfigFile(config.Conf_File)
	//user_proto.RegisterUserHandler(service.Server(), new(User))

	if err := service.Run(); err != nil {
		log.Error(err)
	}
}

//NewNodeClusterAccount function
func NewNodeClusterAccount(nodeinfos api.NodeInfos, value interface{}, pubkey string, prikey string) {
    
    pubkey = "0454f1c2223d553aa6ee53ea1ccea8b7bf78b8ca99f3ff622a3bb3e62dedc712089033d6091d77296547bc071022ca2838c9e86dec29667cf740e5c9e654b6127f"
    prikey = "b799ef616830cd7b8599ae7958fbee56d4c8168ffd5421a16025a398b8a4be45"
	
    blockheader, err := data.BlockHeader()
	if err != nil {
		return
	}

	type NewAccountParam struct {
		Name   string
		Pubkey string
	}

	useraccount := &NewAccountParam {
		Name:   "nodeclustermng",
		Pubkey: pubkey,
	}

	accountbuf, err := pack.Marshal(useraccount)
	if err != nil {
		return
	}

	txAccountSign := &push_sign.TransactionSign{
		Version:     1,
		CursorNum:   blockheader.HeadBlockNum,
		CursorLabel: blockheader.CursorLabel,
		Lifetime:    blockheader.HeadBlockTime + 100,
		Sender:      "delta",
		Contract:    "bottos",
		Method:      "newaccount",
		Param:       accountbuf,
		SigAlg:      1,
	}
	
    msg, err2 := proto.Marshal(txAccountSign)
	if err2 != nil {
		return
	}
	
    seckey, err3 := hex.DecodeString(prikey)
	if err3 != nil {
		return
	}
    
    //Add chainID Flag
    chainID,_:=hex.DecodeString(global_config.CHAIN_ID)
    msg = bytes.Join([][]byte{msg, chainID}, []byte{})

	signature, err4 := Sign(commonutil.Sha256(msg), seckey)
	if err4 != nil {
		return
	}

	txAccount := &bean.TxBean{
		Version:     1,
		CursorNum:   blockheader.HeadBlockNum,
		CursorLabel: blockheader.CursorLabel,
		Lifetime:    blockheader.HeadBlockTime + 100,
		Sender:      "delta",//nodeinfos.Node[0].NodeName,
		Contract:    "bottos",//"nodeclustermng",
		Method:      "newaccount",//"reg",
		Param:       hex.EncodeToString(accountbuf),
		SigAlg:      1,
		Signature:   hex.EncodeToString(signature),
	}
    
	_, err = data.PushTransaction(txAccount)
	if err != nil {
		log.Error("NewNodeClusterAccount ERROR !")
		return
	}
}

//PushNodeClusterTrx function
func PushNodeClusterTrx(nodeinfos api.NodeInfos, value /*interface{}*/ api.StorageDBClusterInfo, pubkey string, prikey string) {
    
    pubkey = "0454f1c2223d553aa6ee53ea1ccea8b7bf78b8ca99f3ff622a3bb3e62dedc712089033d6091d77296547bc071022ca2838c9e86dec29667cf740e5c9e654b6127f"
    prikey = "b799ef616830cd7b8599ae7958fbee56d4c8168ffd5421a16025a398b8a4be45"
	
    blockheader, err := data.BlockHeader()
	if err != nil {
		return
	}

	accountbuf, err := pack.Marshal(&value)
    if err != nil {
        return
    }
    
	txAccountSign := &push_sign.TransactionSign{
		Version:     1,
		CursorNum:   blockheader.HeadBlockNum,
		CursorLabel: blockheader.CursorLabel,
		Lifetime:    blockheader.HeadBlockTime + 100,
		Sender:      "bottos",
		Contract:    "nodeclustermng",
		Method:      "reg",
		Param:       accountbuf,
		SigAlg:      1,
	}
	
    msg, err2 := proto.Marshal(txAccountSign)
	if err2 != nil {
		return
	}
	
    seckey, err3 := hex.DecodeString(prikey)
	if err3 != nil {
		return
	}
    
    //Add chainID Flag
    chainID,_:=hex.DecodeString(global_config.CHAIN_ID)
    msg = bytes.Join([][]byte{msg, chainID}, []byte{})

	signature, err4 := Sign(commonutil.Sha256(msg), seckey)
	if err4 != nil {
		return
	}
	
	txAccount := &bean.TxBean{
		Version:     1,
		CursorNum:   blockheader.HeadBlockNum,
		CursorLabel: blockheader.CursorLabel,
		Lifetime:    blockheader.HeadBlockTime + 100,
		Sender:      "bottos",
		Contract:    "nodeclustermng",
		Method:      "reg",
		Param:       hex.EncodeToString(accountbuf),
		SigAlg:      1,
		Signature:   hex.EncodeToString(signature),
	}
    
	_, err = data.PushTransaction(txAccount)
	if err != nil {
		log.Error("PushTransaction ERROR ! txAccount.Param: ", txAccount.Param)
		return
	}
}

//InetNtoA function
func InetNtoA(ip int64) string {

	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

//InetNtoByte function
func InetAtoByte(ip string) []byte {
    var iplst []byte = make([]byte, 4)

	fmt.Sscanf(ip, "%d.%d.%d.%d", &iplst[0], &iplst[1], &iplst[2], &iplst[3])
    
    return iplst
}

func InetAtoString(ip string) string {
    
    return fmt.Sprintf("%08x", InetAtoN(ip))
}

//InetAtoN function
func InetAtoN(ip string) int64 {
	ipvalue := net.ParseIP(ip).To4()
	if ipvalue == nil {
		return -1
	}
	ret := big.NewInt(0)
	ret.SetBytes(ipvalue)

	return ret.Int64()
}

//SetNodeDBClusterInfo function
func SetNodeDBClusterInfo(nodeinfos api.NodeInfos) {
	var dbclusterinfo api.StorageDBClusterInfo
    var iplst string
	dbclusterinfo.SeedIP = InetAtoString(nodeinfos.Node[0].SeedIp)
	for _, ip := range nodeinfos.Node[0].SlaveIpLst {
        iplst = InetAtoString(ip)

        dbclusterinfo.SlaveIP += iplst
    }
    pubkey := keystore.GetPubKey()
    prikey := keystore.GetPriKey()
    
    PushNodeClusterTrx(nodeinfos, dbclusterinfo, pubkey, prikey)
    
}

//NodeApi interface
type NodeApi interface {
	GetNodeDBInfoList(nodedbinfo *datautil.NodeDBInfo)
}
