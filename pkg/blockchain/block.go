package blockchain

type block struct {
	timestamp     int64
	data          []byte
	prevBlockHash []byte
	hash          []byte
	nonce         int
}

func (b *block) GetTimestamp() int64 {
	return b.timestamp
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
