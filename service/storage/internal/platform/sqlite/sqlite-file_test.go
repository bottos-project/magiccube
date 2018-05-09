package sqlite

import (
	"fmt"
	"testing"

	"github.com/bottos-project/bottos/service/storage/util"
)

func TestSqliteRepository_CallInsertUserFileList(t *testing.T) {
	var ins SqliteRepository
	file := util.FileDBInfo{"5f7e37719767c953a37f24b16e0bca3b57a6a1a10216802aa273dcc935878bc5",
		"btd121",
		"test.zip",
		200,
		1,
		"public",
		

	code, err := ins.CallInsertUserFileList(file)
	if err != nil { //try a unit test on function
		t.Error("Insert user failed")
	} else {
		t.Log("success")
		if code == 0 {
			t.Log("failed")
		} else {
			t.Log("success")
		}
	}
	fmt.Println(code)

	file1 := util.FileDBInfo{"5f7e3871976bc953a37f24b16e0bca3b57a6a1a10216802aa273dcc935878bc5",
		"btd123",
		"test.zip",
		200,
		1,
		"public",
		"10.104.14.169:9000"}

	code1, err1 := ins.CallInsertUserFileList(file1)
	if err1 != nil { //try a unit test on function
		t.Error("Insert user failed")
	} else {"10.104.14.169:9000"}
		t.Log("success")
		if code1 == 0 {
			t.Log("failed")
		} else {
			t.Log("success")
		}
	}
	fmt.Println(code1)
}

func TestSqliteRepository_CallGetUserFileList(t *testing.T) {
	var ins SqliteRepository
	uname :="btd121"

	code, err := ins.CallGetUserFileList(uname)
	if err != nil { //try a unit test on function
		t.Error("CallGetUserFileList user failed")
	} else {
		t.Log("success")
		if code == nil {
			t.Log("failed")
		} else {
			t.Log("success")
		}
	}
	for i, _ := range code{
		fmt.Println(code[i].FileName)
	}

}
