package main

type Blockchain struct {
	Blocks []*Block //存储有序区块
} 
//增加区块到区块中
func (blc *Blockchain) AddBlockToBlockchain(data string,heigth int64,preHash []byte) *Block{
	newBlck := NewBlock(data,heigth,preHash)
	blc.Blocks = append(blc.Blocks,newBlck)
	return newBlck
}
//创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain  {
	//创建创世区块
	genesisBlock := CreateGenesisBlock("GenesisData...")

	return &Blockchain{[]*Block{genesisBlock}}
}