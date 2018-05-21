package bean

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type CoreBaseReturn struct {
	Errcode int64 		`json:"errcode"`
	Msg     string  	`json:"msg"`
	Result  interface{} `json:"result"`
}

type CoreCommonReturn struct {
	Errcode int64 		`json:"errcode"`
	Msg     string  	`json:"msg"`
	Result  struct {
		Trx struct {
			Version     uint32 `json:"version"`
			CursorNum   uint32 `json:"cursor_num"`
			CursorLabel uint32  `json:"cursor_label"`
			Lifetime    uint64 `json:"lifetime"`
			Sender      string `json:"sender"`
			Contract    string `json:"contract"`
			Method      string `json:"method"`
			Param       string `json:"param"`
			SigAlg      uint32 `json:"sig_alg"`
			Signature   string `json:"signature"`
		} `json:"trx"`
		TrxHash string `json:"trx_hash"`
	} `json:"result"`
}

type TxBean struct {
	Version     uint32  `protobuf:"varint,1,opt,name=version" json:"version"`
	CursorNum   uint32  `protobuf:"varint,2,opt,name=cursor_num,json=cursorNum" json:"cursor_num"`
	CursorLabel uint32  `protobuf:"varint,3,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	Lifetime    uint64  `protobuf:"varint,4,opt,name=lifetime" json:"lifetime"`
	Sender      string  `protobuf:"bytes,5,opt,name=sender" json:"sender"`
	Contract    string  `protobuf:"bytes,6,opt,name=contract" json:"contract"`
	Method      string  `protobuf:"bytes,7,opt,name=method" json:"method"`
	Param       string  `protobuf:"bytes,8,opt,name=param" json:"param"`
	SigAlg      uint32  `protobuf:"varint,9,opt,name=sig_alg,json=sigAlg" json:"sig_alg"`
	Signature   string  `protobuf:"bytes,10,opt,name=signature" json:"signature"`
}

type UserTokenBean struct {
	Username string `bson:"username"`
	Token    string `bson:"token"`
	Ctime    int64  `bson:"ctime"`
}

type Did struct {
	Didid string
	Didinfo string
}

type TxPublic struct {
	Sender string `json:"sender"`
}

type Block struct {
	ID                    bson.ObjectId 	`bson:"_id,omitempty"`
	BlockHash             string        	`bson:"block_hash"`
	BlockNumber           uint64        	`bson:"block_number"`
	PrevBlockHash         string        	`bson:"prev_block_hash"`
	Delegate     		  string        	`bson:"delegate"`
	Timestamp             uint64     	    `bson:"timestamp"`
	MerkleRoot 			  string        	`bson:"merkle_root"`
	Transactions          []bson.ObjectId   `bson:"transactions"`
	CreateTime           time.Time     		`bson:"create_time"`
}

type Favorite struct {
	ID          bson.ObjectId   `bson:"_id,omitempty"`
	Contract    string  		`bson:"contract"`
	CursorLabel uint32 			`bson:"cursor_label"`
	CursorNum   uint32			`bson:"cursor_num"`
	Lifetime    uint64 			`bson:"lifetime"`
	Method      string  		`bson:"method"`
	Param       struct {
		Goodsid   string  		`bson:"goodsid"`
		Goodstype string  		`bson:"goodstype"`
		Optype    float64 		`bson:"optype"`
		Username  string  		`bson:"username"`
	} 							`bson:"param"`
	Sender    	string  		`bson:"sender"`
	SigAlg    	uint32 			`bson:"sig_alg"`
	Signature 	string  		`bson:"signature"`
	Version   	uint32 			`bson:"version"`
	CreateTime 	time.Time     	`bson:"create_time"`

}

type Requirement struct {
	ID          bson.ObjectId   `bson:"_id,omitempty"`
	Contract    string  		`bson:"contract"`
	CursorLabel uint32 			`bson:"cursor_label"`
	CursorNum   uint32			`bson:"cursor_num"`
	Lifetime    uint64 			`bson:"lifetime"`
	Method      string  		`bson:"method"`
	Param       struct {
		DataReqId string `bson:"datareqid"`
		Info      struct {
			Description string  `bson:"description"`
			Expiretime  uint64  `bson:"expiretime"`
			Featuretag  uint64  `bson:"featuretag"`
			Optype      uint32  `bson:"optype"`
			Price       uint64  `bson:"price"`
			Reqname     string  `bson:"reqname"`
			Reqtype     uint64  `bson:"reqtype"`
			Samplehash  string  `bson:"samplehash"`
			Username    string  `bson:"username"`
		} `bson:"info"`
	} 							`bson:"param"`
	Sender    	string  		`bson:"sender"`
	SigAlg    	uint32 			`bson:"sig_alg"`
	Signature 	string  		`bson:"signature"`
	Version   	uint32 			`bson:"version"`
	CreateTime 	time.Time     	`bson:"create_time"`
}

type Asset struct {
	ID          bson.ObjectId   `bson:"_id,omitempty"`
	Contract    string  		`bson:"contract"`
	CursorLabel uint32 			`bson:"cursor_label"`
	CursorNum   uint32			`bson:"cursor_num"`
	Lifetime    uint64 			`bson:"lifetime"`
	Method      string  		`bson:"method"`
	Param       struct {
		AssetId string `bson:"assetid"`
		Info    struct {
			AssetName   string  `bson:"assetname"`
			AssetType   string  `bson:"assettype"`
			Description string  `bson:"description"`
			Expiretime  uint64  `bson:"expiretime"`
			Featuretag  string  `bson:"featuretag"`
			Price       uint64  `bson:"price"`
			Samplehash  string  `bson:"samplehash"`
			Samplepath  string  `bson:"samplepath"`
			Signature   string  `bson:"signature"`
			Storagehash string  `bson:"storagehash"`
			Storagepath string  `bson:"storagepath"`
			Uploaddate  uint64  `bson:"uploaddate"`
			Username    string  `bson:"username"`
		} `bson:"info"`
	} 							`bson:"param"`
	Sender    	string  		`bson:"sender"`
	SigAlg    	uint32 			`bson:"sig_alg"`
	Signature 	string  		`bson:"signature"`
	Version   	uint32 			`bson:"version"`
	CreateTime 	time.Time     	`bson:"create_time"`
}