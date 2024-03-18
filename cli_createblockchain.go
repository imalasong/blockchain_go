package main

import (
	"fmt"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	//if !ValidateAddress(address) {
	//	log.Panic("ERROR: Address is not valid")
	//}
	bc := CreateBlockchain(address, nodeID)
	defer bc.db.Close()

	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
