#!/usr/bin/env bash

#GO_PATH=`echo $GOPATH |cut -d'=' -f2`
GO_PATH=
CONSUL_PATH=$GO_PATH/bin/linux_386/
SERVER_PATH=$GO_PATH/bin/

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
CORE_PROC_FILE_DIR=${GO_PATH}/bin

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
	if [ -z "$GO_PATH" ] || [ -z "$SERVER_IPADR" ] || [ -z "$SERVER_PORT" ];
	then
		echo -e "\033[31m *ERROR* You have to specify the variable GO_PATH/SERVER_IPADR/SERVER_PORT in the script!!! \033[0m"
		exit 1
	fi
}

function startcore()
{
    if [ ! -e ${CORE_PROC_FILE_DIR}/core ];
	then
		echo -e "\033[31m *ERROR* Fail to find core process under :"${CORE_PROC_FILE_DIR}" !!! \033[0m"
		exit 1
	fi
	
	if  [ ! -d ${CORE_PROC_FILE_DIR} ] || [ ! -f ${CORE_PROC_FILE_DIR}/config.json ] || [ ! -f ${CORE_PROC_FILE_DIR}/genesis.json ];
	then
		echo -e "\033[31m *ERROR* Fail to find core proess's configuration item : config.json or genesis.json under :"${CORE_PROC_FILE_DIR}" \033[0m"
		exit 1
	fi
	
	#GENESIS_JSON=`cat ${GO_PATH}/bin/data-dir/config.ini | grep -e ^genesis-json | awk '{print $3}'`
	#if [ ! -f "$GENESIS_JSON" ];
	#then
	#	echo -e "\033[31m *ERROR* json file :"$GENESIS_JSON" dosn't exist !!! please reconfigure config.ini \033[0m"
	#	exit 1
	#fi
	
	CHK_CORE=`ps -elf | grep "${CORE_PROC_FILE_DIR}/core" | grep -v grep | wc -l`
	if [ "$CHK_CORE" -lt 1 ];
	then 
		#start Core process  , nohup "command" > myout.file 2>&1 &
        nohup ${CORE_PROC_FILE_DIR}/core 2>&1 & 
        #--http-server-address ${SERVER_IPADR}:${CHAIN_PORT} -m mongodb://126.0.0.1/bottos --resync > core.file 2>&1 &
        sleep 3
	fi
    return
}

function stopcore()
{
    RUNNING_CORE_PID=`ps -elf | grep ${CORE_PROC_FILE_DIR}/core | grep -v grep | awk '{print $4}'`
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
        apt-get install golang-go -y
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
    startcore

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

#main
case $1 in
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
        prepcheck
        sslcheck
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
        echo -e "\033[32m you have to input a parameter , Please run the script like ./startup.sh deploy|start|stop|startcore|stopcore|restartcore !!! \033[0m"
        ;;
esac

exit 0
