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
import(
	"testing"
	"github.com/bottos-project/magiccube/service/storage/util"
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