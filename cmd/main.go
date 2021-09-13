package main

import (
	"fmt"

	"github.com/Yemanden/block-chain/pkg/blockchain"
	"github.com/Yemanden/block-chain/pkg/proofofwork"
)

func main() {
	proofOfWordGenerator := proofofwork.New()

	bc := blockchain.New(proofOfWordGenerator)

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.GetBlocks() {
		fmt.Printf(
			"Prev. hash: %x\nData: %s\nHash: %x\n\n",
			block.GetPrevHash(),
			block.GetData(),
			block.GetHash(),
		)
	}
}
