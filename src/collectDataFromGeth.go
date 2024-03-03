package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ipcPath = "/home/node01/Documents/eth-data/geth.ipc"
	leveldbPath = "/home/node01/Documents/eth-data/geth/chaindata/"
	//
	client, err := ethclient.Dial(ipcPath)
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Geth")

	db, err := initMysql("root", "", "127.0.0.1", 3306, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connect to Mysql")
	defer db.Close()

	showTables(db)
	fmt.Println()

	blockNumber := int64(1)
	endBlockNumber := int64(5)

	start := time.Now()
	for i := blockNumber; i <= endBlockNumber; i++ {
		header, err := client.HeaderByNumber(context.Background(), big.NewInt(i))
		if err != nil {

		}
		blockHash := header.Hash()
		fmt.Println(blockHash.String())
		numTx, err := client.TransactionCount(context.Background(), blockHash)
		if numTx > 0 {
			fmt.Println(i, numTx)
		}

	}

	//client.Client().Call()

	// header, err := client.HeaderByNumber(context.Background(), big.NewInt(4000000))

	//fmt.Println("Block Number: ", header.Number)
	//fmt.Println("Block Hash: ", header.Hash())
	//fmt.Println("Block Coinbase: ", header.Coinbase)
	//fmt.Println("Block gasUsed: ", header.GasUsed)
	//fmt.Println("Block Time: ", header.Time)
	//fmt.Println("Tx Hash: ", header.TxHash)
	//fmt.Println("size: ", header.Size())
	//fmt.Println("numTx: ", numTx)
	//
	//bHash := rpc.BlockNumberOrHash{BlockHash: &blockHash}
	//receipts, err := client.BlockReceipts(context.Background(), bHash)
	//fmt.Println("Block Receipts: ", receipts)
	//
	//tx, err := client.TransactionInBlock(context.Background(), blockHash, 0)
	//txReceipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	//fmt.Println("tx to: ", tx.To())
	//fmt.Println("tx input data: ", tx.Data())
	//fmt.Println("tx Receipts BlockHash: ", txReceipt.BlockHash)
	fmt.Println("Using ", time.Since(start))
	//

}
