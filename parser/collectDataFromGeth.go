package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/eth/tracers/logger"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
	"time"
)

func main() {
	ipcPath = "/home/node01/Documents/eth-data/geth.ipc"
	leveldbPath = "/home/node01/Documents/eth-data/geth/chaindata/"
	//
	client, err := ethclient.Dial(ipcPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")
	blockNumber := big.NewInt(4000000)
	start := time.Now()
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	blockHash := header.Hash()
	numTx, err := client.TransactionCount(context.Background(), blockHash)
	fmt.Println("Block Number: ", header.Number)
	fmt.Println("Block Hash: ", header.Hash())
	fmt.Println("Block Coinbase: ", header.Coinbase)
	fmt.Println("Block gasUsed: ", header.GasUsed)
	fmt.Println("Block Time: ", header.Time)
	fmt.Println("Tx Hash: ", header.TxHash)
	fmt.Println("size: ", header.Size())
	fmt.Println("numTx: ", numTx)

	bHash := rpc.BlockNumberOrHash{BlockHash: &blockHash}
	receipts, err := client.BlockReceipts(context.Background(), bHash)
	fmt.Println("Block Receipts: ", receipts)

	tx, err := client.TransactionInBlock(context.Background(), blockHash, 0)
	txReceipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	fmt.Println("tx Receipts: ", txReceipt)

	fmt.Println("Using ", time.Since(start))
	//
	api := tracers.API{}
	traceConfig := &tracers.TraceConfig{Config: &logger.Config{}}
	res, err := api.TraceTransaction(context.Background(), tx.Hash(), traceConfig)
	fmt.Println("trace: ", res)
}
