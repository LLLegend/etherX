package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	ipcPath = "/home/node01/Documents/eth-data/geth.ipc"
	leveldbPath = "/home/node01/Documents/eth-data/geth/chaindata/"
	//
	//client, err := ethclient.Dial(ipcPath)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Connected")
	//blockNumber := big.NewInt(3000000)
	//start := time.Now()
	//header, err := client.HeaderByNumber(context.Background(), blockNumber)
	//blockHash := header.Hash()
	//numTx, err := client.TransactionCount(context.Background(), blockHash)
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
	//fmt.Println("tx Receipts: ", txReceipt)
	//
	//fmt.Println("Using ", time.Since(start))
	////
	//api := tracers.API{}
	//traceConfig := &tracers.TraceConfig{Config: &logger.Config{}}
	//res, err := api.TraceTransaction(context.Background(), tx.Hash(), traceConfig)
	//fmt.Println("trace: ", res)

	fmt.Println("------------------ Get data from leveldb-------------------")

	db, err := leveldb.OpenFile(leveldbPath, nil)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("LevelDB open successfully")

	//var num uint64
	//num = 3000000
	// Get Block Hash Key By Number
	blkHashKey := getBlockHeaderHashKey(30000)
	// Get Block Hash from Key
	blkHash, _ := db.Get(blkHashKey, nil)

	fmt.Println("-------", blkHash, "---------")

	//headerKey := getBlockHeaderKey(num, blkHash)
	//
	//blkHeaderData, _ := db.Get(headerKey, nil) // headerKey是新的key
	//
	//_byteData := bytes.NewReader(blkHeaderData)
	//blkHeader := new(types.Header)
	//_ = rlp.Decode(_byteData, blkHeader)
	//
	//fmt.Printf("Block Hash: %x \n", blkHeader.Hash())
	//fmt.Printf("Block Coinbase: %x \n", blkHeader.Coinbase)

	_ = db.Close()
}
