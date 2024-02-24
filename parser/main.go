package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"math/big"
	"time"
)

func main() {
	ipcPath = "/home/node01/Documents/eth-data/geth.ipc"
	leveldbPath = "/home/node01/Documents/eth-data/geth/chaindata/"

	client, err := ethclient.Dial(ipcPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")
	blockNumber := big.NewInt(3000000)
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
	//api := tracers.API{}
	//traceConfig := &tracers.TraceConfig{Config: &logger.Config{}}
	//res, err := api.TraceTransaction(context.Background(), tx.Hash(), traceConfig)
	//fmt.Println("trace: ", res)

	fmt.Println("------------------ Get data from leveldb-------------------")

	db, err := leveldb.OpenFile(leveldbPath, nil)
	headerPrefix := []byte("h") // headerPrefix + num (uint64 big endian) + hash -> header
	numSuffix := []byte("n")    // headerPrefix + num (uint64 big endian) + numSuffix -> hash
	blkNum := make([]byte, 8)
	binary.BigEndian.PutUint64(blkNum, uint64(3000000)) // 把num变为 uint64 big endian类型的数据

	hashKey := append(headerPrefix, blkNum...) // headerPrefix + blkNum
	hashKey = append(hashKey, numSuffix...)    // blkNum + headerPrefix + numSuffix
	blkHash, _ := db.Get(hashKey, nil)

	//
	fmt.Println("-------", blkHash, "---------")

	headerKey := append(headerPrefix, blkNum...) // headerPrefix + blkNum
	headerKey = append(headerKey, blkHash...)    // headerPrefix + blkNum + blkHash

	blkHeaderData, _ := db.Get(headerKey, nil) // headerKey是新的key

	_byteData := bytes.NewReader(blkHeaderData)
	blkHeader := new(types.Header)
	_ = rlp.Decode(_byteData, blkHeader)

	fmt.Printf("Block Hash: %x \n", blkHeader.Hash())
	fmt.Printf("Block Coinbase: %x \n", blkHeader.Coinbase)
	_ = db.Close()
}
