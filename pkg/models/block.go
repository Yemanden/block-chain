package models

import (
	"bytes"
	"encoding/gob"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func (b *Block) Serialize() (data []byte, err error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err = encoder.Encode(b)

	return result.Bytes(), err
}

func (b *Block) DeserializeBlock(d []byte) (err error) {
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err = decoder.Decode(&b)
	return
}
