package main

import (
    "bufio"
	"fmt"
	"github.com/code/bottos/service/node/config"
    "github.com/code/bottos/service/node/api"
    "github.com/code/bottos/service/node/keystore"
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
	"os"
    "os/exec"
	"bytes"
	"strings"
	"errors"
    "github.com/code/bottos/service/node/keystore/crypto-go/crypto/aes"
	"runtime"  
    "syscall"  
    slog "github.com/cihub/seelog"

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
	newProcess, err := os.StartProcess(shellPath, argv, attr)  //run script
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Process PID", newProcess.Pid)
	processState, err := newProcess.Wait() //wait for process done
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("processState PID:", processState.Pid())//get PID
	fmt.Println("ProcessExit:", processState.Exited())

	return nil
}

func CreateAccount(nodeinfo api.NodeInfos, UserPwd string) error {
    
    keystore.AccountCreate_Ex(nodeinfo.Node[0].BtoPath, nodeinfo.Node[0].KeyPath, UserPwd)
    
    return nil
}

func CreateAccountOld(nodeinfo api.NodeInfos) error {
	//username := nodeinfo.Node[0].BtoUser
	//var out bytes.Buffer
	//url      := nodeinfo.Node[0].IpAddr + ":" + nodeinfo.Node[0].BtoPort
	command := config.SCRIPT_PATH+" createuser "+"\""+nodeinfo.Node[0].IpAddr+"\" "+"\""+nodeinfo.Node[0].WalletIP+"\" "+"\""+nodeinfo.Node[0].BtoPort+"\" "+"\""+nodeinfo.Node[0].BtoPath+"\" "+"\""+nodeinfo.Node[0].ProdUser+"\" "+"\""+nodeinfo.Node[0].BtoUser+"\""
	fmt.Println("command = ",command)
	if "windows" == config.RUN_PLATFORM {
		if buf,err := api.SshCommand(nodeinfo.Node[0].UserName ,
			nodeinfo.Node[0].PassWord ,
			nodeinfo.Node[0].IpAddr ,
			22 ,
			command); err == nil {
			//fmt.Println(buf)
			if strings.Contains(buf , "error") {
				errinfo := "*ERROR* Fail to create a account !!!"
				fmt.Println(errinfo)
				err = errors.New(errinfo)
				return err
			} else {
				fmt.Println("Create account "+nodeinfo.Node[0].BtoUser+" ok ...")
			}
		}
	} else if "linux" == config.RUN_PLATFORM {
		go exec_shell(command)
		fmt.Println("Create account "+nodeinfo.Node[0].BtoUser+" ok ...")
	}

	return nil
}

func GenerateKeyStone(nodeinfo api.NodeInfos) error {
	//var out bytes.Buffer
	command := config.SCRIPT_PATH+" createkey "+"\""+nodeinfo.Node[0].BtoPath+"\" "+"\""+nodeinfo.Node[0].KeyPath+"\" "
	fmt.Println("command = ",command)
	if "windows" == config.RUN_PLATFORM {
		if _,err := api.SshCommand(nodeinfo.Node[0].UserName ,
			nodeinfo.Node[0].PassWord ,
			nodeinfo.Node[0].IpAddr ,
			22 ,
			command); err != nil {
			errinfo := "*ERROR* Failed to generate the keystone file : "+nodeinfo.Node[0].KeyPath
			fmt.Println(errinfo)
			err = errors.New(errinfo)
		}
	} else if "linux" == config.RUN_PLATFORM {
		go exec_shell(command)
		//fmt.Println("buf = ",buf)
		fmt.Println("Generate keyStone file "+nodeinfo.Node[0].KeyPath+" ok ...")
	}
	return nil
}

