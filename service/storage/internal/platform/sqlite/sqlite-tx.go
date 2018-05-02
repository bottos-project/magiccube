/*The functions of the Sqlite database that provided here are not ready yet，they should never be exposed to users。They are 
in the Bottos's service layer,which delivering database service pluggable to provide users with queries.And we plan to support
it in a future point release.At present, we only support mongodb to provide users with queries.*/
package sqlite

import (
	"errors"
	"fmt"
	"log"

	"github.com/code/bottos/service/storage/util"
)

func (c *SqliteContext) createTx() {
	sqlStmt := `
		create table txinfo(TransactionID VARCHAR(64) PRIMARY KEY,
		Price INTEGER,
		Type VARCHAR(64),
		From VARCHAR(24),
		To VARCHAR(24),
		Date VARCHAR(24),
		BlockNum INTEGER );
		`
	_, err := c.db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *SqliteContext) insertTxInfo(dbtag util.TxDBInfo) error {

	if !c.IsTableExist("txinfo") {
		c.createTx()
	}
	defer c.db.Close()

	stmt, err := c.db.Prepare(util.InserTxInfoSql)
	if err != nil {
		log.Fatal(err)
		return errors.New("InserTxInfoSql sql insert sqlite3 database failed")
	}

	res, err := stmt.Exec(dbtag.TransactionID, dbtag.Price, dbtag.Type, dbtag.From, dbtag.To, dbtag.Date, dbtag.BlockId)
	if err != nil {
		log.Fatal(err)
		return errors.New("InserTxInfoSql sql exec sqlite3 database failed")
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return errors.New("InserTxInfoSql sql raws affected failed")
	}
	return nil

}

// Read
func (c *SqliteContext) readOneTx(tx string) (*util.TxDBInfo, error) {
	rows, err := c.db.Query("SELECT * FROM txinfo where TransactionID=" + tx)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("readOneTx sql query failed")
	}
	defer rows.Close()

	for rows.Next() {
		dbtag := new(util.TxDBInfo)
		err := rows.Scan(&dbtag.TransactionID, &dbtag.Price, &dbtag.Type, &dbtag.From, &dbtag.To, &dbtag.Date, &dbtag.BlockId)
		if err != nil {
			log.Fatal(err)
			return nil, errors.New("readOneTx sql scan failed")
		}
		return dbtag, nil
	}
	return nil, nil
}
func (r *SqliteRepository) CallInsertTxInfo(value util.TxDBInfo) error {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
		return errors.New("ConnectDB failed")
	}
	err = db.insertTxInfo(value)
	if err != nil {
		log.Fatal(err)
		return errors.New("insertTxInfo failed")
	}
	return nil
}

func (r *SqliteRepository) CallGetTx(txid string) (*util.TxDBInfo, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("ConnectDB failed")
	}
	res, err2 := db.readOneTx(txid)
	if err2 != nil {
		log.Fatal(err)
		return nil, errors.New("readOneTx failed")
	}
	return res, nil
}

func (c *SqliteContext) createSync() {
	sqlSync := `
		create table sync(SyncedNumber INTEGER PRIMARY KEY);
		`
	_, err := c.db.Exec(sqlSync)
	if err != nil {
		log.Println(err)
	}
}
func (c *SqliteContext) syncBlockNum() (uint64, error) {
	var num uint64
	if !c.IsTableExist("sync") {
		fmt.Println("create sync")
		c.createSync()
	}
	rows, err := c.db.Query("select count(*) from sync ")
	if err != nil {
		log.Fatal(err)
		fmt.Println("query")
		return 0, errors.New("readOne sql query failed")
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&num)
		if err != nil {
			log.Println(err)
			return 0, errors.New("readOne sql scan failed")
		}
		return num, nil
	}
	return num, nil
}
func (r *SqliteRepository) CallGetSyncBlockCount() (uint64, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
		return 0, errors.New("connectDB failed")
	}

	res, err2 := db.syncBlockNum()
	if err2 != nil {
		log.Fatal(err)
		return 0, errors.New("syncBlockNum failed")
	}
	return res, nil
}
