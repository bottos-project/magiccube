#!/usr/bin/env bash

#GO_PATH=`echo $GOPATH |cut -d'=' -f2`
GOPATH=/mnt/bottos
SYS_GOPATH=/mnt
SYS_GOROOT=/usr/lib/go

GO_PATH=/opt/go
CONSUL_PATH=$GO_PATH/bin/linux_386/
SERVER_PATH=$GO_PATH/bin/
GZ_PACKAGE_DIR=/mnt

MINIO_SITE=https://dl.minio.io/server/minio/release/linux-amd64/minio
MINIO_SERV_SITE=https://raw.githubusercontent.com/minio/minio-service/master/linux-systemd/minio.service
MINIO_RUN_PATH=/usr/local/bin
MINIO_USER=minio-user
MINIO_GRP=minio-user
MINIO_SHR=/usr/local/share/minio
MINIO_COF=/etc/minio

SERVER_IPADR=
SERVER_PORT=9000
WALLET_SERV=
CHAIN_PORT=8888
EXTERNAL_IPADR=

MINIO_VOLUMES_ENV="MINIO_VOLUMES=\""$MINIO_SHR"\""
MINIO_OPTS_ENV="MINIO_OPTS=\"-C "$MINIO_COF" --address "$SERVER_IPADR":"$SERVER_PORT"\""

MINIO_VOLUMES="/usr/local/share/minio"
MINIO_OPTS="-C /etc/minio --address "$SERVER_IPADR":"$SERVER_PORT
CORE_PROC_FILE_DIR=${GO_PATH}/bin/core/cmd_dir

OPT_GO_BIN_GZ_PACK=opt-go-bin.tar.gz
GOPATH_DIR_PACK=GOPATH-DIR.tar.gz
GOLANG_PACK=go1.10.1.linux-amd64.tar.gz

#DO NOT FILL THESE IN SCRIPT
PACKAGE_SVR=""
PACKAGE_SVR_USRNAME=""
PACKAGE_SVR_PWD=""
PACKAGE_SVR_PACKAGE_DIR=""

function check_gzpackages() {
    sudo apt-get install -y tcl tk expect   
    counts=0
    while ( [ -z "$PACKAGE_SVR" ] || [ -z "$PACKAGE_SVR_USRNAME" ] || [ -z "$PACKAGE_SVR_PWD" ] || [ -z "$PACKAGE_SVR_PACKAGE_DIR" ]); do
        if [ $counts -gt 0 ]; then
            echo -e "\033[31m Input wrong. Please input again.\033[0m"
        fi
        read -p "Please input source package server ip:" PACKAGE_SVR
        read -p "Please input source package server packages directory:" PACKAGE_SVR_PACKAGE_DIR
        read -p "Please input source package server user name:" PACKAGE_SVR_USRNAME
        read -p "Please input source package server password:" PACKAGE_SVR_PWD
        counts=1
    done
    
    #if [ ! -d /opt/go/bin/core ]; then

/usr/bin/expect <<-EOF
sudo rm -rf /opt/go/bin/core 2>/dev/null
spawn sudo scp -r $PACKAGE_SVR_USRNAME@$PACKAGE_SVR:$PACKAGE_SVR_PACKAGE_DIR/core /opt/go/bin
set timeout 600

expect {
"*yes/no" { send "yes\r"; exp_continue }
"*password:" { send "$PACKAGE_SVR_PWD\r" }
}

expect eof

EOF

    #fi

    if [ ! -f $GZ_PACKAGE_DIR/$GOLANG_PACK ]; then

/usr/bin/expect <<-EOF

spawn scp $PACKAGE_SVR_USRNAME@$PACKAGE_SVR:$PACKAGE_SVR_PACKAGE_DIR/$GOLANG_PACK $GZ_PACKAGE_DIR
set timeout 600

expect {
"*yes/no" { send "yes\r"; exp_continue }
"*password:" { send "$PACKAGE_SVR_PWD\r" }
}

expect eof

EOF

    fi


    if [ ! -f $GZ_PACKAGE_DIR/$OPT_GO_BIN_GZ_PACK ]; then

/usr/bin/expect <<-EOF

spawn scp $PACKAGE_SVR_USRNAME@$PACKAGE_SVR:$PACKAGE_SVR_PACKAGE_DIR/$OPT_GO_BIN_GZ_PACK $GZ_PACKAGE_DIR
set timeout 600

expect {
"*yes/no" { send "yes\r"; exp_continue }
"*password:" { send "$PACKAGE_SVR_PWD\r" }
}

expect eof

EOF

    fi


    if [ ! -f $GZ_PACKAGE_DIR/$GOPATH_DIR_PACK ]; then

/usr/bin/expect <<-EOF

spawn scp $PACKAGE_SVR_USRNAME@$PACKAGE_SVR:$PACKAGE_SVR_PACKAGE_DIR/$GOPATH_DIR_PACK $GZ_PACKAGE_DIR
set timeout 600

expect {
"*yes/no" { send "yes\r"; exp_continue }
"*password:" { send "$PACKAGE_SVR_PWD\r" }
}

expect eof

EOF

    fi

    if [ ! -f $GZ_PACKAGE_DIR/$OPT_GO_BIN_GZ_PACK ] ; then
        echo -e "\033[31m *ERROR* Please get your missing gz packages [ $OPT_GO_BIN_GZ_PACK ] under directory $GZ_PACKAGE_DIR !!! \033[0m"
        exit 1
    fi
    
    if [ ! -f $GZ_PACKAGE_DIR/$GOPATH_DIR_PACK ]; then
        echo -e "\033[31m *ERROR* Please get your missing gz packages [ $GOPATH_DIR_PACK ] under directory $GZ_PACKAGE_DIR !!! \033[0m"
        exit 1
    fi
   
    if [ ! -f $GZ_PACKAGE_DIR/$GOLANG_PACK ]; then
        echo -e "\033[31m *ERROR* Please get your missing gz packages [ $GOLANG_PACK ] under directory $GZ_PACKAGE_DIR !!! \033[0m"
        exit 1
    fi
    
    echo "!!All files done."    
}

