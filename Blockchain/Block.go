package Blockchain

import (
	"time"
)

type Block struct {
	Timestamp 	 int64
	PreBlockHash []byte
	Data      	 []byte
	Hash 		 []byte
	Nonce		 int
}

/*
 * type BlockHeader struct {
 *     Version    int32
 *     PrevBlock  chainhash.Hash
 *     MerkleRoot chainhash.Hash
 *     Timestamp  time.Time
 *     Bits       uint32
 *     Nonce 	   uint32
 * }
 */


func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte(data), []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
