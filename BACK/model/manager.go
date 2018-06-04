package model

import (
	"../consts"
	"../db"
)

type ManagerWorker struct{}

func (e *ManagerWorker) Accept(visitor Visitor) {
	visitor.visitManager()
}

// start of region Методы манагера

func FindOrderDetailsById(orderId int) (string, int, error) {
	str1, intin, err := db.ReadOrder(orderId)
	if err != nil {
		return "", -1, err
	}

	//	str2, err := ReadCustomersPhoneById()

	return str1, intin, nil
}

func SelectById(id int) string {
	str, _, err := db.ReadOrder(id)
	if err != nil {
		return err.Error()
	} else {
		return str
	}
}

func GetWaitingOrder() []int {
	masInt, err := db.ReadorderWithParam(consts.ORDER_WAITING)
	if len(masInt) > 0 && err == nil {
		return masInt
	}
	return []int{-1}
}

func DeclineOrder(id int) string {
	err := db.UpdateOrder(id, consts.ORDER_DECLINED)
	if err != nil {
		return err.Error()
	} else {
		return "Вы отклонили заказ!"
	}
}

func AcceptOrder(id int) string {
	err := db.UpdateOrder(id, consts.ORDER_ACCEPTED)
	if err != nil {
		return err.Error()
	} else {
		return "Вы одобрили заказ!"
	}
}

// end of region Методы манагера
