package core

import (
	"log"
	"fmt"
)

type Blockchain struct {
	Blocks []*Block
}

//传入区块数据
func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.Apendchain(&newBlock)
}

//创建新区块链
func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.Apendchain(&genesisBlock)
	return &blockchain
}

//新增区块
func (bc *Blockchain) Apendchain(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block ")
	}
}

//打印区块信息
func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("index: %d\n", block.Index)
		fmt.Printf("pre.hash: %s\n", block.PrevBlockHash)
		fmt.Printf("curr.hash: %s\n", block.Hash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("timestamp: %d\n", block.Timestamp)
		fmt.Println()
	}
}

//校验区块
func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}

	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}

	if newBlock.Hash != calculateHash(newBlock) {
		return false
	}
	return true
}
