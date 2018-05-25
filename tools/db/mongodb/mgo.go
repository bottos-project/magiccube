package mgo

import(
	"gopkg.in/mgo.v2"
	"github.com/bottos-project/magiccube/config"
)

var mgoSession *mgo.Session

func Session() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(config.BASE_MONGODB_ADDR)
		if err != nil {
			panic(err)
		}
	}

	mgoSession.SetMode(mgo.Monotonic, true)
	mgoSession.SetPoolLimit(200)
	return mgoSession.Clone()
}


