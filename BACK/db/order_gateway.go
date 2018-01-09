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

func ReadorderWithParam(status int) ([]int, error) {
	var masInt []int

	str := "SELECT id FROM public.\"Order\" WHERE order_status = " + strconv.Itoa(consts.ORDER_WAITING)

	rows, err := CreateConnection(str)
	if err!= nil {
		return []int{-1, -1}, err
	} else {

		for rows.Next() {
			var tempInt int
			err := rows.Scan(&tempInt)
			if err != nil {
				return []int{-1, -1}, err
			} else {
				masInt = append(masInt, tempInt)
			}
		}
		if len(masInt) > 0 {
			return masInt, nil
		}
		rows.Close()
	}

	return []int{-1, -1}, nil

}

func UpdateOrder(orderId int, status int) {
}
