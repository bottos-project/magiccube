/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Data Exchange Client
  Created by Developers Team of Bottos.

  This program is free software: you can distribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with Bottos. If not, see <http://www.gnu.org/licenses/>.
 */
package sqlite

import (
	"fmt"
	"testing"

	"github.com/bottos-project/magiccube/service/storage/util"
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
