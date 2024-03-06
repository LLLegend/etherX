package main

import (
	"encoding/binary"
)

func encodeBlockNumber(number uint64) []byte {
	enc := make([]byte, 8)
	binary.BigEndian.PutUint64(enc, number)
	return enc
}

// headerHashKey = headerPrefix + num (uint64 big endian) + headerHashSuffix
func getBlockHeaderHashKey(num uint64) []byte {
	return append(append([]byte("h"), encodeBlockNumber(num)...), []byte("n")...)
}

func getBlockHeaderKey(number uint64, hash []byte) []byte {
	return append(append(append([]byte("h"), encodeBlockNumber(number)...), hash...))
}
