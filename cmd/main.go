package main

import (
	"fmt"

	"github.com/Yemanden/block-chain/pkg/blockchain"
)

func main() {
	bc := blockchain.New()

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