function unpackpackages() {
	echo "start unpack $GZ_PACKAGE_DIR/$OPT_GO_BIN_GZ_PACK:"
	mkdir -p /opt/go/bin
	tar zxvf $GZ_PACKAGE_DIR/$OPT_GO_BIN_GZ_PACK -C /opt/go
	
	echo "start unpack $GZ_PACKAGE_DIR/$GOPATH_DIR_PACK:"
	tar zxvf $GZ_PACKAGE_DIR/$GOPATH_DIR_PACK -C $SYS_GOPATH
	
    echo "start unpack $GZ_PACKAGE_DIR/$GOLANG_PACK:"
	tar -C /usr/lib -xzf $GZ_PACKAGE_DIR/$GOLANG_PACK
	
    cmd=$(sed  -n "/GOPATH/p" /etc/profile |wc -l)
    if [ $cmd -lt 1 ]; then
       sed -i '$a\export GOPATH="/mnt/bottos"' /etc/profile
    else 
       cmd="/GOROOT/c\GOROOT=\"/usr/lib/go\""
       sed -ir $cmd /etc/profile 
    fi

    cmd=$(sed  -n "/GOROOT/p" /etc/profile |wc -l)
    if [ $cmd -lt 1 ]; then
       sed -i "\$a\export GOROOT=\"$SYS_GOROOT\"" /etc/profile
    else
        cmd="/GOPATH/c\GOPATH=\"/mnt/bottos\""
        sed -ir $cmd /etc/profile
    fi
    
    cmd=$(sed  -n "/GOPATH/p" ~/.profile |wc -l)
    if [ $cmd -lt 1 ]; then
       sed -i '$a\export GOPATH="/mnt/bottos"' ~/.profile
    fi
   
    cmd=$(sed  -n "/GOROOT/p" ~/.profile |wc -l)
    if [ $cmd -lt 1 ]; then
        #cmd=\'+ '$a\export GOROOT='+$SYS_GOROOT+\'
        #echo "LYP--->\$a\export GOROOT=$SYS_GOROOT"
        sed -i "\$a\export GOROOT=\"$SYS_GOROOT\"" ~/.profile
    fi

    cmd=$(sed -n "/export PATH/p" ~/.profile |wc -l)
    if [ $cmd -lt 1 ]; then
        sed -i '$a\export PATH=$PATH:/usr/local/go/bin' ~/.profile
    fi
    
    sudo cp -rf /usr/bin/go /usr/lib

    source /etc/profile
    source ~/.profile
}

