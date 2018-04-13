package bean

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type RequirementBean struct {
	ID                 bson.ObjectId `bson:"_id,omitempty"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data struct {
		DataReqID string `bson:"data_req_id"`
		BasicInfo struct {
			UserName        string `bson:"user_name"`
			SessionID       string `bson:"session_id"`
			RequirementName string `bson:"requirement_name"`
			FeatureTag      uint64 `bson:"feature_tag"`
			SamplePath      string `bson:"sample_path"`
			SampleHash      string `bson:"sample_hash"`
			ExpireTime      uint32 `bson:"expire_time"`
			Price           uint64 `bson:"price"`
			Description     string `bson:"description"`
			PublishDate     uint32 `bson:"publish_date"`
			Signature       string `bson:"signature"`
		} `bson:"basic_info"`
	} `bson:"data"`
	CreatedAt time.Time `bson:"createdAt"`
}

type NodeBean struct {
	ID                 bson.ObjectId `bson:"_id"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data struct {
		NodeID string `bson:"node_id"`
		BasicInfo struct {
			NodeIP      string `bson:"node_ip"`
			NodePort    string `bson:"node_port"`
			NodeAddress string `bson:"node_address"`
		} `bson:"basic_info"`
	} `bson:"data"`
	CreatedAt time.Time `bson:"createdAt"`
}

type TxBean struct {
	ID                 bson.ObjectId `bson:"_id"`
	BlockNum           uint64        `bson:"block_num"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data struct {
		DataDealID string `bson:"data_deal_id"`
		BasicInfo struct {
			UserName  string `bson:"user_name"`
			SessionID string `bson:"session_id"`
			AssetID   string `bson:"asset_id"`
			RandomNum int    `bson:"random_num"`
			Signature string `bson:"signature"`
		} `bson:"basic_info"`
	} `bson:"data"`
	CreatedAt time.Time `bson:"createdAt"`
}

type TransferBean struct {
	ID                 bson.ObjectId `bson:"_id"`
	BlockNum           uint64        `bson:"block_num"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data struct {
		From     string `bson:"from"`
		To       string `bson:"to"`
		Quantity uint64 `bson:"quantity"`
	} `bson:"data"`
	CreatedAt time.Time `bson:"createdAt"`
}

type AssetBean struct {
	ID                 bson.ObjectId `bson:"_id"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data struct {
		AssetID string `bson:"asset_id"`
		BasicInfo struct {
			UserName    string `bson:"user_name"`
			SessionID   string `bson:"session_id"`
			AssetName   string `bson:"asset_name"`
			//FeatureTag  uint64 `bson:"feature_tag"`
			AssetType   string `bson:"asset_type"`
			FeatureTag1 string `bson:"feature_tag1"`
			FeatureTag2 string `bson:"feature_tag2"`
			FeatureTag3 string `bson:"feature_tag3"`
			SamplePath  string `bson:"sample_path"`
			SampleHash  string `bson:"sample_hash"`
			StoragePath string `bson:"storage_path"`
			StorageHash string `bson:"storage_hash"`
			ExpireTime  uint32 `bson:"expire_time"`
			Price       uint64 `bson:"price"`
			Description string `bson:"description"`
			UploadDate  uint32 `bson:"upload_date"`
			Signature   string `bson:"signature"`
		} `bson:"basic_info"`
	} `bson:"data"`
	CreatedAt time.Time `bson:"createdAt"`
}

type FileBean struct {
	ID                 string        `bson:"_id"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data struct {
		FileHash string `bson:"file_hash"`
		BasicInfo struct {
			UserName   string `bson:"user_name"`
			SessionID  string `bson:"session_id"`
			FileSize   uint64 `bson:"file_size"`
			FileName   string `bson:"file_name"`
			FilePolicy string `bson:"file_policy"`
			FileNumber uint64 `bson:"file_number"`
			Signature  string `bson:"signature"`
			AuthPath   string `bson:"auth_path"`
		} `bson:"basic_info"`
	} `bson:"data"`
	CreatedAt time.Time `bson:"createdAt"`
}

type FavoriteBean struct {
	ID                 bson.ObjectId `bson:"_id"`
	BlockNum           int           `bson:"block_num"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data struct {
		UserName  string `bson:"user_name"`
		SessionID string `bson:"session_id"`
		OpType    string `bson:"op_type"`
		GoodsType string `bson:"goods_type"`
		GoodsID   string `bson:"goods_id"`
		Signature string `bson:"signature"`
	} `bson:"data"`
	CreatedAt string `bson:"createdAt"`
}

type BlockBean struct {
	ID                    bson.ObjectId `bson:"_id,omitempty"`
	BlockID               string        `bson:"block_id"`
	BlockNum              uint64        `bson:"block_num"`
	PrevBlockID           string        `bson:"prev_block_id"`
	ProducerAccountID     string        `bson:"producer_account_id"`
	Timestamp             time.Time     `bson:"timestamp"`
	TransactionMerkleRoot string        `bson:"transaction_merkle_root"`
	Transactions          []string      `bson:"transactions"`
	CreatedAt             time.Time     `bson:"createdAt"`
}

type RecordNumLog struct {
	ID                    	bson.ObjectId 	`bson:"_id,omitempty"`
	TxNum              		int   			`bson:"tx_num"`
	TxAmount 				uint64   		`bson:"tx_amount"`
	RequirementNum			int   			`bson:"requirement_num"`
	AssetNum				int   			`bson:"asset_num"`
	AccountNum  			int   			`bson:"account_num"`
	Date					string   		`bson:"date"`
	Timestamp				int   			`bson:"timestamp"`
	CreatedAt 			  	time.Time		`bson:"createdAt"`
}


type PurchaseMesssageBean struct {
	ID                 bson.ObjectId `bson:"_id"`
	BlockNum           uint64        `bson:"block_num"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data struct {
		DataDealID string `bson:"data_deal_id"`
		BasicInfo struct {
			UserName  string `bson:"user_name"`
			SessionID string `bson:"session_id"`
			AssetID   string `bson:"asset_id"`
			RandomNum int    `bson:"random_num"`
			Signature string `bson:"signature"`
		} `bson:"basic_info"`
	} `bson:"data"`
	CreatedAt time.Time `bson:"createdAt"`
}

type DataPreSaleBean struct {
	ID                 bson.ObjectId `bson:"_id"`
	BlockNum           uint64        `bson:"block_num"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data struct {
		DataDealID string `bson:"data_deal_id"`
		BasicInfo struct {
			UserName    string `bson:"user_name"`
			SessionID   string `bson:"session_id"`
			AssetID     string `bson:"asset_id"`
			AssetName   string `bson:"asset_name"`
			DataReqID   string `bson:"data_req_id"`
			DataReqName string `bson:"data_req_name"`
			Consumer    string `bson:"consumer"`
			RandomNum   int    `bson:"random_num"`
			Signature   string `bson:"signature"`
		} `bson:"basic_info"`
	} `bson:"data"`
	CreatedAt time.Time `bson:"createdAt"`
}
