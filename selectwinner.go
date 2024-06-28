package main

import (
	"errors"
	"math/rand"
)

func (n PoSNetwork) SelectWinner() (*Node, error) {
	var winnerPool []*Node
	totalStake := 0
	for _, node := range n.Validators {
		if node.Stake > 0 {
			winnerPool = append(winnerPool, node)
			totalStake += node.Stake
		}
	}
	if winnerPool == nil {
		return nil, errors.New("there are no nodes with stake in the network")
	}
	winnerNumber := rand.Intn(totalStake)
	tmp := 0
	for _, node := range n.Validators {
		tmp += node.Stake
		if winnerNumber < tmp {
			return node, nil
		}
	}
	return nil, errors.New("a winner should have been picked but wasn't")
}
