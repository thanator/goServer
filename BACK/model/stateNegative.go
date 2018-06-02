package model

import (
	_ "bufio"

	"../consts"
	"../db"
)

type StateNegative struct{}

func (s *StateNegative) handlePositive(id int) {
	db.UpdateOrder(id, consts.ORDER_WAITING)
}

func (s *StateNegative) handleNegative(id int) {
	println("state negative + handle positive")
}
