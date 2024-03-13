package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func CreateGenesisBlock() *Block {
	return NewBlock("Genesis block", []byte{})
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	//block := Block{time.Now().Unix(), beforeHash, []byte(data), []byte{}}
	//block.SetHash()
	//return &block
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	sum256 := sha256.Sum256(headers)
	b.Hash = sum256[:]
}

func (b *Block) toString() string {
	var result = "Prev.hash: " + hex.EncodeToString(b.PrevBlockHash) + "\n"
	result += "Data:" + string(b.Data) + "\n"
	result += "Hash:" + hex.EncodeToString(b.Hash) + "\n"
	result += "Nonce:" + strconv.FormatInt(int64(b.Nonce), 10)
	return result
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)

	if err != nil {
		panic("序列化异常" + b.toString())
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)

	if err != nil {
		panic("反序列化异常")
	}

	return &block
}
