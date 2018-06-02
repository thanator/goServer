package model

type State interface {
	handlePositive(int)
	handleNegative(int)
}
