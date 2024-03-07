package main

import (
	"encoding/json"
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

type TracerResponse struct {
	JsonRPC float64         `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
}

//{"jsonrpc":"2.0",
//	"id":1,
//	"result":
//		{
//	"from":"0xf9dff387dcb5cc4cca5b91adb07a95f54e9f1bb6",
//	"gas":"0x249f0",
//	"gasUsed":"0xae5b",
//	"to":"0xc66ea802717bfb9833400264dd12c2bceaa34a6d",
//	"input":"0x095ea7b30000000000000000000000001f5006dff7e123d550abc8a4c46792518401fcaf00000000000000000000000000000000000000000000001043561a8829300000","output":"0x0000000000000000000000000000000000000000000000000000000000000000",
//	"calls":[{
//		"from":"0xc66ea802717bfb9833400264dd12c2bceaa34a6d",
//		"gas":"0x18b16","gasUsed":"0x4e1a",
//		"to":"0xbaf42749e027bb38ce4f23ddae8c84da8a15488f",
//		"input":"0xe1f21c67000000000000000000000000f9dff387dcb5cc4cca5b91adb07a95f54e9f1bb60000000000000000000000001f5006dff7e123d550abc8a4c46792518401fcaf00000000000000000000000000000000000000000000001043561a8829300000",
//		"output":"0x0000000000000000000000000000000000000000000000000000000000000000",
//		"calls":[
//			{
//			"from":"0xbaf42749e027bb38ce4f23ddae8c84da8a15488f",
//			"gas":"0x12387",
//			"gasUsed":"0x340",
//			"to":"0x77a79a78c56504c6c1f7499852b6e1918a6d0ab4",
//			"input":"0xb7009613000000000000000000000000c66ea802717bfb9833400264dd12c2bceaa34a6d000000000000000000000000baf42749e027bb38ce4f23ddae8c84da8a15488fe1f21c6700000000000000000000000000000000000000000000000000000000",
//			"output":"0x0000000000000000000000000000000000000000000000000000000000000001",
//			"value":"0x0",
//			"type":"CALL"
//			},
//			{
//			"from":"0xbaf42749e027bb38ce4f23ddae8c84da8a15488f",
//			"gas":"0x11ba9",
//			"gasUsed":"0x2669",
//			"to":"0x96477a1c968a0e64e53b7ed01d0d6e4a311945c2",
//			"input":"0x0c9fcec9000000000000000000000000f9dff387dcb5cc4cca5b91adb07a95f54e9f1bb60000000000000000000000001f5006dff7e123d550abc8a4c46792518401fcaf00000000000000000000000000000000000000000000001043561a8829300000",
//			"calls":[
//				{
//					"from":"0x96477a1c968a0e64e53b7ed01d0d6e4a311945c2",
//					"gas":"0xb54d",
//					"gasUsed":"0x340",
//					"to":"0x77a79a78c56504c6c1f7499852b6e1918a6d0ab4",
//					"input":"0xb7009613000000000000000000000000baf42749e027bb38ce4f23ddae8c84da8a15488f00000000000000000000000096477a1c968a0e64e53b7ed01d0d6e4a311945c20c9fcec900000000000000000000000000000000000000000000000000000000",
//					"output":"0x0000000000000000000000000000000000000000000000000000000000000001",
//					"value":"0x0",
//					"type":"CALL"}
//			],
//			"logs":[
//				{"address":"0x96477a1c968a0e64e53b7ed01d0d6e4a311945c2",
//					"topics":["0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925","0x000000000000000000000000f9dff387dcb5cc4cca5b91adb07a95f54e9f1bb6","0x0000000000000000000000001f5006dff7e123d550abc8a4c46792518401fcaf"],
//					"data":"0x00000000000000000000000000000000000000000000001043561a8829300000",
//					"position":"0x1"}],
//			"value":"0x0","type":"CALL"
//			},
//			{
//			"from":"0xbaf42749e027bb38ce4f23ddae8c84da8a15488f",
//			"gas":"0xf0e6",
//			"gasUsed":"0x129c",
//			"to":"0xc66ea802717bfb9833400264dd12c2bceaa34a6d",
//			"input":"0x5687f2b8000000000000000000000000f9dff387dcb5cc4cca5b91adb07a95f54e9f1bb60000000000000000000000001f5006dff7e123d550abc8a4c46792518401fcaf00000000000000000000000000000000000000000000001043561a8829300000",
//			"calls":[{"from":"0xc66ea802717bfb9833400264dd12c2bceaa34a6d","gas":"0x8a0a","gasUsed":"0x340","to":"0x77a79a78c56504c6c1f7499852b6e1918a6d0ab4","input":"0xb7009613000000000000000000000000baf42749e027bb38ce4f23ddae8c84da8a15488f000000000000000000000000c66ea802717bfb9833400264dd12c2bceaa34a6d5687f2b800000000000000000000000000000000000000000000000000000000","output":"0x0000000000000000000000000000000000000000000000000000000000000001","value":"0x0","type":"CALL"}],"logs":[{"address":"0xc66ea802717bfb9833400264dd12c2bceaa34a6d","topics":["0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925","0x000000000000000000000000f9dff387dcb5cc4cca5b91adb07a95f54e9f1bb6","0x0000000000000000000000001f5006dff7e123d550abc8a4c46792518401fcaf"],"data":"0x00000000000000000000000000000000000000000000001043561a8829300000","position":"0x1"}],"value":"0x0","type":"CALL"}],"value":"0x0","type":"CALL"}],"value":"0x0","type":"CALL"}}
