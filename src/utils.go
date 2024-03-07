package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
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

func parseTxTraceData(tx *types.Transaction, data map[string]interface{}, sender common.Address) []*TransactionDetail {
	var txds []*TransactionDetail
	var resp TracerBody

	jsonString, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonString))
	// 将 JSON 字符串转换为 string 类型

	err = json.Unmarshal(jsonString, &resp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
	return txds
}
