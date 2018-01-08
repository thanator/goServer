package model

import (
	"../db"
)

/*
func SeeAll() {

}*/

func MakeOrder(typeOfMilc string, milkVolume int, fatMilk float64, deliveryDate string, proizvMilk string, phoneNumber string) (string) {

	idOfMilk, err := db.ReadProductByParams(typeOfMilc, fatMilk, proizvMilk)
	if err != nil {
		return err.Error()
	}
	idOfCustomer, err := db.ReadCustomerByPhone(phoneNumber)
	if err != nil {
		return err.Error()
	} else if idOfCustomer == -1 {
		idOfCustomer, err = db.CreateCustomer(phoneNumber)
		if err != nil {
			return err.Error()
		}
	}

	err = db.CreateOrder(idOfMilk, milkVolume, deliveryDate, idOfCustomer)
	if err != nil {
		return err.Error()
	}

	return "succ"
}
