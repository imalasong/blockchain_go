package main

import "fmt"

func main() {
	//blockchain := NewBlockchain()
	//
	//blockchain.AddBlock("This is XiaoHong")
	//blockchain.AddBlock("Is this XiaoHong")
	//blockchain.AddBlock("Who is this")
	//
	//for _, block := range blockchain.Blocks {
	//	pow := NewProofOfWork(block)
	//	fmt.Println(block.toString())
	//	fmt.Printf("Valid: %v \n", pow.Validate())
	//	fmt.Println()
	//}

	//bc := NewBlockchain()
	//defer bc.db.Close()
	//
	//cli := CLI{bc}
	//cli.Run()

	//data1 := []byte("I like donuts")
	//data2 := []byte("I like donutsca07ca")
	//targetBits := 24
	//target := big.NewInt(1)
	//target.Lsh(target, uint(256-targetBits))
	//fmt.Printf("%x\n", sha256.Sum256(data1))
	//fmt.Printf("%64x\n", target)
	//fmt.Printf("%x\n", sha256.Sum256(data2))
	//
	//fmt.Println(len(sha256.Sum256(data1)))
	//fmt.Println(sha256.Sum256(data1))
out:
	for i := 1; i < 10; i++ {
		if i == 5 {
			continue out
		}
		fmt.Println(i)
	}
	for i := 100; i < 110; i++ {
		//if i == 105 {
		//	continue out
		//}
		fmt.Println(i)
	}

	fmt.Println("2222")
}
