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
	"log"
	"time"

	"errors"

	"github.com/bottos-project/magiccube/service/storage/util"
	_ "github.com/mattn/go-sqlite3"

	//	"github.com/bottos-project/magiccube/service/storage/proto"
)

func (c *SqliteContext) createUser() {
	sqlStmt := `create table userinfo (Username VARCHAR(64) PRIMARY KEY,
		Accountname VARCHAR(64),
		Ownerpubkey VARCHAR(300),
		Activepubkey VARCHAR(300),
		EncyptedInfo VARCHAR(300),
		UserTypes VARCHAR(32),
		RoleType VARCHAR(32) ,
		CompanyName VARCHAR(32) ,
		CompanyAddress VARCHAR(32));
		`
	_, err := c.db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *SqliteContext) insertUserInfo(dbtag util.UserDBInfo) error {

	if !c.IsTableExist("userinfo") {
		c.createUser()
	}
	defer c.db.Close()

	stmt, err := c.db.Prepare(util.InserUserInfoSql)
	if err != nil {
		return errors.New("insertUserInfo sql insert sqlite3 database failed")
	}

	res, err := stmt.Exec(dbtag.Username, dbtag.Accountname, dbtag.Ownerpubkey, dbtag.Activepubkey, dbtag.EncyptedInfo, dbtag.UserType, dbtag.RoleType, dbtag.CompanyName, dbtag.CompanyAddress)
	if err != nil {
		return errors.New("insertUserInfo sql exec sqlite3 database failed")
	}

	_, err = res.RowsAffected()
	if err != nil {
		return errors.New("insertUserInfo sql raws affected failed")
	}
	return nil

}

// Read
func (c *SqliteContext) readOne(user string) (*util.UserDBInfo, error) {
	rows, err := c.db.Query("SELECT * FROM userinfo where Username=" + user)
	if err != nil {
		return nil, errors.New("readOne sql query failed")
	}
	defer rows.Close()

	for rows.Next() {
		dbtag := new(util.UserDBInfo)
		err := rows.Scan(&dbtag.Username, &dbtag.Accountname, &dbtag.Ownerpubkey, &dbtag.Activepubkey, &dbtag.EncyptedInfo, &dbtag.UserType, &dbtag.RoleType, &dbtag.CompanyName, &dbtag.CompanyAddress)
		if err != nil {
			return nil, errors.New("readOne sql scan failed")
		}
		return dbtag, nil
	}
	return nil, nil
}
func (c *SqliteContext) countNum() (uint32, error) {
	var num uint32
	rows, err := c.db.Query("select count(*) from userinfo ")
	if err != nil {
		return 0, errors.New("readOne sql query failed")
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&num)
		if err != nil {
			return 0, errors.New("readOne sql scan failed")
		}
		return num, nil
	}
	return num, nil
}
func (r *SqliteRepository) CallInsertUserInfo(value util.UserDBInfo) error {

	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return errors.New("connectDB failed")
	}
	fmt.Println("caaaaaaa")
	err = db.insertUserInfo(value)
	if err != nil {
		return errors.New("insertUserInfo failed")
	}
	return nil
}
func (r *SqliteRepository) CallGetUserInfo(value string) (*util.UserDBInfo, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, errors.New("connectDB failed")
	}
	res, err2 := db.readOne(value)
	if err2 != nil {
		return nil, errors.New("readOne failed")
	}
	return res, nil
}
func (r *SqliteRepository) CallGetUserNum() (uint32, error) {
	db, err := ConnectDB()
	if err != nil {
		return 0, errors.New("connectDB failed")
	}
	res, err2 := db.countNum()
	if err2 != nil {
		return 0, errors.New("countNum failed")
	}
	return res, nil
}

