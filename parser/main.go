package main

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	leveldbPath = "/home/node01/Documents/eth-data/geth/chaindata/"

	fmt.Println("------------------ Get data from leveldb-------------------")

	db, err := leveldb.OpenFile(leveldbPath, nil)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("LevelDB open successfully")

	var num uint64
	num = 3000000
	// Get Block Hash Key By Number
	fmt.Println("111")
	blkHashKey := getBlockHeaderHashKey(num)
	fmt.Println("222")
	// Get Block Hash from Key
	blkHash, _ := db.Get(blkHashKey, nil)
	fmt.Println("333")

	fmt.Println("-------", blkHash, "---------")

	headerKey := getBlockHeaderKey(num, blkHash)
	//
	blkHeaderData, _ := db.Get(headerKey, nil) // headerKey是新的key
	//
	_byteData := bytes.NewReader(blkHeaderData)
	blkHeader := new(types.Header)
	_ = rlp.Decode(_byteData, blkHeader)
	//
	fmt.Printf("Block Hash: %x \n", blkHeader.Hash())
	fmt.Printf("Block Coinbase: %x \n", blkHeader.Coinbase)

	_ = db.Close()
}
