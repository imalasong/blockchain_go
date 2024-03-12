package main

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{CreateGenesisBlock()}}
}

func (chain *Blockchain) AddBlock(data string) {
	before := chain.Blocks[len(chain.Blocks)-1]
	block := NewBlock(before.Hash, data)
	chain.Blocks = append(chain.Blocks, block)
}
