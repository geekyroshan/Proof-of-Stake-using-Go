package main
import 
"fmt"



func (pos *PoSNetwork) PrintBlockchainInfo() {
	fmt.Println("\nBlockchain Information:")
	for i, block := range pos.Blockchain {
			fmt.Printf("\nBlock %d Info:\n", i)
			fmt.Printf("\tTimestamp: %s\n", block.Timestamp)
			fmt.Printf("\tPrevious Hash: %s\n", block.PrevHash)
			fmt.Printf("\tHash: %s\n", block.Hash)
			fmt.Printf("\tValidator Address: %s\n", block.ValidatorAddr)
	}
}




type PoSNetwork struct {
	Blockchain     []*Block
	BlockchainHead *Block
	Validators     []*Node
}

type Node struct {
	Stake   int
	Address string
}

type Block struct {
	Timestamp     string
	PrevHash      string
	Hash	      string
	ValidatorAddr string
}