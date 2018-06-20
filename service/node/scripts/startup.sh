#!/usr/bin/env bash

USER_HOME_DIR=/home/bottos
GOPATH=$USER_HOME_DIR/mnt/bottos
SYS_GOPATH=/mnt
SYS_GOROOT=/usr/lib/go

GO_PATH=$USER_HOME_DIR/opt/go
CONSUL_PATH=$GOPATH/src/
MICRO_PATH=$GOPATH/src/
SERVER_PATH=$GO_PATH/bin/
GZ_PACKAGE_DIR=/mnt

MINIO_SITE=https://dl.minio.io/server/minio/release/linux-amd64/minio
MINIO_SERV_SITE=https://raw.githubusercontent.com/minio/minio-service/master/linux-systemd/minio.service
MINIO_RUN_PATH=/usr/local/bin
MINIO_USER=bottos
MINIO_GRP=bottos
MINIO_SHR=/usr/local/share/minio
MINIO_COF=/etc/minio

OPT_GO_BIN=$USER_HOME_DIR/opt/go/bin

if [ -z $1 ]; then
    echo -e "\033[32m you have to input a parameter , Please run the script like ./startup.sh deploy|update|start|buildstart|stop|startcore|stopcore|restartcore !!! \033[0m"
    exit 1
fi

if [ $1 != "stop" ]; then 
    read -p "Please input your server ip address:" SERVER_IPADR
else
    SERVER_IPADR="127.0.0.1"    
fi

SERVER_PORT=9000
WALLET_SERV=$SERVER_IPADR
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

function miniocheck()
{
    if [ -z "$SERVER_IPADR" ] || [ -z "$SERVER_PORT" ];
    then
		echo -e "\033[31m *ERROR* You need specify variable SERVER_IPADR and SERVER_PORT first !!! \033[0m"
        exit 1
    fi
    return
}

function usercheck()
{
	SCRIPT_USER=`whoami`
	if [ -z $1 ] || [ $SCRIPT_USER != $1 ];
	then
		echo -e "\033[31m *ERROR* Please run the script as $1 !!! \033[0m"
		exit 1
	fi
	return
}

function setminio()
{
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
	chmod +x minio.service
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
    echo "HERE"	
    if [ -z "`dpkg -l | grep libsasl2-dev`" ] || [ -z "`dpkg -l | grep libssl-dev`" ];
	then
	echo -e "\033[33m *WRAN* set ssl runtime environment , it may take a moment ... \033[0m"
    	apt-get install pkg-config libssl-dev libsasl2-dev -y
    fi

    cd $USER_HOME_DIR
     
    wget https://github.com/mongodb/mongo-c-driver/releases/download/1.10.1/mongo-c-driver-1.10.1.tar.gz --directory-prefix=$USER_HOME_DIR
    if [ ! -f $USER_HOME_DIR/mongo-c-driver-1.10.1.tar.gz ];
    then
        echo "*ERROR* Fail to find the tar file: mongo-c-driver-1.9.2.tar.gz !"
         exit 1
    fi

    tar xzf $USER_HOME_DIR/mongo-c-driver-1.10.1.tar.gz
    
    cd $USER_HOME_DIR/mongo-c-driver-1.10.1
    mkdir cmake-build
    cd cmake-build
    cmake -DENABLE_AUTOMATIC_INIT_AND_CLEANUP=OFF ..
    make
    make install

    #./configure --disable-automatic-init-and-cleanup --enable-static
    #make
    #make install
    #cd ..
    
    wget https://github.com/mongodb/mongo-cxx-driver/archive/r3.2.0.tar.gz --directory-prefix=$USER_HOME_DIR
    if [ ! -f $USER_HOME_DIR/r3.2.0.tar.gz ];
    then
        echo "*ERROR* Fail to find the tar file: r3.2.0.tar.gz !"
        exit 1
    fi
    
    tar -xzvf $USER_HOME_DIR/r3.2.0.tar.gz
    cd $USER_HOME_DIR/mongo-cxx-driver-r3.2.0/build
    cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=/usr/local ..
    make EP_mnmlstc_core
    make
    make install
    cd ..

    rm -rf $USER_HOME_DIR/mongo-c*
    rm -rf $USER_HOME_DIR/r3.2.0*
}

