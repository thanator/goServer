package db

import (
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

func ReadOrder(orderId int) (string, error) {
	var returnString string

	str := "SELECT o.id, p.milktype, o.volume_all, o.order_type, c.customer_phone, o.order_date, o.deliver_till, o.order_status FROM public.\"Order\" as o INNER JOIN public.\"Product\" as p ON o.product_id = p.id INNER JOIN public.\"Customer\" as c ON o.customer_id = c.id WHERE o.id = " + strconv.Itoa(orderId)

	rows, err := CreateConnection(str)
	if err != nil {
		return "", err
	} else {
		for rows.Next() {
			var col1 string
			var col2 string
			var col3 string
			var col4 int
			var col5 string
			var col6 string
			var col7 string
			var col8 int
			err := rows.Scan(&col1, &col2, &col3, &col4, &col5, &col6, &col7, &col8)
			if err != nil {
				return "", err
			} else {
				returnString += "Заказ №" + col1 + ", тип: " + col2 + "," + consts.TYPE_OF_ORDER[col4] + ", статус: " + consts.ORDER_STATUS[col8] + "\n"
			}
		}
		if len(returnString) > 0 {
			return returnString, nil
		}
		rows.Close()
	}
	return "", nil
}

func ReadorderWithParam(status int) ([]int, error) {
	var masInt []int

	str := "SELECT id FROM public.\"Order\" WHERE order_status = " + strconv.Itoa(status)

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

func UpdateOrder(orderId int, status int) (error) {
	str := "UPDATE public.\"Order\" SET order_status = " + strconv.Itoa(status) + " WHERE id = " + strconv.Itoa(orderId)
	_, err := CreateConnection(str)
	if err!= nil {
		return err
	} else {
		return nil
	}

}
