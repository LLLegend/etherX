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

	blockNumber := int64(1000000)
	endBlockNumber := int64(1000100)

	start := time.Now()
	genesis, _ := client.HeaderByNumber(context.Background(), big.NewInt(0))
	parentHash := genesis.Hash().String()
	config := OnlyTopCallWithLog{OnlyTopCall: "False", WithLog: "True"}
	tracerConfig := TraceConfig{Tracer: "callTracer", TracerConfig: config}
	for i := blockNumber; i <= endBlockNumber; i++ {
		var block Block
		block = *new(Block)

		header, err := client.HeaderByNumber(context.Background(), big.NewInt(i))
		if err != nil {
			panic(err)
		}

		blockHash := header.Hash()
		numTx, err := client.TransactionCount(context.Background(), blockHash)

		// write block data to mysql
		go func() {
			block.BlockNumber = i
			block.BlockHash = blockHash.String()
			block.ParentHash = parentHash
			block.Coinbase = header.Coinbase.String()
			block.Timestamp = time.Unix(int64(header.Time), 0)
			block.GasUsed = header.GasUsed
			block.GasLimit = header.GasLimit
			block.BlockSize = int64(header.Size())
			block.Difficulty = header.Difficulty.Uint64()
			block.Extra = string(header.Extra)
			block.ExternalTxCount = int64(numTx)
			block.InternalTxCount = 0
			//err = insertBlocks(db, block)
			//if err != nil {
			//	panic(err)
			//}
		}()

		for j := 0; j < int(numTx); j++ {
			var txb *TransactionBackground
			txb = new(TransactionBackground)

			tx, err := client.TransactionInBlock(context.Background(), blockHash, 0)
			if err != nil {
				panic(err)
			}

			var resp string

			req := TraceTransactionRequest{tx.Hash().String(), tracerConfig}
			err = client.Client().Call(&resp, "debug_traceTransaction", req)

			fmt.Println(resp)
			fmt.Println(req)
			fmt.Println(1)
			txb.BlockNumber = i
			txb.TxHash = tx.Hash().String()
			txb.PositionInBlock = j
			txb.GasLimit = tx.Gas()
			txb.Timestamp = tx.Time()
			txb.Nonce = tx.Nonce()

		}

		parentHash = block.BlockHash
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
