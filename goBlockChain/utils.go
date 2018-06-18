package main

import (
	"encoding/binary"
	"bytes"
	"log"
)

//import "bytes"
//将int64转化为字节数组
func IntToHex(num int64) []byte  {
	buff := new(bytes.Buffer)
	err := binary.Write(buff,binary.BigEndian,num)
	if err!=nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
//产生创世区块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data,1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}
