package bean


type CoreCommonReturn struct {
	Errcode int64 		`json:"errcode"`
	Msg     string  	`json:"msg"`
	Result  interface{} `json:"result"`
}

type TxBean struct {
	Version     uint64  `protobuf:"varint,1,opt,name=version" json:"version"`
	CursorNum   uint64  `protobuf:"varint,2,opt,name=cursor_num,json=cursorNum" json:"cursor_num"`
	CursorLabel uint64  `protobuf:"varint,3,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	Lifetime    uint64  `protobuf:"varint,4,opt,name=lifetime" json:"lifetime"`
	Sender      string  `protobuf:"bytes,5,opt,name=sender" json:"sender"`
	Contract    string  `protobuf:"bytes,6,opt,name=contract" json:"contract"`
	Method      string  `protobuf:"bytes,7,opt,name=method" json:"method"`
	Param       string `protobuf:"bytes,8,opt,name=param" json:"param"`
	SigAlg      uint64 `protobuf:"varint,9,opt,name=sig_alg,json=sigAlg" json:"sig_alg"`
	Signature   string `protobuf:"bytes,10,opt,name=signature" json:"signature"`
}

type UserTokenBean struct {
	Username string `bson:"username"`
	Token    string `bson:"token"`
	Ctime    int64  `bson:"ctime"`
}

type Did struct {
	Didid string
	Didstr string
}

