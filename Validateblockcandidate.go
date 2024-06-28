package main

import "errors"

func (n PoSNetwork) ValidateBlockCandidate(newBlock *Block) error {
	if n.BlockchainHead.Hash != newBlock.PrevHash {
		return errors.New("blockchain HEAD hash is not equal to new block previous hash")
	}

	if n.BlockchainHead.Timestamp >= newBlock.Timestamp {
		return errors.New("blockchain HEAD timestamp is greater than or equal to new block timestamp")
	}

	if NewBlockHash(n.BlockchainHead) != newBlock.Hash {
		return errors.New("new block hash of blockchain HEAD does not equal new block hash")
	}
	return nil
}
