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
 
package rqlite

//import (
//	"errors"
//	"github.com/rqlite/gorqlite"
//	"time"
//	"fmt"
//	"os"
//)

//type RqliteRepository struct {
//	rqlEndpoint  string
//	rqlAccessKey string
//	rqlSecretKey string
//}

//// NewMinioRepository creates a new MinioRepository
//func NewRqliteRepository(endpoint string, accessKey string, secretKey string) *RqliteRepository {
//	return &RqliteRepository{rqlEndpoint: endpoint, rqlAccessKey: accessKey, rqlSecretKey: secretKey}
//}
//func init(endporint string) *Connection {

//	conn = gorliqte.Open(endporint)

//	conn.SetConsistencyLevel("none")
//	// set the http timeout.  Note that rqlite has various internal timeouts, but this
//	// timeout applies to the http.Client and its work.  It is measured in seconds.
//	conn.SetTimeout(10)
//	return &conn;

//}

//func insertUserInfo(userInfoJson string){

//	// simulate database/sql Prepare()
//	statements := make ([]string,0)
//	pattern := "INSERT INTO secret_agents(id, hero_name, abbrev) VALES (%d, '%s', '%3s')"
//	statements = append(statements,fmt.Sprintf(pattern,125718,"Speed Gibson","Speed"))
//	statements = append(statements,fmt.Sprintf(pattern,209166,"Clint Barlow","Clint"))
//	statements = append(statements,fmt.Sprintf(pattern,44107,"Barney Dunlap","Barney"))
//	results, err := conn.Write(statements)

//	// now we have an array of []WriteResult

//	for n, v := range WriteResult {
//		fmt.Printf("for result %d, %d rows were affected\n",n,v.RowsAffected)
//		if ( v.Err != nil ) {
//			fmt.Printf("   we have this error: %s\n",v.Err.Error())
//		}
//	}

//	// or if we have an auto_increment column
//	res, err := conn.WriteOne("INSERT INTO foo (name) values ('bar')")
//	fmt.Printf("last insert id was %d\n",res.LastInsertID)

//	// just like database/sql, you're required to Next() before any Scan() or Map()

//	// note that rqlite is only going to send JSON types - see the encoding/json docs
//	// which means all numbers are float64s.  gorqlite will convert to int64s for you
//	// because it is convenient but other formats you will have to handle yourself

//	var id int64
//	var name string
//	rows, err := conn.QueryOne("select id, name from secret_agents where id > 500")
//	fmt.Printf("query returned %d rows\n",rows.NumRows)
//	for rows.Next() {
//		err := response.Scan(&id, &name)
//		fmt.Printf("this is row number %d\n",response.RowNumber)
//		fmt.Printf("there are %d rows overall%d\n",response.NumRows)
//	}

//	// just like WriteOne()/Write(), QueryOne() takes a single statement,
//	// while Query() takes a []string.  You'd only use Query() if you wanted
//	// to transactionally group a bunch of queries, and then you'd get back
//	// a []QueryResult

//	// alternatively, use Next()/Map()

//	for rows.Next() {
//		m, err := response.Map()
//		// m is now a map[column name as string]interface{}
//		id := m["name"].(float64) // the only json number type
//		name := m["name"].(string)
//	}

//	// get rqlite cluster information
//	leader, err := conn.Leader()
//	// err could be set if the cluster wasn't answering, etc.
//	fmt.Println("current leader is"leader)
//	peers, err := conn.Peers()
//	for n, p := range peers {
//		fmt.Printf("cluster peer %d: %s\n",n,p)
//	}

//	// turn on debug tracing to the io.Writer of your choice.
//	// gorqlite will verbosely write bery granular debug information.
//	// this is similar to perl's DBI->Trace() facility.
//	// note that this is done at the package level, not the connection
//	// level, so you can debug Open() etc. if need be.

//	f, err := os.OpenFile("/tmp/deep_insights.log",OS_RDWR|os.O_CREATE|os.O_APPEND,0644)
//	gorqlite.TraceOn(f)

//	// change my mind and watch the trace
//	gorqlite.TraceOn(os.Stderr)

//	// turn off
//	gorqlite.TraceOff()
//}
