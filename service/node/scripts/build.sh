#!/bin/bash

#set the services at startup
CONF_FILE=/etc/rc.local
#specify tmp key.json , it containes owner and active key
JSON_FILE=key.json
#specify the diractory to save keystone file
KEYS_PATH=./
#specify the full path for keystone file
KEYS_FILE=$KEYS_PATH""keystone.bto
#
ENCRYPT_PASSWORD=123456
#set a temp file to record two private keys
TMP_FILE=/tmp/.tmp.txt
ENCRYPT_FILE=/tmp/.encrypt.txt

COMMON_STRATUP="/opt/go/bin/startup.sh start"



function usercheck(){
	SCRIPT_USER=`whoami`
	if [ $SCRIPT_USER != "root" ];
	then
		echo "*ERROR* Please run the script as root !!!"
		exit 1
	fi
	
	return 0
}

function paramcheck() {
	if [ $1 -lt $2 ];
	then
		echo "*ERROR* too little parameter !!!"
		exit 1
	fi
}

function setserv(){
	CHECK_RST=`grep -e "^$1" $CONF_FILE`
	if [ -z "$CHECK_RST" ];
	then
		sed -ir "/^exit 0/i\\$1" $CONF_FILE
	fi
	
	return 0
}

function createkey(){
	EOSD_DIR=$1"/eosd"
	EOSC_DIR=$1"/eosc"
	
	if [ ! -e $EOSD_DIR ] || [ ! -e $EOSC_DIR ];
	then
		echo "*ERROR* command "$EOSD_DIR" or "$EOSC_DIR" doesn't exist !!!"
		exit 1
	fi
	
	if [ ! -f $TMP_FILE ]; 
	then
		echo "*ERROR* The temp file : "$TMP_FILE" doesn't exist , please create account first !!!"
		exit 2
	fi
	
	KEY_NAME=(owner active)

	KEY_PRI_A=`sed -n 1p $TMP_FILE`
	KEY_PRI_B=`sed -n 2p $TMP_FILE`
	
	ENC_KEY_PRI_A=`echo ${KEY_PRI_A} | openssl enc -aes-256-cfb -e -base64 -k ${ENCRYPT_PASSWORD} -salt`
	ENC_KEY_PRI_B=`echo ${KEY_PRI_B} | openssl enc -aes-256-cfb -e -base64 -k ${ENCRYPT_PASSWORD} -salt`
	
	if [ -f $ENCRYPT_FILE ];
	then
		rm -rf $ENCRYPT_FILE
	fi
	
	KEY_INFO=("$ENC_KEY_PRI_A" "$ENC_KEY_PRI_B")

	# generate keystone with private key
	if [ -f $JSON_FILE ];
	then
		echo "Delete old private key json file ..."
		rm -rf $JSON_FILE
	fi
	
	CURRENT_TIME=`date "+%Y-%m-%d %H:%M:%S"`
	
	printf '{\n' | tee >> $JSON_FILE
	printf '\t\"account\":\n'| tee >> $JSON_FILE
	printf '\t\"data\":[\n'| tee >> $JSON_FILE
	for ((i=0;i<${#KEY_INFO[@]};i++))
	do
		printf '\t{\n' | tee >> $JSON_FILE
		num=$(echo $((${#KEY_INFO[@]}-1)))
		if [ "$i" == "${num}" ];
		then
			printf "\t\t\"${KEY_NAME[$i]}\":\"${KEY_INFO[$i]}\"}\n" | tee >> $JSON_FILE
		else
			printf "\t\t\"${KEY_NAME[$i]}\":\"${KEY_INFO[$i]}\"},\n" | tee >> $JSON_FILE
		fi
	done
	
	printf "\t]\n" >> $JSON_FILE
	printf "\t\"timestamp\":\"$CURRENT_TIME\"\n"| tee >> $JSON_FILE
	printf "}" >> $JSON_FILE
	
	if [ ! -f $JSON_FILE ];
	then
		echo "*ERROR* Fail to fine the key file :"$JSON_FILE
		exit 1
	fi
	
	mv $JSON_FILE $2
	
	return 0
}

function createuser(){

	EOS_SERV=$1
	WALLET_SERV=$2
	EOS_PORT=$3
	EOSC_DIR=$4
	PROD_NAME=$5
	USR_NAME=$6
	
	if [ ! -d $EOSC_DIR ];
	then
		echo "*ERROR* Diractory "$EOSC_DIR" doesn't exist !!!"
		exit 1
	fi
	
	KEY_PUB_A=`${EOSC_DIR}/eosc create key | grep -i "Public key:" | awk '{print $3}'`
	KEY_PUB_B=`${EOSC_DIR}/eosc create key | grep -i "Public key:" | awk '{print $3}'`
	
	KEY_PRI_A=`${EOSC_DIR}/eosc create key | grep -i "Private key:" | awk '{print $3}'`
	KEY_PRI_B=`${EOSC_DIR}/eosc create key | grep -i "Private key:" | awk '{print $3}'`
	
	${EOSC_DIR}/eosc --host $EOS_SERV --port $EOS_PORT --wallet-host $WALLET_SERV create account -s $PROD_NAME $USR_NAME $KEY_PUB_A $KEY_PUB_B
	
	if [ -f $TMP_FILE ];
	then
		rm -rf $TMP_FILE
	fi
	
	echo $KEY_PRI_A >> $TMP_FILE
	echo $KEY_PRI_B >> $TMP_FILE

	return 0
}

function chkfile(){
	FILE_PATH=$1
	if [ ! -f $FILE_PATH ];
	then
		echo "*ERROR* File doesn't exist !!!"
		return 1
	fi

	return 0
}

#main
case $1 in
        setserv)
            usercheck
            paramcheck $# 3
            echo "setserv ..."
            SERV_CMD=$2"/"$3" > "$3".file 2>&1"
            chmod a+x $2"/"$3
            setserv "$COMMON_STRATUP"
            $SERV_CMD &
            ;;
        createkey)
            usercheck
            paramcheck $# 2
            echo "createkey ..."
            EOS_PATH=$2
            #optional
            KEY_PATH=$3
            if [ "" == "$KEY_PATH" ];
            then
                KEY_PATH=$KEYS_FILE
            fi
            createkey "$EOS_PATH" "$KEY_PATH"
            ;;
        createuser)
            usercheck
            paramcheck $# 7
            echo "createuser ..."
            EOS_IP=$2
            WALLET_IP=$3
            EOS_PRT=$4
            EOS_PATH=$5
            PROD_USER=$6
            USER_NAME=$7
            createuser "$EOS_IP" "$WALLET_IP" "$EOS_PRT" "$EOS_PATH" "$PROD_USER" "$USER_NAME"
            ;;
        chkfile)
            paramcheck $# 2
            chkfile $2
            ;;			
        usage)
            ;;
        *)
            ;;
esac

exit 0

