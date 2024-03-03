package main

const createBlock string = ""

const insertBlock string = "INSERT INTO blocks (block_number, block_hash, parent_hash, coinbase, timestamp, gas_used, gas_limit, block_size, difficulty, extra, external_tx_count, internal_tx_count) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
