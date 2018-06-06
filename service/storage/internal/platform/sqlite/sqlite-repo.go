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
	"database/sql"
	"errors"

	"fmt"
)

// DefaultDbpath db path
const DefaultDbpath string = "./bottos.db"

// SqlTableExists is or not
const SqlTableExists string = "select count(*) from sqlite_master where type= 'table' and name="

// SqliteRepository struct
type SqliteRepository struct {
}

// NewSqliteRepository creates a new SqliteRepository
func NewSqliteRepository() *SqliteRepository {
	return &SqliteRepository{}
}

// SqliteContext struct
type SqliteContext struct {
	db *sql.DB
}

// ConnectDB ...
func ConnectDB() (*SqliteContext, error) {
	db, err := sql.Open("sqlite3", DefaultDbpath)
	if err != nil {
		return nil, errors.New("ConnectDB sql open sqlite3 database failed")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.New("ConnectDB db ping failed")
	}
	return &SqliteContext{db}, nil
}

// IsTableExist is or not
func (c *SqliteContext) IsTableExist(table string) bool {
	sql := SqlTableExists + "'" + table + "';"
	var num uint32
	rows, err := c.db.Query(sql)
	if err != nil {
		fmt.Println(sql)
		return false
	}
	//select count(*) from sqlite_master where type= 'table' and name='sync';
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			return false
		}
		if num == 0 {
			return false
		}
	}
	return true
}
