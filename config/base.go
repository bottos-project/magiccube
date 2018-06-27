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

package config

const (
	// BASE_CHAIN_IP BASE CHAIN IP
	BASE_CHAIN_IP = "139.219.198"
	//BASE_CHAIN_IP = "139.217.206.43"
	// BASE_CHAIN_PORT BASE CHAIN PORT
	BASE_CHAIN_PORT = "8689"
	// BASE_CHAIN_URL BASE CHAIN URL
	BASE_CHAIN_URL = "http://" + BASE_CHAIN_IP + ":" + BASE_CHAIN_PORT + "/"
	// BASE_RPC BASE RPC
	BASE_RPC = "http://" + BASE_CHAIN_IP + ":8080/rpc"
	// BASE_MONGODB_IP BASE_MONGODB_IP
	BASE_MONGODB_IP = "127.0.0.1"
	// BASE_MONGODB_PORT BASE_MONGODB_PORT
	BASE_MONGODB_PORT = "27017"
	// BASE_MONGODB_ADDR BASE_MONGODB_ADDR
	BASE_MONGODB_ADDR = BASE_MONGODB_IP + ":" + BASE_MONGODB_PORT
	// DB_NAME DB_NAME
	DB_NAME = "bottos"
	// BASE_MINIO_IP BASE_MINIO_IP
	BASE_MINIO_IP = BASE_CHAIN_IP
	// BASE_MINIO_PORT BASE_MINIO_PORT
	BASE_MINIO_PORT = "9000"
	// BASE_MINIO_ADDR BASE_MINIO_ADDR
	BASE_MINIO_ADDR = BASE_MINIO_IP + ":" + BASE_MINIO_PORT
	// BASE_MINIO_ACCESS_KEY BASE_MINIO_ACCESS_KEY
	BASE_MINIO_ACCESS_KEY = ""
	// BASE_MINIO_SECRET_KEY BASE_MINIO_SECRET_KEY
	BASE_MINIO_SECRET_KEY = ""
	// BASE_LOG_CONF BASE_LOG_CONF
	BASE_LOG_CONF = "config/log.json"

	//CHAIN_ID
	CHAIN_ID ="00000000000000000000000000000000"

	// TOKEN_EXPIRE_TIME Token expired
	TOKEN_EXPIRE_TIME = 2 * 60 * 60

	//EnableVerification Enable/Disable verification code
	EnableVerification = true
)
