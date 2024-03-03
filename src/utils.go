package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func parseTxData(tx *types.Transaction, sender common.Address, status uint64) *TransactionDetail {
	var txd TransactionDetail
	txd.From = sender.String()
	txd.To = tx.To().String()
	txd.Value = tx.Value().Int64()
	txd.FunctionSelector = ""
	txd.LogTopics = ""
	txd.LogData = ""
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

func parseTxTraceData(tx *types.Transaction, data interface{}, sender common.Address) []*TransactionDetail {
	var txds []*TransactionDetail

	var tracerResponse TracerResponse
	tracerResponse = data.(TracerResponse)
	fmt.Println(tracerResponse)
	return txds
}
