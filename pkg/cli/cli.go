package cli

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/Yemanden/block-chain/pkg/blockchain"
	"github.com/Yemanden/block-chain/pkg/proofofwork"
)

type CLI interface {
	Run() (err error)
}

type cli struct {
	bc blockchain.BlockChain
}

func (c *cli) Run() (err error) {
	c.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err = addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			return
		}
	case "printchain":
		err = printChainCmd.Parse(os.Args[2:])
		if err != nil {
			return
		}
	default:
		c.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		c.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		err = c.printChain()
	}
	return
}

func (c *cli) addBlock(data string) {
	err := c.bc.AddBlock(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success!")
}

func (c *cli) printChain() error {
	bci := c.bc.Iterator()

	for {
		block, err := bci.Next()
		if err != nil {
			return err
		}

		pow := proofofwork.NewProofOfWork(block)
		fmt.Printf(
			"Prev. hash: %x\nData: %s\nHash: %x\nPoW: %s\n\n",
			block.PrevBlockHash,
			block.Data,
			block.Hash,
			strconv.FormatBool(pow.Validate()))

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return nil
}

func (c *cli) validateArgs() {
	if len(os.Args) < 2 {
		c.printUsage()
		os.Exit(1)
	}
}

func (c *cli) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

func New(blockChain blockchain.BlockChain) CLI {
	return &cli{bc: blockChain}
}
