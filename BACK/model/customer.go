package model

import (
	"time"
	"../db"
)

/*
func SeeAll() {

}*/

func MakeOrder(typeOfMilc string, milkVolume int, fatMilk int, deliveryDate time.Time, proizvMilk string, phoneNumber string) {
	var idOfMilk = db.ReadProductByParams(typeOfMilc, fatMilk, proizvMilk)
	var idOfCustomer = db.ReadCustomerByPhone(phoneNumber)

	db.CreateOrder(idOfMilk, milkVolume, deliveryDate, idOfCustomer)
}