func CheckKeyStore(nodeinfo api.NodeInfos, UserPwd string) error {
	username := nodeinfo.Node[0].BtoUser
	url      := nodeinfo.Node[0].IpAddr + ":" + nodeinfo.Node[0].BtoPort

	if accountinfo , err := api.GetAccountInfo(url , username); err != nil || accountinfo.AccountName != username {
		//user doesn't exist
		fmt.Println("*WARN* Account doesn't exist , create it ...")
		if err = CreateAccount(nodeinfo, UserPwd); err != nil {return err}
		//if err = GenerateKeyStore(nodeinfo);err != nil {return err}
	}else{
		//check if keystone file exists, it need add public key check in the feture
		fmt.Println("Check account exists ok ...")
		command := config.SCRIPT_PATH+" chkfile "+"\""+nodeinfo.Node[0].KeyPath+"\""
		//fmt.Println("command = ",command)
		if "windows" == config.RUN_PLATFORM {
			if buf,err := api.SshCommand(nodeinfo.Node[0].UserName ,
				nodeinfo.Node[0].PassWord ,
				nodeinfo.Node[0].IpAddr ,
				22 ,
				command); err != nil || len(buf) != 0 {
				errinfo := "*ERROR* Failed to search the keystone file : "+nodeinfo.Node[0].KeyPath
				fmt.Println(errinfo)
				err = errors.New(errinfo)
				return err
			} else {
				fmt.Println("Check keystone file ok ...")
			}

		}else if "linux" == config.RUN_PLATFORM {
            filename := nodeinfo.Node[0].UserName + ".keystore"
            filepath := nodeinfo.Node[0].KeyPath + "/" + filename 
			
            if api.PathExist(filepath) == false {
				errinfo := "*ERROR* Failed to search the keystone file : "+filepath
				log.Println(errinfo)
				err = errors.New(errinfo)
				return err
			} else {
				log.Println("Check keystone file ok ...Now decrypt file:", filepath)
                key, Account := aes.KeyDecrypt(filepath, UserPwd)
                log.Println("DECRYPT KEYSTORE DONE! Account:",Account,", key:", key)
			}
		}
	}

	return nil
}

func InitServer(nodeinfo api.NodeInfos) error {
	// read configurated servers from config file and set it

	for i := 0;i < len(nodeinfo.Node[0].ServLst);i++ {
		command := config.SCRIPT_PATH+" setserv "+"\""+nodeinfo.Node[0].ServPath+"\" "+"\""+nodeinfo.Node[0].ServLst[i]+"\""
		fmt.Println("command = ",command)
		if "windows" == config.RUN_PLATFORM {
			//command := config.SCRIPT_PATH+" setserv "+"\""+nodeinfo.Node[0].ServPath+"\" "+"\""+nodeinfo.Node[0].ServLst[i]+"\""
			//fmt.Println("command = ",command)
			if buf,err := api.SshCommand(nodeinfo.Node[0].UserName ,
				nodeinfo.Node[0].PassWord ,
				nodeinfo.Node[0].IpAddr ,
				22 ,
				command); err == nil {
				fmt.Println("buf = ",buf)
			}
		}else if "linux" == config.RUN_PLATFORM {
			//wg.Add(1)
			go exec_shell(command)

		}
	}
	//wg.Wait()

	fmt.Println("Init server ok ...")
	return nil
}

func InitDatabase(nodeinfo api.NodeInfos) error {
	if "windows" == config.RUN_PLATFORM {
		//
	}else if "linux" == config.RUN_PLATFORM {
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

func main() {
    var inputReader *bufio.Reader
    var input, input1, input2 string
    var err error

    inputReader = bufio.NewReader(os.Stdin)
    fmt.Println("Please your password: ")
    input1, err = inputReader.ReadString('\n')
    
    fmt.Println("Please input your password again: ")
    input2, err = inputReader.ReadString('\n')
    
    if input1 != input2 || err != nil || len(input1) <= 0 || input1 == "\n" {
        fmt.Println("Input error! Failed to start node.")
        return
    }

    input = input1
    
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

	if api.PathExist(config.CONFIG_FILE) == false {
		fmt.Println("*ERROR* configuration file :",config.CONFIG_FILE," doesn't exist !!!")
		return
	}

	//read configuration from json file
	var nodeinfos api.NodeInfos
	nodeinfos = api.ReadFile(config.CONFIG_FILE)

	//check if exists keystone file
    if CheckKeyStore(nodeinfos, UserPwd) != nil {return}
	
    log.Println("now call Save_ip_ponix_to_blockchain")
	api.Save_ip_ponix_to_blockchain()

	//set server according to the json file
	if InitServer(nodeinfos) != nil {return}

	//init mangodb
	if InitDatabase(nodeinfos) != nil {return}
	
    fmt.Println("Starting Server ...")
    
    daemon(0, 1)  

	//go monitorConfigFile(config.Conf_File)
	//user_proto.RegisterUserHandler(service.Server(), new(User))
	
    if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
