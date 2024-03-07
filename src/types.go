package main

import (
	"time"
)

type Block struct {
	BlockNumber     uint64
	BlockHash       string
	ParentHash      string
	Coinbase        string
	Timestamp       time.Time
	GasUsed         uint64
	GasLimit        uint64
	BlockSize       uint64
	Difficulty      uint64
	Extra           string
	ExternalTxCount uint64
	InternalTxCount uint64
}

type TransactionBackground struct {
	BlockNumber     uint64
	TxHash          string
	PositionInBlock uint64
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
	Value            uint64
	FunctionSelector string
	LogTopics        string
	LogData          string
	CallLevel        uint64
	CallPosition     uint64
	TxType           string
	Status           string
	Input            string
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

type LogBody struct {
	Address  string   `json:"address"`
	Topics   []string `json:"topics"`
	Data     string   `json:"data"`
	Position string   `json:"position"`
}

type TracerBody struct {
	From    string       `json:"from"`
	GasUsed string       `json:"gasUsed"`
	Gas     string       `json:"gas"`
	To      string       `json:"to"`
	Input   string       `json:"input"`
	Output  string       `json:"output,omitempty"`
	Calls   []TracerBody `json:"calls,omitempty"`
	Type    string       `json:"type"`
	Logs    []LogBody    `json:"logs,omitempty"`
	Value   string       `json:"value"`
}
