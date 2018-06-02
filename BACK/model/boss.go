package model

import (
	"../consts"
	"../db"
)

type BossWorker struct{}

func (e *BossWorker) Accept(visitor Visitor) {
	visitor.visitBoss()
}

// start of region Методы босса

func FindProductById(productId int) string {
	return db.ReadProductById(productId)
}

func FindOrderById(orderId int) (string, int, error) {
	return db.ReadOrder(orderId)
}

func FindProductAll() string {
	return db.ReadAllProducts()
}

func FindOrderAll() string {
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
		tempStr, _, _ := db.ReadOrder(element)
		returnStr += tempStr
	}

	return returnStr
}

func FindAllProductIds() []string {
	return db.ReadAllProductsIds()
}

func FindAllOrderIds() []int {

	var masInt []int

	// пройтись по всем заказам и получить их ID-шники
	masInt1, err := db.ReadorderWithParam(consts.ORDER_WAITING)
	if err != nil {
		return []int{-1}
	}
	if masInt1[0] != -1 {
		for _, element := range masInt1 {
			masInt = append(masInt, element)
		}
	}
	masInt2, err := db.ReadorderWithParam(consts.ORDER_ACCEPTED)
	if err != nil {
		return []int{-1}
	}
	if masInt2[0] != -1 {
		for _, element := range masInt2 {
			masInt = append(masInt, element)
		}
	}
	masInt3, err := db.ReadorderWithParam(consts.ORDER_DECLINED)
	if err != nil {
		return []int{-1}
	}
	if masInt3[0] != -1 {
		for _, element := range masInt3 {
			masInt = append(masInt, element)
		}
	}
	return masInt
}

func SpisatProduct(id int) {
	db.UpdateProduct(id)
}

// end of region Методы босса
