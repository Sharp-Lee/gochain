package Blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
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

// 创建新区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte(data), []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// 序列化block为一个字节数组
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil{
		log.Panic(err)
	}

	return result.Bytes()
}

// 解码序列化后的字节数组，返回block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil{
		log.Panic(err)
	}

	return &block
}