function sslcheck()
{
	#if [ -z "`dpkg -l | grep libsasl2-dev`" ] || [ -z "`dpkg -l | grep libssl-dev`" ];
	#then
	#	echo -e "\033[33m *WRAN* set ssl runtime environment , it may take a moment ... \033[0m"
	    ssldepends
	#fi
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
    if [ ! -e $USER_HOME_DIR/opt/go/bin/core/bottos ];
	then
		echo -e "\033[31m *ERROR* Fail to find bottos process under : $USER_HOME_DIR/opt/go/bin/core !!! \033[0m"
		exit 1
	fi
	
	if [ ! -f $USER_HOME_DIR/opt/go/bin/core/chainconfig.json ] || [ ! -f $USER_HOME_DIR/opt/go/bin/core/genesis.json ];
	then
		echo -e "\033[31m *ERROR* Fail to find core proess's configuration item : config.json or genesis.json under :"${CORE_PROC_FILE_DIR}" \033[0m"
		exit 1
	fi

    cp -f $USER_HOME_DIR/opt/go/bin/*.json $USER_HOME_DIR
    cp -f $USER_HOME_DIR/opt/go/bin/*.json $USER_HOME_DIR/opt/go/bin/core
    cp -f $USER_HOME_DIR/opt/go/bin/*.json $USER_HOME_DIR/opt/go/bin/core/cmd_dir
    
    if [ ! -d ${CORE_PROC_FILE_DIR} ];
	then
		echo -e "\033[31m *ERROR* Directory does not exist:"${CORE_PROC_FILE_DIR}" \033[0m"
		exit 1
	fi

    CHK_CORE=`ps -elf | grep "$USER_HOME_DIR/opt/go/bin/core/bottos" | grep -v grep | wc -l`
    if [ "$CHK_CORE" -lt 1 ];
	then 
		#start Core process  , nohup "command" > myout.file 2>&1 &
		nohup $USER_HOME_DIR/opt/go/bin/core/bottos 2>&1 & 
        	#--http-server-address ${SERVER_IPADR}:${CHAIN_PORT} -m mongodb://126.0.0.1/bottos --resync > core.file 2>&1 &
        	sleep 3
	fi
    return
}

function stopcore()
{
    RUNNING_CORE_PID=`ps -elf | grep $USER_HOME_DIR/opt/go/bin/core/bottos | grep -v grep | awk '{print $4}'`
    if [ -z "$RUNNING_CORE_PID" ];
    then
        echo -e "\033[33m *WRAN* bottos process hadn't been running ... \033[0m"
        return
    fi

    kill -SIGINT $RUNNING_CORE_PID
    ps -ef | grep "$USER_HOME_DIR/opt/go/bin/core/bottos" | grep -v grep | cut -c 9-15 | xargs kill -s 9

    return
}

function restartcore()
{
    stopcore
    startcore
}

function startminio()
{
	#systemctl start minio
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
        rm -rf /var/lib/dpkg/lock 2>/dev/null
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
	sleep 1
	/usr/lib/go/bin/./go build github.com/bottos-project/bottos/bcli
        cp -rf bcli ${CORE_PROC_FILE_DIR} 2>/dev/null
	#${CORE_PROC_FILE_DIR}/./bcli newaccount -name usermng -pubkey 0454f1c2223d553aa6ee53ea1ccea8b7bf78b8ca99f3ff622a3bb3e62dedc712089033d6091d77296547bc071022ca2838c9e86dec29667cf740e5c9e654b6127f &
	#${CORE_PROC_FILE_DIR}/./bcli deploycode -contract usermng -wasm $CORE_PROC_FILE_DIR/contract/usermng.wasm &
	
    ${CORE_PROC_FILE_DIR}/./bcli newaccount -name nodeclustermng -pubkey 0454f1c2223d553aa6ee53ea1ccea8b7bf78b8ca99f3ff622a3bb3e62dedc712089033d6091d77296547bc071022ca2838c9e86dec29667cf740e5c9e654b6127f &
	${CORE_PROC_FILE_DIR}/./bcli deploycode -contract nodeclustermng -wasm $CORE_PROC_FILE_DIR/contract/nodeclustermng.wasm &
	sleep 1
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

    	cd $USER_HOME_DIR/opt/go/bin
	
	# start consul for go-micro
	nohup ${CONSUL_PATH}consul agent -dev > consul.log 2>&1 &
	sleep 1

	# start micro service
	nohup ${MICRO_PATH}micro api > micro.log 2>&1 &
	# start mongodb service
	sleep 3
    	
        startcore
	
	#echo "start minio"
	# setup minio
	if [ ! -e /usr/local/bin/minio ];
	then
	    miniocheck
		setminio
	fi

	startminio

	startcontract
	
    # start node service , other services will be started by node server
    	#nohup ${SERVER_PATH}node > node.file 2>&1
    	${SERVER_PATH}./node
	
        return
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
	ps -ef | grep -w ${SERVER_PATH}"user" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	ps -ef | grep -w ${SERVER_PATH}"useApi" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	ps -ef | grep ${SERVER_PATH}"storage" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	ps -ef | grep "$USER_HOME_DIR/opt/go/bin/core/bottos" | grep -v grep | cut -c 9-15 | xargs kill -s 2
	sleep 2
	ps -ef | grep "$USER_HOME_DIR/opt/go/bin/core/bottos" | grep -v grep | cut -c 9-15 | xargs kill -s 2
        
        miniopid=$(pidof minio)
        kill -9 $miniopid 2>/dev/null
        datapid=$(pidof data)
    	kill -9 $datapid 2>/dev/null
    	datApipid=$(pidof datApi)
    	kill -9 $datApipid 2>/dev/null
	sleep 1

	ps -ef | grep -w ${SERVER_PATH}"./node" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	sleep 1
       
        service mongodb stop
	#ps -ef | grep "mongodb" | grep -v grep | cut -c 9-15 | xargs kill -s 9
	sleep 1

	systemctl stop minio

	ps -ef | grep ${MICRO_PATH}"micro api" | grep -v grep | cut -c 9-15 | xargs kill -s 9

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

    eth0_ip=$SERVER_IPADR
    #if [ ! -z $GOPATH ]; then
        rm -rf $GOPATH/src/github.com/bottos-project/bottos    2>&1>/dev/null
        rm -rf $GOPATH/src/github.com/bottos-project/magiccube 2>&1>/dev/null
        rm -rf $GOPATH/src/github.com/bottos-project/crypto-go 2>&1>/dev/null

	cd $GOPATH/src/github.com/bottos-project/
	
        git clone https://github.com/bottos-project/bottos.git 
        git clone https://github.com/bottos-project/magiccube.git 
        git clone https://github.com/bottos-project/crypto-go.git
        git clone https://github.com/howeyc/gopass.git
	git clone https://github.com/bottos-project/msgpack-go.git   	

        cd $USER_HOME_DIR
        
        cmd="/MONGO_DB_URL=/c\MONGO_DB_URL=\"$eth0_ip\""
	sed -ir "$cmd" $GOPATH/src/github.com/bottos-project/magiccube/service/node/config/config.go
        cmd="/WALLET_IP=/c\WALLET_IP=\"$eth0_ip\""
        sed -ir "$cmd" $GOPATH/src/github.com/bottos-project/magiccube/service/node/config/config.go
    	
	cp -rf $GOPATH/src/github.com/bottos-project/bottos/chainconfig.json $USER_HOME_DIR/opt/go/bin
	cp -rf $GOPATH/src/github.com/bottos-project/bottos/genesis.json $USER_HOME_DIR/opt/go/bin

        cmd="/option_db/c\\\"option_db\":\"$eth0_ip:27017\","
        sed -ir $cmd $USER_HOME_DIR/opt/go/bin/chainconfig.json
        cmd="/api_service_enable/c\\\"api_service_enable\":true,"
        sed -ir $cmd $USER_HOME_DIR/opt/go/bin/chainconfig.json
        
	cp -rf $GOPATH/src/github.com/bottos-project/magiccube/service/node/config/config.json $USER_HOME_DIR/opt/go/bin/
        cmd="/\"ipAddr\"/c\\\"ipAddr\":\"$eth0_ip\","
        sed -ir $cmd $USER_HOME_DIR/opt/go/bin/config.json
        cmd="/walletIP/c\\\"walletIP\":\"$eth0_ip\","
        sed -ir $cmd $USER_HOME_DIR/opt/go/bin/config.json
        chmod 777 $GOPATH/src/github.com/bottos-project/* -R
    #fi
    cp -rf $GOPATH/src/github.com/bottos-project/magiccube/service/node/scripts/build.sh $USER_HOME_DIR/opt/go/bin/ 2>/dev/null
    cp -rf $GOPATH/src/github.com/bottos-project/magiccube/vendor/* $GOPATH/src  2>/dev/null 
   	
    cp -rf $GOPATH/src/github.com/bottos-project/magiccube/config $USER_HOME_DIR/opt/go/bin
    
    cp -rf $GOPATH/src/github.com/bottos-project/bottos/corelog.xml $USER_HOME_DIR/opt/go/bin/core 2>/dev/null

    cp -rf $USER_HOME_DIR/opt/go/bin/*.json $USER_HOME_DIR/ 2>/dev/null
    cp -rf $USER_HOME_DIR/opt/go/bin/*.json $USER_HOME_DIR/opt/go/bin/core 2>/dev/null
    cp -rf $USER_HOME_DIR/opt/go/bin/*.json $USER_HOME_DIR/opt/go/bin/core/cmd_dir 2>/dev/null    

    chown bottos:bottos $GOPATH/src -R   
 
    echo "\n Cloning all is done. Please try ./startup.sh buildstart for auto-build then, or try ./startup.sh start for directly start."
    
    return	
}

function setenv()
{
    export GOPATH=/home/bottos/mnt/bottos
    export GOROOT=/usr/lib/go
    
    cp -rf $GOPATH/src/github.com/bottos-project/bottos/bcli/cliconfig.json $USER_HOME_DIR/opt/go/bin/core/cmd_dir
    cp -rf $GOPATH/src/github.com/bottos-project/bottos/bcli/cliconfig.json $USER_HOME_DIR/opt/go/bin
}

function build_all_modules()
{
    /usr/lib/go/bin/./go build github.com/bottos-project/bottos
    /usr/lib/go/bin/./go build github.com/bottos-project/magiccube/service/node
    /usr/lib/go/bin/./go build github.com/bottos-project/magiccube/service/user/useApi
    /usr/lib/go/bin/./go build github.com/bottos-project/magiccube/service/user
    /usr/lib/go/bin/./go build github.com/bottos-project/magiccube/service/asset
    /usr/lib/go/bin/./go build github.com/bottos-project/magiccube/service/storage
    /usr/lib/go/bin/./go build github.com/bottos-project/magiccube/service/requirement/reqApi
    /usr/lib/go/bin/./go build github.com/bottos-project/magiccube/service/requirement
    /usr/lib/go/bin/./go build github.com/bottos-project/magiccube/service/exchange
    /usr/lib/go/bin/./go build github.com/bottos-project/magiccube/service/dashboard/dasApi
    /usr/lib/go/bin/./go build github.com/bottos-project/magiccube/service/dashboard   
    /usr/lib/go/bin/./go build github.com/bottos-project/magiccube/service/data
    /usr/lib/go/bin/./go build github.com/bottos-project/magiccube/service/data/datApi
    
    cp -f bottos      $USER_HOME_DIR/opt/go/bin/core/ 2>/dev/null
    
    path=`pwd`
    if [ $path != "$USER_HOME_DIR/opt/go/bin" ];
    then
        cp -f node        $USER_HOME_DIR/opt/go/bin
        cp -f user        $USER_HOME_DIR/opt/go/bin
        cp -f useApi     $USER_HOME_DIR/opt/go/bin
        cp -f asset       $USER_HOME_DIR/opt/go/bin
        cp -f storage     $USER_HOME_DIR/opt/go/bin
        cp -f requirement $USER_HOME_DIR/opt/go/bin
        cp -f reqApi      $USER_HOME_DIR/opt/go/bin
        cp -f exchange    $USER_HOME_DIR/opt/go/bin
        cp -f dasApi      $USER_HOME_DIR/opt/go/bin
        cp -f dashboard   $USER_HOME_DIR/opt/go/bin
        cp -f data        $USER_HOME_DIR/opt/go/bin
        cp -f datApi      $USER_HOME_DIR/opt/go/bin
    fi

    cd $USER_HOME_DIR/opt/go/bin
}

function swcheck () {
    rm -rf $GOPATH/src/* 2>/dev/null
    rm -rf $OPT_GO_BIN/* 2>/dev/null	
    rm -rf magiccube     2>/dev/null
    rm -rf /home/bottos/.cache 2>/dev/null
    cd $USER_HOME_DIR
    #if [ ! -d $MINIO_SHR ] || [ ! -d $MINIO_COF ];
    #then	
        git clone https://github.com/bottos-project/magiccube.git 
        cp ./magiccube/vendor/minio /usr/local/bin
        cp ./magiccube/vendor/minio .
	if [ ! -e /usr/local/bin/minio ];
	then
	    miniocheck
		setminio
	fi
        cp -rf ./magiccube/vendor/minio $GOPATH/src
    	rm -rf ./magiccube 2>/dev/null    
    	
	mkdir -p $MINIO_SHR 2>/dev/null
        mkdir -p $MINIO_COF 2> /dev/null
        chown bottos:bottos $MINIO_SHR   -R	
        chown bottos:bottos $MINIO_COF   -R	
    #fi
    
    mkdir -p $OPT_GO_BIN 2>/dev/null
    mkdir -p $OPT_GO_BIN/core 2>/dev/null
    mkdir -p $OPT_GO_BIN/core/cmd_dir 2>/dev/null

    mkdir -p $USER_HOME_DIR/opt 2>/dev/null	
    mkdir -p $USER_HOME_DIR/mnt/bottos/src/github.com/bottos-project/ 2>/dev/null
    mkdir /home/bto  2>/dev/null   
    
    chown bottos:bottos /usr/bin/mongo* 
    chown bottos:bottos /home/bto   -R
    chown bottos:bottos $OPT_GO_BIN -R	
    chown bottos:bottos $OPT_GO_BIN/* -R	
    chown bottos:bottos $USER_HOME_DIR/mnt/bottos/src/github.com/bottos-project -R	
    chown bottos:bottos $USER_HOME_DIR/opt   -R
    chown bottos:bottos $USER_HOME_DIR/opt/* -R
    chown bottos:bottos $USER_HOME_DIR/mnt/bottos -R	
    chown bottos:bottos $USER_HOME_DIR/mnt/bottos/* -R	

    wget https://storage.googleapis.com/golang/go1.10.1.linux-amd64.tar.gz --directory-prefix=$USER_HOME_DIR
    
    if [ ! -f $USER_HOME_DIR/go1.10.1.linux-amd64.tar.gz* ]; then
        echo "Download golang package failed!"
        exit 1
    fi	
    tar -xzvf $USER_HOME_DIR/go1.10.1.linux-amd64.tar.gz -C /usr/local
    tar -xzvf $USER_HOME_DIR/go1.10.1.linux-amd64.tar.gz -C /usr/lib
    
    rm -rf $USER_HOME_DIR/go1.10.1.linux-amd64.tar.gz 2>&1>/dev/null
}

function setgopath() {
    cmd=$(sed  -n "/GOPATH/p" /etc/profile |wc -l)
    if [ $cmd -lt 1 ]; then
       sed -i '$a\export GOPATH="/home/bottos/mnt/bottos"' /etc/profile
    else
       $(sed -ir "/GOPATH/c\export GOPATH=\"/home/bottos/mnt/bottos\"" /etc/profile)
    fi
	
    cmd=$(sed  -n "/GOROOT/p" /etc/profile |wc -l)
   
     if [ $cmd -lt 1 ]; then
       sed -i "\$a\\export GOROOT=\"$SYS_GOROOT\"" /etc/profile
    else
       $(sed -ir "/GOROOT/c\export GOROOT=\"/usr/lib/go\"" /etc/profile)
    fi
    
    cmd=$(sed  -n "/GOPATH/p" $USER_HOME_DIR/.bashrc |wc -l)
    if [ $cmd -lt 1 ]; then
       sed -i '$a\export GOPATH="/home/bottos/mnt/bottos"' $USER_HOME_DIR/.bashrc
    fi
    
    cmd=$(sed  -n "/GOROOT/p" $USER_HOME_DIR/.bashrc |wc -l)
    if [ $cmd -lt 1 ]; then
        sed -i "\$a\export GOROOT=\"$SYS_GOROOT\"" $USER_HOME_DIR/.bashrc
    fi
    
    cmd=$(sed -n "/export PATH/p" $USER_HOME_DIR/.bashrc |wc -l)
    if [ $cmd -lt 1 ]; then
        sed -i '$a\export PATH=$PATH:/usr/lib/go/bin' $USER_HOME_DIR/.bashrc
    fi
    
    cmd=$(sed -n "/export PATH/p" /etc/profile |wc -l)
    if [ $cmd -lt 1 ]; then
        sed -i '$a\export PATH=$PATH:/usr/lib/go/bin' /etc/profile
    fi
    
    eth0_ip=$SERVER_IPADR
    cmd="/bind_ip/c\bind_ip=$eth0_ip"
    sed -ir $cmd /etc/mongodb.conf

    export GOPATH="/home/bottos/mnt/bottos"
    export GOROOT="/usr/lib/go"

    if [ $(echo $PATH|grep "\/usr\/lib\/go"|wc -l) -lt 1 ]; then
        export PATH=$PATH:/usr/lib/go/bin
    fi	
}

#main
case $1 in
    "update")
        usercheck "bottos"
        download_git_newcode        
        ;;
    "buildstart")
        usercheck "bottos"
        stopserv
        
        setenv
        build_all_modules
        varcheck
        service mongodb start
        startserv
        ;;
    "start")
        usercheck "bottos"
        stopserv
          
        setenv
        varcheck
        service mongodb start
        startserv 
        ;;
    "stop")
        usercheck "bottos"
        stopserv
        ;;
    "deploy")
        usercheck "root"
        varcheck   
        prepcheck
        sslcheck
	swcheck
        setgopath
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

#exit 0

