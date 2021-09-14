package main

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"

	"github.com/Yemanden/block-chain/pkg/blockchain"
	"github.com/Yemanden/block-chain/pkg/cli"
	"github.com/Yemanden/block-chain/pkg/proofofwork"
)

const dbFile = "blockchain.db"

func main() {
	proofOfWordGenerator := proofofwork.New()
	db, err := bolt.Open(dbFile, os.ModePerm, nil)
	if err != nil {
		fmt.Println("failed connect to DB: " + err.Error())
		os.Exit(1)
	}
	defer db.Close()

	bc, err := blockchain.New(proofOfWordGenerator, db)
	if err != nil {
		fmt.Println("failed create of blockchain: " + err.Error())
		os.Exit(2)
	}

	cli := cli.New(bc)
	err = cli.Run()
	if err != nil {
		fmt.Println("cli error: " + err.Error())
		os.Exit(3)
	}
}