func (c *SqliteContext) insertUserToken(username string, token string) error {

	if !c.IsTableExist("tokeninfo") {
		sqlStmt := `
		create table tokeninfo (Username VARCHAR(64) PRIMARY KEY,
		Token VARCHAR(64) UNIQUE,
		InsertTime INTEGER );
		`
		log.Println("create table tokeninfo")
		_, err := c.db.Exec(sqlStmt)
		if err != nil {
			log.Println(err)
		}
	}
	defer c.db.Close()

	stmt, err := c.db.Prepare(util.InsertUserTokenSql)
	if err != nil {
		log.Println(err)
		return errors.New("Prepare insertUserToken sql insert sqlite3 database failed")
	}
	insertTime := time.Now().Unix()
	res, err := stmt.Exec(username, token, insertTime)
	if err != nil {
		log.Println(err)
		return errors.New("insertUserToken sql exec sqlite3 database failed")
	}

	_, err = res.RowsAffected()
	if err != nil {
		return errors.New("insertUserToken sql raws affected failed")
	}
	return nil

}
func (c *SqliteContext) getToken(username string, token string) (*util.TokenDBInfo, error) {
	if username == "" && token == "" {
		return nil, errors.New("para error")
	}
	sql := "select * from tokeninfo where Token= '" + token + "'or Username = '" + username + "';"
	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, errors.New("getToken sql query failed")
	}
	defer rows.Close()
	fmt.Println(sql)
	if rows.Next() {
		dbtag := new(util.TokenDBInfo)
		err := rows.Scan(&dbtag.Username, &dbtag.Token, &dbtag.InsertTime)
		if err != nil {
			return nil, errors.New("getToken sql scan failed")
		}
		return dbtag, nil
	}
	return nil, nil
}
func (c *SqliteContext) delToken(username string, token string) (uint32, error) {
	if username == "" && token == "" {
		return 0, errors.New("para error")
	}
	stmt, err := c.db.Prepare("delete from tokeninfo where Username=? or Token=? ")
	if err != nil {
		log.Println(err)
		return 0, errors.New("delToken sqlite3 database failed")
	}
	res, err := stmt.Exec(username, token)
	if err != nil {
		log.Println(err)
		return 0, errors.New("delToken exec sql exec sqlite3 database failed")
	}

	_, err = res.RowsAffected()
	if err != nil {
		return 0, errors.New("delToken sql raws affected failed")
	}
	return 1, nil
}
func (r *SqliteRepository) CallInsertUserToken(username string, token string) (uint32, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return 0, errors.New("connectDB failed")
	}
	err2 := db.insertUserToken(username, token)
	if err2 != nil {
		log.Println(err2)
		return 0, errors.New("insertUserToken failed")
	}
	return 1, nil
}
func (r *SqliteRepository) CallGetUserToken(username string, token string) (*util.TokenDBInfo, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return nil, errors.New("connectDB failed")
	}
	var tokeninfo *util.TokenDBInfo
	tokeninfo, err2 := db.getToken(username, token)
	if err2 != nil {
		log.Println(err2)
		return nil, errors.New("getToken failed")
	}
	return tokeninfo, nil
}

func (r *SqliteRepository) CallDelUserToken(username string, token string) (uint32, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return 0, errors.New("connectDB failed")
	}
	code, err2 := db.delToken(username, token)
	if err2 != nil {
		log.Println(err2)
		return 0, errors.New("delToken failed")
	}
	return code, nil
}

func (r *SqliteRepository) CallTokenAging(timeout int64) error {
	c, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return errors.New("connectDB failed")
	}
	defer c.db.Close()
	currentTime := time.Now().Unix()
	insertTime := currentTime - timeout
	fmt.Println("insertTime")
	stmt, err2 := c.db.Prepare("delete from tokeninfo where InserTime < 1519634294 ")
	if err2 != nil {
		log.Println(err2)
		return errors.New("delete from tokeninfo where InsertTime failed")
	}
	res, err3 := stmt.Exec(insertTime)
	log.Println(err3)

	_, err = res.RowsAffected()
	log.Println(err)

	return nil
}
