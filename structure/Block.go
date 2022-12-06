package structure

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

type Block struct {
	Height        int64
	PrevBlockHash []byte
	Data          []byte
	Timestamp     int64
	Hash          []byte
	Nonce         int64
}

func (block *Block) SetHash() {
	heightBytes := IntToHex(block.Height)
	timeBytes := []byte(strconv.FormatInt(block.Timestamp, 2))

	blockBytes := bytes.Join(heightBytes,block.PrevBlockHash,block.Data,timeBytes)
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}


func NewBlock(data string,height int64, prevBlockHash []byte) *Block {
	block := &Block{
		height,
		PrevBlockHash,
		data
	}
	block.SetHash()
}