package main

import (
	"strconv"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Block struct {
	//区块高度
	Height int64
	//上一区块HASH
	PreBlockHash []byte
	//交易数据
	Data []byte
	//时间戳
	Timestamp int64
	//本区块Hash
	Hash []byte
	//工作量证明
	Nonce int64
}

func (block *Block) SetHash() {
	//1，Height 转化为字节数组
	heightBytes := IntToHex(block.Height)
	//2，将时间戳转化为字节数组
	timeString := strconv.FormatInt(block.Timestamp,2)//第二个参数表示转成2进制字符串
	timeBytes := []byte(timeString)
	//3，拼接所有的属性
	blockBytes := bytes.Join([][]byte{heightBytes,block.PreBlockHash,block.Data,timeBytes,block.Hash},[]byte{})
	//4，将拼接的字符数组转成hash
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}

//区块序列化
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

//区块反序列化
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil{
		log.Panic(err)
	}
	return &block
}
//创建新区块
func NewBlock(data string,height int64,preBlockHash []byte) *Block {
	//创建区块
	block := &Block{Height:height,Data:[]byte(data),PreBlockHash:preBlockHash}
	//设置hash
	//block.SetHash()
	//调用工作量证明方法返回有效的hash值和nonce
	pow := NewProofOfWork(block)
	hash,nonce := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

