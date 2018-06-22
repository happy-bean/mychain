package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

//区块
type Block struct {
	//区块头
	//区块编号
	Index int64
	//区块时间戳 创建区块的时间
	Timestamp int64
	//上一个区块的hash值
	PrevBlockHash string
	//当前区块的hash值
	Hash string

	//区块体
	//区块数据
	Data string
}

//计算hash值
func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

//创建区块
func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

//创建创世区块
func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	preBlock = GenerateNewBlock(preBlock, "genesis block")
	return preBlock
}
