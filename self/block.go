package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	PreHash   []byte
	Data      []byte
	Hash      []byte
}

func CreateGenesisBlock() *Block {
	return NewBlock([]byte{}, "Genesis block")
}

func NewBlock(beforeHash []byte, data string) *Block {
	block := Block{time.Now().Unix(), beforeHash, []byte(data), []byte{}}
	block.SetHash()
	return &block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PreHash, b.Data, timestamp}, []byte{})
	sum256 := sha256.Sum256(headers)
	b.Hash = sum256[:]
}

func (b *Block) toString() string {
	var result string = "Prev.hash: " + string(b.PreHash) + "\n"
	result += "Data:" + string(b.Data) + "\n"
	result += "Hash:" + string(b.Hash) + "\n"
	return result
}
