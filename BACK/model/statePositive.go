package model

import (
	_ "bufio"
)

type StatePositive struct{}

func (s *StatePositive) handlePositive(id int) {
	println("state positive + handle positive")
}

func (s *StatePositive) handleNegative(id int) {
	println("state positive + handle positive")
}
