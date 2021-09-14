package blockchain

import (
	"github.com/boltdb/bolt"

	"github.com/Yemanden/block-chain/pkg/models"
)

type Iterator interface {
	Next() (nextBlock models.Block, err error)
}

type iterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (i *iterator) Next() (nextBlock models.Block, err error) {
	err = i.db.View(func(tx *bolt.Tx) error {
		var err error

		bucket := tx.Bucket([]byte(blocksBucket))
		encodedBlock := bucket.Get(i.currentHash)
		err = nextBlock.DeserializeBlock(encodedBlock)
		return err
	})
	if err != nil {
		return
	}

	i.currentHash = nextBlock.PrevBlockHash

	return
}
