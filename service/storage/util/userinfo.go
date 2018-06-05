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

package util
// UserInfo struct 
type UserInfo struct {
	Username     string `db:"username"`
	Accountname  string `db:"accountname"`
	Ownerpubkey  string `db:"ownerpubkey"`
	Activepubkey string `db:"activepubkey"`
	User_type    string `db:"user_type"`
	Role_type    string `db:"role_type"`
	Info         string `db:"info"`
}
//UserDBInfo struct
type UserDBInfo struct {
	Username       string `db:"username"`
	Accountname    string `db:"accountname"`
	Ownerpubkey    string `db:"owner_pub_key"`
	Activepubkey   string `db:"active_pub_key"`
	EncyptedInfo   string `db:"encypted_info"`
	UserType       string `db:"user_type"`
	RoleType       string `db:"role_type"`
	CompanyName    string `db:"company_name"`
	CompanyAddress string `db:"company_address"`
}
//TokenDBInfo struct
type TokenDBInfo struct {
	Username   string `db:"username"`
	Token      string `db:"token"`
	InsertTime int64  `db:"insert_time"`
}
// InserUserInfoSql string
const InserUserInfoSql string = "insert into userinfo(Username, Accountname, Ownerpubkey,Activepubkey,EncyptedInfo,UserType,RoleType,CompanyName,CompanyAddress) values(?,?,?,?,?,?,?,?,?)"
//QueryUserInfoSql string

const QueryUserInfoSql string = "select * from userinfo"
//InsertUserTokenSql string
const InsertUserTokenSql string = "insert into tokeninfo(Username, Token,InsertTime) values(?,?,?)"
//DefualtAgingTime int64
const DefualtAgingTime int64 = 30 * 60
