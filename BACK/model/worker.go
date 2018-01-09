package model

import (
	"database/sql"
	"../db"
	"../consts"
)

func Login(name string, password string) (string) {
	var tempPass string
	tempPass = db.ReadWokerPassword(name)

	//TODO add URL-s
	if password == tempPass {
		switch name {
		case "manager":
			return "URL1"
		case "boss":
			return "URL2"
		default:
			return ""
		}
	}
	return ""
}

// start of region Методы манагера

/*func SelectById(id int) (string) {
	someShit, err := db.ReadOrder(id)
	if err != nil {
		return err.Error()
	} else {

		return "some string"
	}
}*/

func GetWaitingOrder() ([]int) {
	 masInt, err := db.ReadorderWithParam(consts.ORDER_WAITING)
	 if len(masInt) > 0 && err == nil {
	 	return masInt
	 }
	 return []int{-1}
}

func DeclineOrder(id int) {
	db.UpdateOrder(id, consts.ORDER_DECLINED)
}

func AcceptOrder(id int) {
	db.UpdateOrder(id, consts.ORDER_ACCEPTED)
}

// end of region Методы манагера

// start of region Методы босса

func FindProductById(productId int) (*sql.Rows, error) {
	return nil, nil
}

func FindOrderById(orderId int) (*sql.Rows, error) {
	return nil, nil
}

func FindProductAll() (*sql.Rows, error) {
	return nil, nil
}

func FindOrderAll() (*sql.Rows, error) {
	return nil, nil
}

func SpisatProduct() {

}

// end of region Методы босса
