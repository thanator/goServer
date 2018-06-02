package model

import (
	_ "bufio"

	"../consts"
	"../db"
)

type StateWaiting struct{}

func (s *StateWaiting) handlePositive(id int) {
	db.UpdateOrder(id, consts.ORDER_ACCEPTED)
}

func (s *StateWaiting) handleNegative(id int) {
	db.UpdateOrder(id, consts.ORDER_DECLINED)
}
