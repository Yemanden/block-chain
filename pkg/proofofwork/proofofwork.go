package proofofwork

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/Yemanden/block-chain/pkg/models"
)

const (
	targetBits = 24 // 24 бита в 16ричной системе - 6 нулей в левой части хэша
	maxNonce   = math.MaxInt64
)

type ProofOfWork interface {
	Run() (nonce int, hashByte []byte)
	Validate() bool
}

type proofOfWork struct {
	block  models.Block
	target *big.Int
}

func (p *proofOfWork) Run() (nonce int, hashByte []byte) {
	var hashInt big.Int
	var hash [32]byte

	fmt.Printf("Mining the block containing \"%s\"\n", p.block.Data)
	for nonce = 0; nonce < maxNonce; nonce++ {
		data := p.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		// преобразуем хэш к численному виду для более простого сравнения
		hashInt.SetBytes(hash[:])
		// если полученный хэш меньше таргетного
		if hashInt.Cmp(p.target) == -1 {
			break
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (p *proofOfWork) Validate() bool {
	var hashInt big.Int

	data := p.prepareData(p.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(p.target) == -1

	return isValid
}

func (p *proofOfWork) prepareData(nonce int) []byte {
	// объединяем данные, в качестве разделителя - ничего
	data := bytes.Join(
		[][]byte{
			p.block.PrevBlockHash,
			p.block.Data,
			intToHex(p.block.Timestamp),
			intToHex(int64(targetBits)),
			intToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func intToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func NewProofOfWork(b models.Block) ProofOfWork {
	target := big.NewInt(1)                  // 63 нуля и одна 1
	target.Lsh(target, uint(256-targetBits)) // смещение влево на (256-targetBits)/4 символов. 4 - количество бит в символе hex

	pow := &proofOfWork{b, target}

	return pow
}
