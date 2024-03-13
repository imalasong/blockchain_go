package main

import (
	"fmt"
)

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

const subsidy = 50

func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()

	return &tx
}

func (tr *Transaction) SetID() {

}

func (tr *Transaction) IsCoinbase() bool {
	return tr.Vin[0].Vout == -1
}

//
//func NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
//	var inputs []TXInput
//	var outputs []TXOutput
//
//	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
//
//	if acc < amount {
//		log.Panic("ERROR: Not enough funds")
//	}
//
//	// Build a list of inputs
//	for txid, outs := range validOutputs {
//		txID, _ := hex.DecodeString(txid)
//
//		for _, out := range outs {
//			input := TXInput{txID, out, from}
//			inputs = append(inputs, input)
//		}
//	}
//
//	// Build a list of outputs
//	outputs = append(outputs, TXOutput{amount, to})
//	if acc > amount {
//		outputs = append(outputs, TXOutput{acc - amount, from}) // a change
//	}
//
//	tx := Transaction{nil, inputs, outputs}
//	tx.SetID()
//
//	return &tx
//}
