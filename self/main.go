package main

import "fmt"

func main() {
	blockchain := NewBlockchain()

	blockchain.AddBlock("This is XiaoHong")
	blockchain.AddBlock("Is this XiaoHong")
	blockchain.AddBlock("Who is this")

	for _, block := range blockchain.Blocks {
		fmt.Println(block.toString())
	}
}