function miniocheck(){
    if [ -z "$SERVER_IPADR" ] || [ -z "$SERVER_PORT" ];
    then
		echo -e "\033[31m *ERROR* You need specify variable SERVER_IPADR and SERVER_PORT first !!! \033[0m"
        exit 1
    fi
    return
}

function usercheck(){
	SCRIPT_USER=`whoami`
	if [ $SCRIPT_USER != "root" ];
	then
		echo -e "\033[31m *ERROR* Please run the script as root !!! \033[0m"
		exit 1
	fi
	return
}

function setminio()
{
	apt-get update
	curl -O $MINIO_SITE
	if [ ! -e "./minio" ];
	then
		echo -e "\033[31m *ERROR* Fail to fetch the minio from the site: "$MINIO_SITE" \033[0m"
		exit 1
	fi

	mv minio /usr/local/bin/

	chmod a+x /usr/local/bin/minio

	useradd -r $MINIO_USER -s /sbin/nologin
	if [ -z `egrep "^"$MINIO_USER /etc/passwd 2> /dev/null` ];
	then
		echo -e "\033[31m *ERROR* Failed to create user for minio !!! \033[0m"
		exit 1
	fi

	chown $MINIO_USER:$MINIO_GRP /usr/local/bin/minio -R

	mkdir $MINIO_SHR

	chown $MINIO_USER:$MINIO_GRP $MINIO_SHR -R

	mkdir $MINIO_COF

	chown $MINIO_USER:$MINIO_GRP $MINIO_COF -R

	if [ -f /etc/default/minio ];
	then
		rm -rf /etc/default/minio
	fi

	echo $MINIO_VOLUMES_ENV >> /etc/default/minio
	echo $MINIO_OPTS_ENV >> /etc/default/minio
	source /etc/default/minio

	curl -O $MINIO_SERV_SITE
	if [ ! -f "./minio.service" ];
	then
		echo -e "\033[31m *ERROR* Fail to fetch the minio.service from the site: "${MINIO_SERV_SITE}" \033[0m"
		exit 1
	fi

	cp minio.service /etc/systemd/system

	systemctl daemon-reload
	systemctl enable minio

	if [ ! -z "`ufw status | grep inactive`" ];
	then
		ufw allow $SERVER_PORT
		systemctl restart ufw
	fi
}

function ssldepends()
{
    apt-get install pkg-config libssl-dev libsasl2-dev -y
    if [ ! -f ${GO_PATH}/bin/mongo-c-driver-1.9.2.tar.gz ] || [ ! -f ${GO_PATH}/bin/mongo-cxx-driver-r3.2.0-rc1.tar.gz ];
    then
        echo "*ERROR* Fail to find the tar file: mongo-c-driver-1.9.2.tar.gz or mongo-cxx-driver-r3.2.0-rc1.tar.gz from the path "${GO_PATH}" !!!"
        exit 1
    fi
    
    tar xzf mongo-c-driver-1.9.2.tar.gz
    cd mongo-c-driver-1.9.2
    ./configure --disable-automatic-init-and-cleanup --enable-static
    make
    make install
    cd ..

    tar xzf mongo-cxx-driver-r3.2.0-rc1.tar.gz
    cd mongo-cxx-driver-r3.2.0-rc1/build
    cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=/usr/local ..
    make EP_mnmlstc_core
    make
    make install
    cd ..

    rm -rf mongo-c-driver-1.9.2
    rm -rf mongo-cxx-driver-r3.2.0-rc1
}

function sslcheck()
{
	if [ -z "`dpkg -l | grep libsasl2-dev`" ] || [ -z "`dpkg -l | grep libssl-dev`" ];
	then
		echo -e "\033[33m *WRAN* set ssl runtime environment , it may take a moment ... \033[0m"
	    ssldepends
	fi
}

function varcheck()
{
	if [ -z "$GO_PATH" ] || [ -z "$SERVER_IPADR" ] || [ -z "$SERVER_PORT" ] || [ -z "$GZ_PACKAGE_DIR" ] || [ -z "$SYS_GOPATH" ] || [ -z "$SYS_GOROOT" ];
	then
		echo -e "\033[31m *ERROR* You have to specify the variable SYS_GOPATH/SYS_GOROOT/GO_PATH/SERVER_IPADR/SERVER_PORT/GZ_PACKAGE_DIR in the script!!! \033[0m"
		exit 1
	fi
}

