package db

import (
	"database/sql"
	"time"
	"strconv"
	"../consts"
)

func CreateOrder(productId int, orderVolume int, date string, customerId int) (error) {

	someDate1, someDate2, someDate3 := time.Now().Date()

	currentDate := someDate2.String() + " " + strconv.Itoa(someDate3) + ", " + strconv.Itoa(someDate1)

	str := "INSERT INTO public.\"Order\" VALUES (nextval('\"Order_id_seq\"'::regclass), " + strconv.Itoa(productId) + ", " + strconv.Itoa(orderVolume) + ", " + strconv.Itoa(consts.ORDER_SELL) + ", " + strconv.Itoa(customerId) + ", '" + currentDate + "', '" + date + "', " + strconv.Itoa(consts.ORDER_WAITING) + ")"

	_, err := CreateConnection(str)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func ReadOrder(orderId int) (*sql.Rows, error) {
	return nil, nil
}

func ReadorderWithParam(status int) (*sql.Rows, error) {
	return nil, nil
}

func UpdateOrder(orderId int, status int) {
}
