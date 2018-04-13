package mongodb

import (
	"fmt"
	"testing"
)

func TestMongodbAsset_CallGetDayAssetNum(t *testing.T) {
	ins := MongoRepository{"47.98.47.148:27017"}

	code, err := ins.CallGetDayAssetNum(6)
	fmt.Println(code)
	fmt.Println(err)
	fmt.Println("code")
}