function startcore()
{
    if [ ! -e /opt/go/bin/core/core ];
	then
		echo -e "\033[31m *ERROR* Fail to find core process under : /opt/go/bin/core !!! \033[0m"
		exit 1
	fi
	
	if [ ! -f /opt/go/bin/core/chainconfig.json ] || [ ! -f /opt/go/bin/core/genesis.json ];
	then
		echo -e "\033[31m *ERROR* Fail to find core proess's configuration item : config.json or genesis.json under :"${CORE_PROC_FILE_DIR}" \033[0m"
		exit 1
	fi

    cp -f /opt/go/bin/core/chainconfig.json /opt/go/bin
	cp -f /opt/go/bin/core/genesis.json /opt/go/bin

    if [ ! -d ${CORE_PROC_FILE_DIR} ];
	then
		echo -e "\033[31m *ERROR* Directory does not exist:"${CORE_PROC_FILE_DIR}" \033[0m"
		exit 1
	fi
	
    sudo rm -r /opt/go/bin/datadir 2>/dev/null
	sudo rm -r ${CORE_PROC_FILE_DIR}/datadir 2>/dev/null
	sudo rm -r $GO_PATH/bin/datadir 2>/dev/null

	#GENESIS_JSON=`cat ${GO_PATH}/bin/data-dir/config.ini | grep -e ^genesis-json | awk '{print $3}'`
	#if [ ! -f "$GENESIS_JSON" ];
	#then
	#	echo -e "\033[31m *ERROR* json file :"$GENESIS_JSON" dosn't exist !!! please reconfigure config.ini \033[0m"
	#	exit 1
	#fi
	
	CHK_CORE=`ps -elf | grep "/opt/go/bin/core/core" | grep -v grep | wc -l`
	if [ "$CHK_CORE" -lt 1 ];
	then 
		#start Core process  , nohup "command" > myout.file 2>&1 &
		sudo nohup /opt/go/bin/core/core 2>&1 & 
        	#--http-server-address ${SERVER_IPADR}:${CHAIN_PORT} -m mongodb://126.0.0.1/bottos --resync > core.file 2>&1 &
        	sleep 3
	fi
    return
}

function stopcore()
{
    RUNNING_CORE_PID=`ps -elf | grep /opt/go/bin/core | grep -v grep | awk '{print $4}'`
    if [ -z "$RUNNING_CORE_PID" ];
    then
        echo -e "\033[33m *WRAN* core process hadn't been running ... \033[0m"
        return
    fi

    kill -SIGINT $RUNNING_CORE_PID
    ps -ef | grep ${SERVER_PATH}"core" | grep -v grep | cut -c 9-15 | xargs kill -s 9

    return
}

function restartcore()
{
    stopcore
    startcore
}

function startminio()
{
	systemctl start minio
	minio server $MINIO_OPTS $MINIO_VOLUMES > minio.file 2>&1 &
	if [ "`ps -elf | grep minio | grep -v grep | wc -l`" -lt 1 ];
	then
		echo -e "\033[31m *ERROR* Failed to start minio service \033[0m"
		exit 1
	fi

}

function prepcheck()
{
	echo -e "\033[32m ==================================== \033[0m"
	echo -e "\033[32m = Prepare to check the environment = \033[0m"
	echo -e "\033[32m ==================================== \033[0m"
    #check if golang had been installed , if not , install it
    if [ -n "`cat /etc/issue | grep -i Ubuntu`" ] && [ -z "`which go`" ];
	then
	    echo -e "\033[33m *WRAN* golang hadn't been installed , install it currently ... \033[0m"
		apt-get update
        #apt-get install golang-go -y
	fi

	#check if git had been installed , if not , install it
	if [ -n "`cat /etc/issue | grep -i Ubuntu`" ] && [ -z "`which git`" ];
	then
	    echo -e "\033[33m *WRAN* git hadn't been installed , install it currently ... \033[0m"
        apt-get install git -y
	fi
	
	#check if cmake had been installed , if not , install it
	if [ -n "`cat /etc/issue | grep -i Ubuntu`" ] && [ -z "`dpkg -l | grep cmake`" ];
	then
	    echo -e "\033[33m *WRAN* cmake hadn't been installed , install it currently ... \033[0m"
        apt-get install cmake -y
	fi
	apt-get upgrade cmake -y

	#check if mongodb had been installed , if not , install it
	if [ -n "`cat /etc/issue | grep -i Ubuntu`" ] && [ -z "`dpkg -l | grep mongodb`" ];
	then
		echo -e "\033[33m *WRAN* mongodb hadn't been installed , install it currently ... \033[0m"
		apt-get install mongodb -y
	fi

	apt-get install jq -y

	echo -e "\033[32m check the environment ok ... \033[0m"
	return
}

