package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func initMysql(username string, password string, host string, port int, DBName string) (*sql.DB, error) {
	var db *sql.DB
	dsn := "root:19990902Aa@@tcp(127.0.0.1:3306)/ETHERX"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func showTables(db *sql.DB) {
	res, _ := db.Query("SHOW TABLES")

	var table string
	for res.Next() {
		err := res.Scan(&table)
		if err != nil {
			return
		}
		fmt.Println(table)
	}
}

func insertBlocks(db *sql.DB, block Block) error {
	_, err := db.Exec(
		insertBlock,
		block.blockNumber,
		block.blockHash,
		block.parentHash,
		block.coinbase,
		block.timestamp,
		block.gasUsed,
		block.gasLimit,
		block.logsBloom,
		block.blockSize,
		block.difficulty,
		block.extra,
		block.externalTxCount,
		block.internalTxCount,
	)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return nil
}
