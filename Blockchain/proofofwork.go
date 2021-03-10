package Blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"gochain/utils"
	"math"
	"math/big"
)

// 难度值，表示hash的前24位必须是0
const targetBits = 24
const maxNonce = math.MaxInt64

/*
 * 工作量证明对象结构体
 * blockf: 指向当前需要证明的块
 * target：计算寻找的hash必须要小于target
 */
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// 创建工作量证明对象，此处默认设置目标值为 0x0000010000...（58个0）
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target,uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

// 工作量证明用到的数据有：PrevBlockHash, Data, Timestamp, targetBits, nonce
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PreBlockHash,
			pow.block.Data,
			utils.IntToHex(pow.block.Timestamp),
			utils.IntToHex(int64(targetBits)),
			utils.IntToHex(int64(nonce)),
		},
		[]byte{})

	return data
}

// 寻找有效hash
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r%x\n", hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("Nonce \"%v\"\n", nonce)
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// 验证工作量
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
