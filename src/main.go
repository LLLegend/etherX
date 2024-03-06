package main

import (
	"bytes"
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/bloom"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"runtime"
)

func main() {
	leveldbPath = "/home/node01/Documents/eth-data/geth/chaindata"

	fmt.Println("------------------ Get data from pebble-------------------")

	maxMemTableSize := (1<<31)<<(^uint(0)>>63) - 1
	cache := 16
	// Two memory tables is configured which is identical to leveldb,
	// including a frozen memory table and another live one.
	memTableLimit := 2
	memTableSize := cache * 1024 * 1024 / 2 / memTableLimit
	if memTableSize >= maxMemTableSize {
		memTableSize = maxMemTableSize - 1
	}
	opt := &pebble.Options{
		// Pebble has a single combined cache area and the write
		// buffers are taken from this too. Assign all available
		// memory allowance for cache.
		Cache:        pebble.NewCache(int64(cache * 1024 * 1024)),
		MaxOpenFiles: 16,

		// The size of memory table(as well as the write buffer).
		// Note, there may have more than two memory tables in the system.
		MemTableSize: uint64(memTableSize),

		// MemTableStopWritesThreshold places a hard limit on the size
		// of the existent MemTables(including the frozen one).
		// Note, this must be the number of tables not the size of all memtables
		// according to https://github.com/cockroachdb/pebble/blob/master/options.go#L738-L742
		// and to https://github.com/cockroachdb/pebble/blob/master/db.go#L1892-L1903.
		MemTableStopWritesThreshold: memTableLimit,

		// The default compaction concurrency(1 thread),
		// Here use all available CPUs for faster compaction.
		MaxConcurrentCompactions: func() int { return runtime.NumCPU() },

		// Per-level options. Options for at least one level must be specified. The
		// options for the last level are used for all subsequent levels.
		Levels: []pebble.LevelOptions{
			{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
			{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
			{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
			{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
			{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
			{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
			{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
		},
		ReadOnly: false,
		//EventListener: &pebble.EventListener{
		//	CompactionBegin: db.onCompactionBegin,
		//	CompactionEnd:   db.onCompactionEnd,
		//	WriteStallBegin: db.onWriteStallBegin,
		//	WriteStallEnd:   db.onWriteStallEnd,
		//},
	}
	// Disable seek compaction explicitly. Check https://github.com/ethereum/go-ethereum/pull/20130
	// for more details.
	opt.Experimental.ReadSamplingMultiplier = -1

	db, err := pebble.Open(leveldbPath, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Pebble open successfully")

	var num uint64
	num = 3000000
	// Get Block Hash Key By Number
	fmt.Println("111")
	blkHashKey := getBlockHeaderHashKey(num)
	fmt.Println("222")
	// Get Block Hash from Key
	blkHash, closer, err := db.Get(blkHashKey)
	if err != nil {
		panic(err)
	}
	err = closer.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("333")

	fmt.Println("-------", blkHash, "---------")

	headerKey := getBlockHeaderKey(num, blkHash)
	//
	blkHeaderData, closer, _ := db.Get(headerKey) // headerKey是新的key
	//
	err = closer.Close()
	if err != nil {
		return
	}
	_byteData := bytes.NewReader(blkHeaderData)
	blkHeader := new(types.Header)
	_ = rlp.Decode(_byteData, blkHeader)
	//
	fmt.Printf("Block Hash: %x \n", blkHeader.Hash())
	fmt.Printf("Block Coinbase: %x \n", blkHeader.Coinbase)

	_ = db.Close()
}
