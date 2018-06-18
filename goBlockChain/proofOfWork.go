package main

import (
	"math/big"
	"bytes"
	"crypto/sha256"
)

type ProofOfWork struct {
	Block *Block
	target *big.Int
}

func (pow *ProofOfWork) prepareData(nonce int) []byte{
	data := bytes.Join([][]byte{
		pow.Block.PreBlockHash,
		pow.Block.Data,
		IntToHex(pow.Block.Timestamp),
		IntToHex(int64(targetBit)),
		IntToHex(int64(nonce)),
		IntToHex(int64(pow.Block.Height)),
	},
	[]byte{})
	return data
}
//验证区块是否有效
func (proofOfWork *ProofOfWork) IsValid() bool {
	var hashInt big.Int
	hashInt.SetBytes(proofOfWork.Block.Hash)
	if proofOfWork.target.Cmp(&hashInt) == 1{
		return true
	}
	return false
}
func (proofOfWork *ProofOfWork) Run() ([]byte,int64) {

	//1.将Block的属性拼接成字节数组
	//2.生成hash
	//3.判断hash有效性如果满足条件，跳出循环
	nonce := 0
	var hashInt big.Int //存储新生成的hash
	var hash [32]byte
	for {
		//准备数据
		dataBytes := proofOfWork.prepareData(nonce)
		//生成hash
		hash = sha256.Sum256(dataBytes)
		//将hash存储到hashInt
		hashInt.SetBytes(hash[:])
		//判断hashInt是否小于Block里面的target
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce = nonce + 1
	}
	return hash[:],int64(nonce)
}

const targetBit = 16  //256位hash值里前面至少有16个0
//新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {

	//1.创建一个初始值为1的target
	target := big.NewInt(1)
	//2.左移256-targetBit位
	target = target.Lsh(target,256-targetBit)

	return &ProofOfWork{Block: block,target:target}
}
