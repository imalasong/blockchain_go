package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

const dbFile = "./a.db"
const blocksBucket = "blocks"

func NewBlockchain() *Blockchain {
	//return &Blockchain{[]*Block{CreateGenesisBlock()}}
	var tip []byte
	db, _ := bolt.Open(dbFile, 0600, nil)

	_ = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := CreateGenesisBlock()
			b, _ := tx.CreateBucket([]byte(blocksBucket))
			_ = b.Put(genesis.Hash, genesis.Serialize())
			_ = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip, db}

	return &bc
}

func (bc *Blockchain) AddBlock(data string) {
	//before := chain.Blocks[len(chain.Blocks)-1]
	//block := NewBlock(before.Hash, data)
	//chain.Blocks = append(chain.Blocks, block)
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		fmt.Println("添加块失败1")
		panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		_ = b.Put(newBlock.Hash, newBlock.Serialize())
		_ = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})
	if err != nil {
		fmt.Println("添加块失败2")
		panic(err)
	}
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}
