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
	ipcPath := "/home/node01/Documents/eth-data/geth.ipc"

	client, err := ethclient.Dial(ipcPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")
	blockNumber := big.NewInt(3000000)
	start := time.Now()
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	fmt.Println("Block Number: ", header.Number)
	fmt.Println("Block Hash: ", header.Hash())
	fmt.Println("Block Coinbase: ", header.Coinbase)
	fmt.Println("Block gasUsed: ", header.GasUsed)
	fmt.Println("Block Time: ", header.Time)
	fmt.Println("Tx Hash: ", header.TxHash)
	fmt.Println("size: ", header.Size())
	fmt.Println("Using ", time.Since(start))

}
