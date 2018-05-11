package bean


type CoreCommonReturn struct {
	Errcode int64 `json:"errcode"`
	Msg     string  `json:"msg"`
	Result  interface{}
}

type UserTokenBean struct {
	Username string `bson:"username"`
	Token    string `bson:"token"`
	Ctime    int64  `bson:"ctime"`
}
