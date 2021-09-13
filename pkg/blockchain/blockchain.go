package blockchain

type BlockChain interface {
	AddBlock(data string)
	GetBlocks() []Block
}

type blockChain struct {
	blocks []Block
}

func (b *blockChain) AddBlock(data string) {
	prevBlock := b.blocks[len(b.blocks)-1]
	nb := newBlock(data, prevBlock.GetHash())
	b.blocks = append(b.blocks, nb)
}

func (b *blockChain) GetBlocks() []Block {
	return b.blocks
}

func New() BlockChain {
	return &blockChain{[]Block{newGenesisBlock()}}
}
