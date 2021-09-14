package blockchain

import (
	"time"

	"github.com/boltdb/bolt"

	"github.com/Yemanden/block-chain/pkg/models"
)

const blocksBucket = "blocks"

type powGenerator interface {
	Generate(block models.Block) (int, []byte)
}

type BlockChain interface {
	AddBlock(data string) (err error)
	Iterator() Iterator
}

type blockChain struct {
	powGenerator powGenerator

	tip []byte
	db  *bolt.DB
}

func (b *blockChain) AddBlock(data string) (err error) {
	var lastHash []byte

	_ = b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		lastHash = bucket.Get([]byte("l"))

		return nil
	})

	newBlock := b.newBlock(data, lastHash)

	err = b.db.Update(func(tx *bolt.Tx) error {
		var (
			data []byte
			err  error
		)
		bucket := tx.Bucket([]byte(blocksBucket))
		data, err = newBlock.Serialize()
		if err != nil {
			return err
		}
		err = bucket.Put(newBlock.Hash, data)
		if err != nil {
			return err
		}
		err = bucket.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			return err
		}
		b.tip = newBlock.Hash

		return nil
	})
	return
}

func (b *blockChain) Iterator() Iterator {
	return &iterator{b.tip, b.db}
}

func (b *blockChain) newBlock(data string, prevBlockHash []byte) (block models.Block) {
	block = models.Block{Timestamp: time.Now().Unix(), Data: []byte(data), PrevBlockHash: prevBlockHash, Hash: []byte{}}
	nonce, hash := b.powGenerator.Generate(block)

	block.Hash = hash[:]
	block.Nonce = nonce
	return
}

func (b *blockChain) newGenesisBlock() models.Block {
	return b.newBlock("Genesis Block", []byte{})
}

func New(powGenerator powGenerator, db *bolt.DB) (BlockChain, error) {
	bc := &blockChain{
		powGenerator: powGenerator,
	}
	var tip []byte

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			var data []byte
			genesis := bc.newGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return err
			}
			data, err = genesis.Serialize()
			if err != nil {
				return err
			}
			err = b.Put(genesis.Hash, data)
			if err != nil {
				return err
			}
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				return err
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	bc.tip = tip
	bc.db = db

	return bc, nil
}
