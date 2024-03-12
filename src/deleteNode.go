package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

func main() {

	pebblePath := "/home/node01/Documents/eth-data-2/geth/chaindata/"
	pebbleAncientPath := "/home/node01/Documents/eth-data-2/geth/chaindata/ancient/chain"

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/d4cee2a05a2d453a8f83b7b3f9f89b75")
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}

	config := OpenOptions{Type: dbPebble, Directory: pebblePath, AncientsDirectory: pebbleAncientPath, Namespace: "eth/db/chaindata/", Cache: 2048,
		Handles: 2048, ReadOnly: false, Ephemeral: false}
	db, _ := openKeyValueDatabase(config)

	// 4100000 - 4101200 10 -100
	blockNumber := uint64(10)
	endBlockNumber := uint64(100000)

	total := int64(0)
	for i := blockNumber; i <= endBlockNumber; i++ {
		header, _ := client.HeaderByNumber(context.Background(), big.NewInt(int64(i)))
		rootHash := header.Root

		start := time.Now()

		err := db.Delete(rootHash.Bytes())
		if err != nil {
			fmt.Println(i, " Block delete error")
			panic(err)
		}
		since := time.Since(start)
		fmt.Println("Delete Block ", i, " Using ", since)

		total += since.Nanoseconds()
	}
	fmt.Println("Average delete ns: ", float64(total)/100.0)

}
