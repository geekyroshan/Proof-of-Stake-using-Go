package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	// set random seed
	rand.Seed(time.Now().UnixNano())

	// this will generate an initial PoS network including a blockchain with a genesis block
	genesisTime := time.Now().String()
	pos := &PoSNetwork{
		Blockchain: []*Block{
			{
				Timestamp:     genesisTime,
				PrevHash:      "",
				Hash:          newHash(genesisTime),
				ValidatorAddr: "",
			},
		},
	}
	pos.BlockchainHead = pos.Blockchain[0]

	// instantiate nodes to act as validators in our network
	pos.Validators = pos.NewNode(60)
	pos.Validators = pos.NewNode(40)

	// build 5 additions to the blockchain
	for i := 0; i < 5; i++ {
		winner, err := pos.SelectWinner()
		if err != nil {
			log.Fatal(err)
		}
		winner.Stake += 10
		pos.Blockchain, pos.BlockchainHead, err = pos.GenerateNewBlock(winner)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Round ", i)
		fmt.Println("\tAddress:", pos.Validators[0].Address, "- Stake:", pos.Validators[0].Stake)
		fmt.Println("\tAddress:", pos.Validators[1].Address, "- Stake:", pos.Validators[1].Stake)
	}

	pos.PrintBlockchainInfo()
}
