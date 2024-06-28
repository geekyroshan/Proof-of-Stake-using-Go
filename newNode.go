package main
import (
   "math/rand"
	 "fmt" 
)

func (n PoSNetwork) NewNode(stake int) []*Node {
	newNode := &Node{
		Stake:   stake,
		Address: randAddress(),
	}
	n.Validators = append(n.Validators, newNode)
	return n.Validators
}

func randAddress() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}