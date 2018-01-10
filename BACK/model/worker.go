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

func SelectById(id int) (string) {
	someShit, err := db.ReadOrder(id)
	if err != nil {
		return err.Error()
	} else {
		return someShit
	}
}

func GetWaitingOrder() ([]int) {
	 masInt, err := db.ReadorderWithParam(consts.ORDER_WAITING)
	 if len(masInt) > 0 && err == nil {
	 	return masInt
	 }
	 return []int{-1}
}

func DeclineOrder(id int) (string) {
	err := db.UpdateOrder(id, consts.ORDER_DECLINED)
	if err != nil {
		return err.Error()
	} else {
		return "Вы отклонили заказ!"
	}
}

func AcceptOrder(id int) (string){
	err := db.UpdateOrder(id, consts.ORDER_ACCEPTED)
	if err != nil {
		return err.Error()
	} else {
		return "Вы одобрили заказ!"
	}
}

// end of region Методы манагера

// start of region Методы босса

func FindProductById(productId int) (*sql.Rows, error) {
	return nil, nil
}

func FindOrderById(orderId int) (string, error) {
	return db.ReadOrder(orderId)
}

func FindProductAll() (string) {
	return db.ReadAllProducts()
}

func FindOrderAll() (string) {
	var masInt []int
	var returnStr string

	// пройтись по всем заказам и получить их ID-шники
	masInt1, err := db.ReadorderWithParam(consts.ORDER_WAITING)
	if err != nil {
		return err.Error()
	}
	if masInt1[0] != -1 {
		for _, element := range masInt1 {
			masInt = append(masInt, element)
		}
	}
	masInt2, err := db.ReadorderWithParam(consts.ORDER_ACCEPTED)
	if err != nil {
		return err.Error()
	}
	if masInt2[0] != -1 {
		for _, element := range masInt2 {
			masInt = append(masInt, element)
		}
	}
	masInt3, err := db.ReadorderWithParam(consts.ORDER_DECLINED)
	if err != nil {
		return err.Error()
	}
	if masInt3[0] != -1 {
		for _, element := range masInt3 {
			masInt = append(masInt, element)
		}
	}
	// получить описание заказов
	for _, element := range masInt {
		tempStr,_ := db.ReadOrder(element)
		returnStr += tempStr
	}

	return returnStr
}

func SpisatProduct() {

}

// end of region Методы босса
