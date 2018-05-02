/*The functions of the Sqlite database that provided here are not ready yet，they should never be exposed to users.They are 
in the Bottos's service layer,which delivering database service pluggable to provide users with queries.And we plan to support
it in a future point release. At present, we only support mongodb to provide users with queries.*/
package sqlite

import (
	//	"database/sql"
	"fmt"
	//	_ "github.com/mattn/go-sqlite3"
	"errors"
	"log"

	"github.com/code/bottos/service/storage/util"
)

func (c *SqliteContext) InsertUserfile(file util.FileDBInfo) error {
	if !c.IsTableExist("fileinfo") {
		sqlStmt := `
		create table fileinfo (FileHash VARCHAR(64) PRIMARY KEY,
		Username VARCHAR(64),
		FileName VARCHAR(64),
		FileSize INTEGER,
		FileNumber INTEGER,
		FilePolicy VARCHAR(64),
		AuthorizedStorage VARCHAR(64));
		`
		log.Println("create table fileinfo")
		_, err := c.db.Exec(sqlStmt)
		if err != nil {
			log.Println(err)
		}
	}
	defer c.db.Close()

	stmt, err := c.db.Prepare(util.InsertUserFileSql)
	if err != nil {
		log.Println(err)
		return errors.New("Prepare InsertUserFileSql sql insert sqlite3 database failed")
	}
	res, err := stmt.Exec(file.FileHash, file.Username, file.FileName, file.FileSize, file.FileNumber, file.FilePolicy, file.AuthorizedStorage)
	if err != nil {
		log.Println(err)
		return errors.New("InsertUserFileSql sql exec sqlite3 database failed")
	}

	_, err = res.RowsAffected()
	if err != nil {
		return errors.New("InsertUserFileSql sql raws affected failed")
	}
	return nil
}


func (c *SqliteContext) getUserfile(username string )([]*util.FileDBInfo, error) {
	sql := "select * from fileinfo where Username= '" + username + "';"
	rows, err := c.db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("getuserlist sql query failed")
	}
	defer rows.Close()
	fmt.Println(sql)

	var files = []*util.FileDBInfo{}
	for rows.Next() {
		dbtag := new(util.FileDBInfo)
		err := rows.Scan(&dbtag.FileHash, &dbtag.Username, &dbtag.FileName, &dbtag.FileSize, &dbtag.FileNumber, &dbtag.FilePolicy, &dbtag.AuthorizedStorage)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("getToken sql scan failed")
		}
		fmt.Println(dbtag.FileHash)
		files = append(files, dbtag)
	}
	return files, nil
}
func (r *SqliteRepository) CallInsertUserFileList(file util.FileDBInfo) (int32, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return 0, errors.New("connectDB failed")
	}
	err2 := db.InsertUserfile(file)
	if err2 != nil {
		log.Println(err2)
		return 0, errors.New("InsertUserfile failed")
	}
	return 1, nil
}
func (r *SqliteRepository) CallGetUserFileList(username string) ([]*util.FileDBInfo, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return nil, errors.New("connectDB failed")
	}
	files,err2 := db.getUserfile(username)
	if err2 != nil {
		log.Println(err2)
		return nil, errors.New("InsertUserfile failed")
	}
	return files, nil
}
