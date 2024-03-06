package main

import (
	"bytes"
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	leveldbPath = "/home/node01/Documents/eth-data/geth/chaindata"

	fmt.Println("------------------ Get data from leveldb-------------------")

	db, err := pebble.Open(leveldbPath, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Pebble open successfully")

	var num uint64
	num = 3000000
	// Get Block Hash Key By Number
	fmt.Println("111")
	blkHashKey := getBlockHeaderHashKey(num)
	fmt.Println("222")
	// Get Block Hash from Key
	blkHash, closer, _ := db.Get(blkHashKey)
	err = closer.Close()
	if err != nil {
		return
	}
	fmt.Println("333")

	fmt.Println("-------", blkHash, "---------")

	headerKey := getBlockHeaderKey(num, blkHash)
	//
	blkHeaderData, closer, _ := db.Get(headerKey) // headerKey是新的key
	//
	err = closer.Close()
	if err != nil {
		return
	}
	_byteData := bytes.NewReader(blkHeaderData)
	blkHeader := new(types.Header)
	_ = rlp.Decode(_byteData, blkHeader)
	//
	fmt.Printf("Block Hash: %x \n", blkHeader.Hash())
	fmt.Printf("Block Coinbase: %x \n", blkHeader.Coinbase)

	_ = db.Close()
}
