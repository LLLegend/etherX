package main

import (
	"encoding/binary"
	"fmt"
)

func encodeBlockNumber(number uint64) []byte {
	fmt.Println("111-111-111")
	enc := make([]byte, 8)
	binary.BigEndian.PutUint64(enc, number)
	fmt.Println("111-111-222")
	return enc
}

// headerHashKey = headerPrefix + num (uint64 big endian) + headerHashSuffix
func getBlockHeaderHashKey(num uint64) []byte {
	fmt.Println("111-111")
	return append(append([]byte("h"), encodeBlockNumber(num)...), []byte("n")...)
}

func getBlockHeaderKey(number uint64, hash []byte) []byte {
	return append(append(append([]byte("h"), encodeBlockNumber(number)...), hash...))
}
