package model

import (
	_ "bufio"
)

type StateEmpty struct{}

func (s *StateEmpty) handlePositive(id int) {
	println("IM EMPTY")
}

func (s *StateEmpty) handleNegative(id int) {
	println("IM EMPTY")
}
