package main

import (
	"time"
)

type Block struct {
	BlockNumber     int64
	BlockHash       string
	ParentHash      string
	Coinbase        string
	Timestamp       time.Time
	GasUsed         uint64
	GasLimit        uint64
	BlockSize       int64
	Difficulty      uint64
	Extra           string
	ExternalTxCount int64
	InternalTxCount int64
}

type TransactionBackground struct {
	BlockNumber     int64
	TxHash          string
	PositionInBlock int
	Timestamp       time.Time
	GasUsed         uint64
	GasLimit        uint64
	Nonce           uint64
	ExecutionStatus string
	CallTree        string
}

type TransactionDetail struct {
	From             string
	To               string
	Value            int64
	FunctionSelector string
	LogTopics        string
	LogData          string
	CallLevel        uint64
	CallPosition     uint64
	TxType           string
	Status           string
}

type Transaction struct {
	TransactionBackground *TransactionBackground
	TransactionDetail     *TransactionDetail
}

type TracerConfig struct {
	Tracer       string             `json:"tracer"`
	TracerConfig OnlyTopCallWithLog `json:"tracerConfig"`
}

type OnlyTopCallWithLog struct {
	OnlyTopCall bool `json:"onlyTopCall"`
	WithLog     bool `json:"withLog"`
}

type TracerResponse struct {
	From    string           `json:"from"`
	GasUsed string           `json:"gasUsed"`
	To      string           `json:"to"`
	Input   string           `json:"input"`
	Output  string           `json:"output"`
	Calls   []TracerResponse `json:"calls"`
	Type    string           `json:"type"`
}
