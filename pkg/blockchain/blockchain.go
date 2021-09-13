package blockchain

import (
	"time"

	"github.com/Yemanden/block-chain/pkg/models"
)

type powGenerator interface {
	Generate(block models.Block) (int, []byte)
}

type BlockChain interface {
	AddBlock(data string)
	GetBlocks() []models.Block
}

type blockChain struct {
	blocks       []models.Block
	powGenerator powGenerator
}

func (b *blockChain) AddBlock(data string) {
	prevBlock := b.blocks[len(b.blocks)-1]
	nb := b.newBlock(data, prevBlock.GetHash())
	b.blocks = append(b.blocks, nb)
}

func (b *blockChain) GetBlocks() []models.Block {
	return b.blocks
}

func (b *blockChain) newBlock(data string, prevBlockHash []byte) models.Block {
	newBlock := &block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	nonce, hash := b.powGenerator.Generate(newBlock)

	newBlock.hash = hash[:]
	newBlock.nonce = nonce
	return newBlock
}

func (b *blockChain) newGenesisBlock() models.Block {
	return b.newBlock("Genesis Block", []byte{})
}

func New(powGenerator powGenerator) BlockChain {
	bc := &blockChain{
		powGenerator: powGenerator,
	}
	bc.blocks = []models.Block{bc.newGenesisBlock()}
	return bc
}
