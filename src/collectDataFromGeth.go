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

type Block struct {
	blockNumber     int64
	blockHash       string
	parentHash      string
	coinbase        string
	timestamp       time.Time
	gasUsed         uint64
	gasLimit        uint64
	blockSize       int64
	difficulty      uint64
	extra           string
	externalTxCount int64
	internalTxCount int64
}

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
	endBlockNumber := int64(50)

	start := time.Now()
	genesis, _ := client.HeaderByNumber(context.Background(), big.NewInt(0))
	parentHash := genesis.Hash().String()
	for i := blockNumber; i <= endBlockNumber; i++ {
		var block Block
		block = *new(Block)

		header, err := client.HeaderByNumber(context.Background(), big.NewInt(i))
		if err != nil {

		}
		blockHash := header.Hash()

		numTx, err := client.TransactionCount(context.Background(), blockHash)
		if numTx > 0 {
			fmt.Println(i, numTx)
		}
		//bloomByte, _ := header.Bloom.MarshalText()

		block.blockNumber = i
		block.blockHash = blockHash.String()
		block.parentHash = parentHash
		block.coinbase = header.Coinbase.String()
		block.timestamp = time.Unix(int64(header.Time), 0)
		block.gasUsed = header.GasUsed
		block.gasLimit = header.GasLimit
		//block.logsBloom = string(bloomByte)
		block.blockSize = int64(header.Size())
		block.difficulty = header.Difficulty.Uint64()
		block.extra = string(header.Extra)
		block.externalTxCount = int64(numTx)
		block.internalTxCount = 0

		err = insertBlocks(db, block)
		if err != nil {
			panic(err)
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
