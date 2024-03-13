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
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func CreateGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	//block := Block{time.Now().Unix(), beforeHash, []byte(data), []byte{}}
	//block.SetHash()
	//return &block
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.HashTransactions(), timestamp}, []byte{})
	sum256 := sha256.Sum256(headers)
	b.Hash = sum256[:]
}

func (b *Block) toString() string {
	var result = "Prev.hash: " + hex.EncodeToString(b.PrevBlockHash) + "\n"
	result += "Data:" + hex.EncodeToString(b.HashTransactions()) + "\n"
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

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}
