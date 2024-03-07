package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	pebblePath := "/home/node01/Documents/eth-data/geth/chaindata/"
	pebbleAncientPath := "/home/node01/Documents/eth-data/geth/chaindata/ancient/chain"

	fmt.Println("------------------ Get data from pebble-------------------")

	config := OpenOptions{Type: dbPebble, Directory: pebblePath, AncientsDirectory: pebbleAncientPath, Namespace: "eth/db/chaindata/", Cache: 2048,
		Handles: 2048, ReadOnly: false, Ephemeral: false}
	db, _ := openKeyValueDatabase(config)
	// Get Block Hash Key By Number
	hash := common.HexToHash("0x305fb171c3f9f626d122da8d7a261bc6f01311eaea02d05e4467bea3a8bd07ae")
	blockNumKey := headerNumberKey(hash)
	blockNum, err := db.Get(blockNumKey)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("blockNum: ", blockNum)
	number := binary.BigEndian.Uint64(blockNum)
	fmt.Println(number)

	hash = common.HexToHash("0x0e066f3c2297a5cb300593052617d1bca5946f0caa0635fdb1b85ac7e5236f34")
	state, err := db.Get(hash.Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Println("state: ", state)

	account := common.HexToAddress("0xb4bfEfC30A60B87380e377F8B96CC3b2E65A8F64")
	accountHash := crypto.Keccak256Hash(account.Bytes())
	valueKey := accountTrieValueKey(accountHash)
	value, err := db.Get(valueKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Value: ", value)

	_byteData := bytes.NewReader(value)
	valuedata := new(types.StateAccount)
	_ = rlp.Decode(_byteData, valuedata)
	fmt.Println(valuedata)

	storageHash := valuedata.Root
	codeHash := common.BytesToHash(valuedata.CodeHash)
	codeKey := accountTrieCodeKey(codeHash)
	code, err := db.Get(codeKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("code: ", code)

	storageKey := accountTrieStorageKey(accountHash, storageHash)
	storage, err := db.Get(storageKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("storage: ", storage)

	//hk := headerKey(hash, number)
	//header, err := db.Get(hk)
	//if err != nil {
	//	panic(err)
	//}
	//_byteData := bytes.NewReader(header)
	//blkHeader := new(types.Header)
	//_ = rlp.Decode(_byteData, blkHeader)
	////
	//fmt.Printf("Block Hash: %x \n", blkHeader.Hash())
	//fmt.Printf("Block Coinbase: %x \n", blkHeader.Coinbase)

	_ = db.Close()
}
