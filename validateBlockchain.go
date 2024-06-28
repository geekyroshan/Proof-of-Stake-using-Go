package main
import "errors"

func (n PoSNetwork) ValidateBlockchain() error {
	if len(n.Blockchain) <= 1 {
		return nil
	}

	currBlockIdx := len(n.Blockchain)-1
	prevBlockIdx := len(n.Blockchain)-2

	for prevBlockIdx >= 0 {
		currBlock := n.Blockchain[currBlockIdx]
		prevBlock := n.Blockchain[prevBlockIdx]
		if currBlock.PrevHash != prevBlock.Hash {
			return errors.New("blockchain has inconsistent hashes")
		}

		if currBlock.Timestamp <= prevBlock.Timestamp {
			return errors.New("blockchain has inconsistent timestamps")
		}

		if NewBlockHash(prevBlock) != currBlock.Hash {
			return errors.New("blockchain has inconsistent hash generation")
		}
		currBlockIdx--
		prevBlockIdx--
	}
	return nil
}