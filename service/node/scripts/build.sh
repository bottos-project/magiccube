#!/bin/bash

#set the services at startup
CONF_FILE=/etc/rc.local

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
            #echo "setserv ..."
            SERV_CMD=$2"/"$3" > "$3".file 2>&1"
            chmod a+x $2"/"$3
            setserv "$COMMON_STRATUP"
            $SERV_CMD &
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

