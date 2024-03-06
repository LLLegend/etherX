package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
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
	rpcClient, err := rpc.DialHTTP("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	defer rpcClient.Close()

	fmt.Println("Connected to Geth")

	db, err := initMysql("root", "", "127.0.0.1", 3306, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connect to Mysql")
	defer db.Close()

	blockNumber := int64(1)
	endBlockNumber := int64(1500000)

	start := time.Now()
	parentBlock, _ := client.HeaderByNumber(context.Background(), big.NewInt(blockNumber-1))
	parentHash := parentBlock.Hash().String()
	// onlyTopCallWithLog := OnlyTopCallWithLog{OnlyTopCall: false, WithLog: true}
	// tracerConfig := TracerConfig{Tracer: "callTracer", TracerConfig: onlyTopCallWithLog}

	txnum := 0
	numi := 0
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
		block.BlockNumber = i
		block.BlockHash = blockHash.String()
		block.ParentHash = parentHash
		block.Coinbase = header.Coinbase.String()
		block.Timestamp = time.Unix(int64(header.Time), 0)
		block.GasUsed = header.GasUsed
		block.GasLimit = header.GasLimit
		block.BlockSize = int64(header.Size())
		block.Difficulty = header.Difficulty.Uint64()
		block.Extra = hex.EncodeToString(header.Extra)
		block.ExternalTxCount = int64(numTx)
		block.InternalTxCount = 0
		err = insertBlocks(db, block)
		if err != nil {
			panic(err)
		}
		// Find Tx
		//for j := 0; j < int(numTx); j++ {
		//	var txb *TransactionBackground
		//	txb = new(TransactionBackground)
		//
		//	//var txds []*TransactionDetail
		//
		//	// Get Tx
		//	tx, err := client.TransactionInBlock(context.Background(), blockHash, uint(j))
		//	if err != nil {
		//		panic(err)
		//	}
		//	//// Get Sender in the Tx
		//	//sender, err := client.TransactionSender(context.Background(), tx, blockHash, uint(j))
		//	//if err != nil {
		//	//	panic(err)
		//	//}
		//	// Tx that doesn't have internal Tx
		//	if len(tx.Data()) == 0 {
		//		//txReceipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		//		//if err != nil {
		//		//	panic(err)
		//		//}
		//		// txd := parseTxData(tx, sender, txReceipt.Status)
		//		// txds = append(txds, txd)
		//	} else {
		//		var resp interface{}
		//		if err := rpcClient.Call(&resp, "debug_traceTransaction", tx.Hash().String(), tracerConfig); err != nil {
		//			log.Fatal(err)
		//		}
		//		// txd := parseTxTraceData(tx, resp, sender)
		//		// txds = append(txds, txd...)
		//		numi += 1
		//
		//	}
		//	txb.BlockNumber = i
		//	txb.TxHash = tx.Hash().String()
		//	txb.PositionInBlock = j
		//	txb.GasLimit = tx.Gas()
		//	txb.Timestamp = tx.Time()
		//	txb.Nonce = tx.Nonce()
		//
		//	txnum += 1
		//
		//}

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
	fmt.Println("Num Tx contains internal Tx", numi)
	fmt.Println("Total Tx: ", txnum)
	//

}
