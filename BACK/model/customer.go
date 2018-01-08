package model

import (
	"../db"
)

/*
func SeeAll() {

}*/

func MakeOrder(typeOfMilc string, milkVolume int, fatMilk float64, deliveryDate string, proizvMilk string, phoneNumber string) (string) {

	var idOfMilk, err = db.ReadProductByParams(typeOfMilc, fatMilk, proizvMilk)
	if err!=nil {
		return err.Error()
	}
	var idOfCustomer = db.ReadCustomerByPhone(phoneNumber)

	db.CreateOrder(idOfMilk, milkVolume, deliveryDate, idOfCustomer)

	return ""
}
