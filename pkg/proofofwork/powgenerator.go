package proofofwork

import "github.com/Yemanden/block-chain/pkg/models"

type PoWGenerator interface {
	Generate(block models.Block) (int, []byte)
}

type powGenerator struct{}

func (p *powGenerator) Generate(block models.Block) (int, []byte) {
	return NewProofOfWork(block).Run()
}

func New() PoWGenerator {
	return &powGenerator{}
}
