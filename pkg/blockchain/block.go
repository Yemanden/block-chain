package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block interface {
	SetHash()
	GetPrevHash() []byte
	GetData() []byte
	GetHash() []byte
}

type block struct {
	timestamp     int64
	data          []byte
	prevBlockHash []byte
	hash          []byte
}

func (b *block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.timestamp, 10))
	headers := bytes.Join([][]byte{b.prevBlockHash, b.data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.hash = hash[:]
}

func (b *block) GetPrevHash() []byte {
	return b.prevBlockHash
}

func (b *block) GetData() []byte {
	return b.data
}

func (b *block) GetHash() []byte {
	return b.hash
}

func newBlock(data string, prevBlockHash []byte) Block {
	b := &block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	b.SetHash()
	return b
}

func newGenesisBlock() Block {
	return newBlock("Genesis Block", []byte{})
}
