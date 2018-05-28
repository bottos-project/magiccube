package config

const (
	BASE_CHAIN_IP   = "139.217.206.43"
	BASE_CHAIN_PORT = "8689"
	BASE_CHAIN_URL  = "http://" + BASE_CHAIN_IP + ":" + BASE_CHAIN_PORT + "/"
	BASE_RPC        = "http://" + BASE_CHAIN_IP + ":8080/rpc"
	//BASE_RPC        		= "http://127.0.0.1:8080/rpc"

	BASE_MONGODB_PORT = "27017"
	BASE_MONGODB_ADDR = BASE_CHAIN_IP + ":" + BASE_MONGODB_PORT
	DB_NAME           = "bottos"

	BASE_MINIO_IP         = "xx"
	BASE_MINIO_PORT       = "9000"
	BASE_MINIO_ADDR       = BASE_MINIO_IP + ":" + BASE_MINIO_PORT
	BASE_MINIO_ACCESS_KEY = ""
	BASE_MINIO_SECRET_KEY = ""

	BASE_LOG_CONF	= "config/log.json"

	//Token expired
	TOKEN_EXPIRE_TIME = 2 * 60 * 60

	//Enable/Disable verification code
	Enable_verification = true

)
