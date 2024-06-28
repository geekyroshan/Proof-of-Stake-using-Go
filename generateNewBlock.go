package main
import "time"


func (n PoSNetwork) GenerateNewBlock(Validator *Node) ([]*Block, *Block, error) {
	if err := n.ValidateBlockchain(); err != nil {
		Validator.Stake -= 10
		return n.Blockchain, n.BlockchainHead, err
	}

	currentTime := time.Now().String()

	newBlock := &Block {
		Timestamp: currentTime,
		PrevHash: n.BlockchainHead.Hash,
		Hash: NewBlockHash(n.BlockchainHead),
		ValidatorAddr: Validator.Address,
	}

	if err := n.ValidateBlockCandidate(newBlock); err != nil {
		Validator.Stake -= 10
		return n.Blockchain, n.BlockchainHead, err
	} else {
		n.Blockchain = append(n.Blockchain, newBlock)
	}
	return n.Blockchain, newBlock, nil
}