package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	pebblePath := "/home/node01/Documents/eth-data/geth/chaindata/"
	pebbleAncientPath := "/home/node01/Documents/eth-data/geth/chaindata/ancient/chain"

	fmt.Println("------------------ Get data from pebble-------------------")

	//cache := 2048
	//handles := 2048
	//
	//// The max memtable size is limited by the uint32 offsets stored in
	//// internal/arenaskl.node, DeferredBatchOp, and flushableBatchEntry.
	////
	//// - MaxUint32 on 64-bit platforms;
	//// - MaxInt on 32-bit platforms.
	////
	//// It is used when slices are limited to Uint32 on 64-bit platforms (the
	//// length limit for slices is naturally MaxInt on 32-bit platforms).
	////
	//// Taken from https://github.com/cockroachdb/pebble/blob/master/internal/constants/constants.go
	//maxMemTableSize := (1<<31)<<(^uint(0)>>63) - 1
	//
	//// Two memory tables is configured which is identical to leveldb,
	//// including a frozen memory table and another live one.
	//memTableLimit := 2
	//memTableSize := cache * 1024 * 1024 / 2 / memTableLimit
	//
	//// The memory table size is currently capped at maxMemTableSize-1 due to a
	//// known bug in the pebble where maxMemTableSize is not recognized as a
	//// valid size.
	////
	//// TODO use the maxMemTableSize as the maximum table size once the issue
	//// in pebble is fixed.
	//if memTableSize >= maxMemTableSize {
	//	memTableSize = maxMemTableSize - 1
	//}
	//
	//opt := &pebble.Options{
	//	// Pebble has a single combined cache area and the write
	//	// buffers are taken from this too. Assign all available
	//	// memory allowance for cache.
	//	Cache:        pebble.NewCache(int64(cache * 1024 * 1024)),
	//	MaxOpenFiles: handles,
	//
	//	// The size of memory table(as well as the write buffer).
	//	// Note, there may have more than two memory tables in the system.
	//	MemTableSize: uint64(memTableSize),
	//
	//	// MemTableStopWritesThreshold places a hard limit on the size
	//	// of the existent MemTables(including the frozen one).
	//	// Note, this must be the number of tables not the size of all memtables
	//	// according to https://github.com/cockroachdb/pebble/blob/master/options.go#L738-L742
	//	// and to https://github.com/cockroachdb/pebble/blob/master/db.go#L1892-L1903.
	//	MemTableStopWritesThreshold: memTableLimit,
	//
	//	// The default compaction concurrency(1 thread),
	//	// Here use all available CPUs for faster compaction.
	//	MaxConcurrentCompactions: func() int { return runtime.NumCPU() },
	//
	//	// Per-level options. Options for at least one level must be specified. The
	//	// options for the last level are used for all subsequent levels.
	//	Levels: []pebble.LevelOptions{
	//		{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
	//		{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
	//		{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
	//		{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
	//		{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
	//		{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
	//		{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
	//	},
	//	ReadOnly: false,
	//}
	//opt.EnsureDefaults()
	//// Disable seek compaction explicitly. Check https://github.com/ethereum/go-ethereum/pull/20130
	//// for more details.
	//opt.Experimental.ReadSamplingMultiplier = -1

	//db, err := pebble.Open(leveldbPath, opt)
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//fmt.Println("Pebble open successfully")

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

	hk := headerKey(hash, number)
	header, err := db.Get(hk)
	if err != nil {
		panic(err)
	}
	_byteData := bytes.NewReader(header)
	blkHeader := new(types.Header)
	_ = rlp.Decode(_byteData, blkHeader)
	//
	fmt.Printf("Block Hash: %x \n", blkHeader.Hash())
	fmt.Printf("Block Coinbase: %x \n", blkHeader.Coinbase)

	_ = db.Close()
}
