package model

type Order struct {
	ThisState State
}

func (o *Order) ManageOrderAccept(id int) {
	if id != -1 {
		o.ThisState.handlePositive(id)
	}
}

func (o *Order) ManageOrderDeny(id int) {
	if id != -1 {
		o.ThisState.handleNegative(id)
	}
}

func (o *Order) SetState(state State) {
	o.ThisState = state
}