function startcontract()
{
	sudo ${CORE_PROC_FILE_DIR}/./cmd newaccount -name usermng -pubkey 7QBxKhpppiy7q4AcNYKRY2ofb3mR5RP8ssMAX65VEWjpAgaAnF &
	sudo ${CORE_PROC_FILE_DIR}/./cmd deploycode -contract usermng -wasm $CORE_PROC_FILE_DIR/contract/usermng.wasm &
	echo "===CONTRACT DONE==="
}

function startserv()
{
	echo -e "\033[32m ==================================== \033[0m"
	echo -e "\033[32m =       Start bottos server        = \033[0m"
	echo -e "\033[32m ==================================== \033[0m"
	if [ ! -d "$GO_PATH" ];
	then
		echo -e "\033[31m *ERROR* Failed to find variable GO_PATH !!! \033[0m"
	    exit 1
	fi

	echo -e "\033[32m check GO_PATH="$GO_PATH" ok ... \033[0m"
	if [ ! -d "$CONSUL_PATH" ] || [ ! -d "$SERVER_PATH" ];
	then
		echo -e "\033[31m *ERROR* directory $CONSUL_PATH or $SERVER_PATH doesn't exist !!! \033[0m"
		exit 1
	fi

	# start consul for go-micro
	nohup ${CONSUL_PATH}consul agent -dev > consul.log 2>&1 &
	sleep 1

	# start micro service
	nohup ${SERVER_PATH}micro api > micro.log 2>&1 &
	# start mongodb service
	#service mongodb start
	sleep 3
	echo `ps -ef|grep micro`
        echo "startcore"
        startcore
	
	echo "start minio"
	# setup minio
	if [ ! -e /usr/local/bin/minio ];
	then
	    miniocheck
		setminio
	fi

	startminio

	# start node service , other services will be started by node server
    	#nohup ${SERVER_PATH}node > node.file 2>&1
    	${SERVER_PATH}./node
	
	startcontract
}

function stopserv()
{
	echo "GOPATH=$GO_PATH"
	ps -ef | grep -w ${SERVER_PATH}"exchange" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	ps -ef | grep -w ${SERVER_PATH}"excApi" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	ps -ef | grep -w ${SERVER_PATH}"asset" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	ps -ef | grep -w ${SERVER_PATH}"assApi" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	ps -ef | grep -w ${SERVER_PATH}"requirement" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	ps -ef | grep -w ${SERVER_PATH}"reqApi" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	ps -ef | grep -w ${SERVER_PATH}"dashboard" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	ps -ef | grep -w ${SERVER_PATH}"dasApi" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	ps -ef | grep -w ${SERVER_PATH}"identity" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	ps -ef | grep -w ${SERVER_PATH}"ideApi" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	ps -ef | grep ${SERVER_PATH}"storage" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	ps -ef | grep ${SERVER_PATH}"core" | grep -v grep | cut -c 9-15 | xargs kill -s 9
    miniopid=$(pidof minio)
    kill -9 $miniopid 2>/dev/null
    datapid=$(pidof data)
    kill -9 $datapid 2>/dev/null
    datApipid=$(pidof datApi)
    kill -9 $datApipid 2>/dev/null

	sleep 1

	ps -ef | grep -w ${SERVER_PATH}"./node" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	sleep 1

	ps -ef | grep "mongodb" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	sleep 1

	systemctl stop minio

	ps -ef | grep ${SERVER_PATH}"micro api" | grep -v grep | cut -c 9-15 | xargs kill -s 9

	ps -ef | grep ${CONSUL_PATH}"consul agent -dev" | grep -v grep | cut -c 9-15 | xargs kill -s 9

	echo -e "\033[32m Stop service ok ... \033[0m"

	return
}

