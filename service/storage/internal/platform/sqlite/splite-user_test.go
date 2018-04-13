package sqlite
import(
	"testing"
	"github.com/code/bottos/service/storage/util"
	"fmt"
)

func TestSqliteRepository_CallInsertUserInfo(t *testing.T) {
	var ins SqliteRepository
	value := util.UserDBInfo{
	        Username : "usermng",
		Accountname:"usermng",
		Ownerpubkey: "0xabcd1",
		Activepubkey: "0xabcd2",
		EncyptedInfo :"0xefgh",
		UserType :"person",
		RoleType:"provider",
		CompanyName :"tuzi network",
		CompanyAddress:"zhangjiang",
	}
	err := ins.CallInsertUserInfo(value)
	if err != nil { //try a unit test on function
		t.Error("Insert user failed")
	} else {
		t.Log("success")
	}
}
func TestSqliteRepository_QueryUserInfo(t *testing.T) {
	var ins SqliteRepository
	if value, e := ins.CallGetUserInfo("usermng"); value != nil || e != nil { //try a unit test on function
		t.Error("Insert user failed")
	} else {
		fmt.Println(value.Accountname)
	}
}