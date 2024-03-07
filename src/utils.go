package main

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

func parseTxData(tx *types.Transaction, sender common.Address, status uint64) *TransactionDetail {
	var txd TransactionDetail
	txd.From = sender.String()
	txd.To = tx.To().String()
	txd.Value = tx.Value().Uint64()
	txd.FunctionSelector = ""
	txd.LogTopics = ""
	txd.LogData = ""
	txd.Input = ""
	txd.CallLevel = 1
	txd.CallPosition = 1
	txd.TxType = "Transfer"
	if status == 1 {
		txd.Status = "Failed"
	} else {
		txd.Status = "Succeed"
	}
	return &txd
}

func parseTxTraceData(tx *types.Transaction, data map[string]interface{}) []*TransactionDetail {
	var callLevel *uint64
	var callPosition *uint64

	*callLevel = 1
	*callPosition = 1

	var txds []*TransactionDetail
	var resp TracerBody

	byteData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteData, &resp)
	if err != nil {
		log.Println(err)
	}

	var txd TransactionDetail
	txd.From = resp.From
	txd.To = resp.To
	txd.Value = hexStringToUint64(resp.Value)
	txd.FunctionSelector = ""

	// TO DO parse Log
	txd.LogTopics = ""
	txd.LogData = ""

	txd.Input = resp.Input
	txd.CallLevel = *callLevel
	txd.CallPosition = *callPosition
	txd.TxType = resp.Type
	txds = append(txds, &txd)

	for i := 0; i < len(resp.Calls); i++ {

	}

	return txds
}

func parseTraceData(data TracerBody, level *uint64) {

}

func hexStringToUint64(data string) uint64 {
	return 0
}

func findFunctionSignature() {

}