function download_git_newcode()
{ 
    echo "WARNING: THIS STEP WILL OVERWRITE YOUR CODES UNDER GOPATH. ARE YOU SURE? $1? (Y/N) (default: N) __"
    read dorm
    dorm=${dorm:=N}
    if [ $dorm = N ]; then
        echo "OK. BYE"
        exit 1
    fi

    echo "Please input your eth0's IP address:"
    read eth0_ip

    if [ ! -z $GOPATH ]; then
        sudo rm -rf $GOPATH/src/github.com/bottos-project/bottos 2>&1>/dev/null
        sudo rm -rf $GOPATH/src/github.com/bottos-project/core   2>&1>/dev/null
        sudo rm -rf $GOPATH/src/github.com/bottos-project/bottos/service/node/keystore/crypto-go 2>&1>/dev/null
        sudo git clone https://github.com/bottos-project/bottos.git $GOPATH/src/github.com/bottos-project/bottos
        
        path=`pwd`
        cd $GOPATH/src/github.com/bottos-project/bottos/service/node/keystore
        sudo git clone https://github.com/bottos-project/crypto-go.git
        cd $path
        sudo git clone https://github.com/bottos-project/core.git $GOPATH/src/github.com/bottos-project/core
        
        cmd="/WALLET_IP/c\WALLET_IP=\"$eth0_ip\""
        sudo sed -ir $cmd $GOPATH/src/github.com/bottos-project/bottos/service/node/config/config.go
        cmd="/ipAddr/c\\\"ipAddr\":\"$eth0_ip,\""
        sudo sed -ir $cmd /opt/go/bin/config.json
        cmd="/walletIP/c\\\"walletIP\":\"$eth0_ip,\""
        sudo sed -ir $cmd /opt/go/bin/config.json
        cmd="/bind_ip/c\bind_ip=$eth0_ip"
        sudo sed -ir $cmd /etc/mongodb.conf
        sudo chmod 777 $GOPATH/src/github.com/bottos-project/* -R
        echo "\n Cloning all is done. Please try ./startup.sh buildstart for auto-build then, or try ./startup.sh start for directly start."
    fi
}

function build_all_modules()
{
    export GOPATH=/mnt/bottos
    export GOROOT=/usr/lib/go
    
    /usr/lib/go/bin/./go build github.com/bottos-project/core
    /usr/lib/go/bin/./go build github.com/bottos-project/bottos/service/node
    /usr/lib/go/bin/./go build github.com/bottos-project/bottos/service/asset
    /usr/lib/go/bin/./go build github.com/bottos-project/bottos/service/storage
    /usr/lib/go/bin/./go build github.com/bottos-project/bottos/service/requirement
    /usr/lib/go/bin/./go build github.com/bottos-project/bottos/service/exchange
    /usr/lib/go/bin/./go build github.com/bottos-project/bottos/service/dashboard/dasApi
    /usr/lib/go/bin/./go build github.com/bottos-project/bottos/service/dashboard   
    /usr/lib/go/bin/./go build github.com/bottos-project/bottos/service/data
    /usr/lib/go/bin/./go build github.com/bottos-project/bottos/service/data/datApi

    cp -f core        /opt/go/bin/core
    cp -f node        /opt/go/bin
    cp -f asset       /opt/go/bin
    cp -f storage     /opt/go/bin
    cp -f requirement /opt/go/bin
    cp -f exchange    /opt/go/bin
    cp -f dasApi      /opt/go/bin
    cp -f dashboard   /opt/go/bin
    cp -f data        /opt/go/bin
    cp -f datApi      /opt/go/bin

    cp -f /opt/go/bin/log.xml  /opt/go/bin/config
    cp -f /opt/go/bin/log.xml  /opt/go/bin/config/log-req.xml

    cd /opt/go/bin
}

#main
case $1 in
    "update")
        download_git_newcode        
        ;;
    "buildstart")
        build_all_modules
        usercheck
        varcheck
        service mongodb start
        startserv
        ;;
	"start")
        usercheck
        varcheck
        service mongodb start
        startserv 
        ;;
    "stop")
        usercheck
        stopserv
        ;;
    "deploy")
        usercheck
        varcheck   
	    check_gzpackages
        prepcheck
        sslcheck
	    unpackpackages
        ;;
    "startcore")
        startcore
        ;;
    "stopcore")
        stopcore
        ;;
    "restartcore")
        restartcore
        ;;
    "help"|*)
        echo -e "\033[32m you have to input a parameter , Please run the script like ./startup.sh deploy|update|start|buildstart|stop|startcore|stopcore|restartcore !!! \033[0m"
        ;;
esac

exit 0
