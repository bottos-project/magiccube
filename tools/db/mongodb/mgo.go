package mgo

import(
	"gopkg.in/mgo.v2"
)

const URL = "47.98.47.148:27017" //mongodb连接字符串
var mgoSession *mgo.Session

func Session() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(URL)
		if err != nil {
			panic(err)
		}
	}

	mgoSession.SetMode(mgo.Monotonic, true)
	mgoSession.SetPoolLimit(200)
	return mgoSession.Clone()
}


