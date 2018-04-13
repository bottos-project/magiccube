package sqlite
//import(
//	"testing"
//	"fmt"
//)
//
//func TestSqliteRepository_CallInsertUserToken(t *testing.T) {
//	var ins SqliteRepository
//
//	code,err := ins.CallInsertUserToken("abcdef","12345646")
//	if err != nil { //try a unit test on function
//		t.Error("Insert user failed")
//	} else {
//		t.Log("success")
//		if code ==0 {
//			t.Log("failed")
//		}else {
//			t.Log("success")
//		}
//	}
//}
//func TestSqliteRepository_CallGetUserToken(t *testing.T) {
//	var ins SqliteRepository
//	if value, e := ins.CallGetUserToken("123456"); value != 1 || e != nil { //try a unit test on function
//		t.Error("Insert user failed")
//	} else {
//		fmt.Println(value)
//	}
//}