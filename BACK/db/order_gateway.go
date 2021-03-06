package db

import (
	"strconv"
	"time"

	"../consts"
)

func CreateOrder(productId int, orderVolume int, date string, customerId int) error {

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

func ReadOrder(orderId int) (string, int, error) {
	var returnString string
	i := -1
	str := "SELECT o.id, p.typemilk, o.volume_all, o.order_type, c.customer_phone, o.order_date, o.deliver_till, o.order_status FROM public.\"Order\" as o INNER JOIN public.\"Product\" as p ON o.product_id = p.id INNER JOIN public.\"Customer\" as c ON o.customer_id = c.id WHERE o.id = " + strconv.Itoa(orderId)

	rows, err := CreateConnection(str)
	if err != nil {
		return "", i, err
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
				return "", i, err
			} else {
				i = col8
				returnString += "Заказ №" + col1 + ", тип: " + col2 + "," + consts.TYPE_OF_ORDER[col4] + ", статус: " + consts.ORDER_STATUS[col8] + ", номер заказчика: " + col5 + "\n"
			}
		}
		if len(returnString) > 0 {
			return returnString, i, nil
		}
		rows.Close()
	}
	return "", i, nil
}

func ReadorderWithParam(status int) ([]int, error) {
	var masInt []int

	str := "SELECT id FROM public.\"Order\" WHERE order_status = " + strconv.Itoa(status)

	rows, err := CreateConnection(str)
	if err != nil {
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

func UpdateOrder(orderId int, status int) error {
	str := "UPDATE public.\"Order\" SET order_status = " + strconv.Itoa(status) + " WHERE id = " + strconv.Itoa(orderId)
	_, err := CreateConnection(str)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func ReadAllOrderIds() []int {
	var returnMas []int

	str := "SELECT id FROM public.\"Order\""

	rows, err := CreateConnection(str)
	if err != nil {
		return []int{-1}
	} else {
		for rows.Next() {
			var tempInt int
			err := rows.Scan(&tempInt)
			if err != nil {
				return []int{-1}
			} else {
				returnMas = append(returnMas, tempInt)
			}
		}
		rows.Close()
		if len(returnMas) > 0 {
			return returnMas
		} else {
			return []int{-1}
		}
	}
}

func ExportOrder() string {

	var returnStr string

	str := "select xmlelement(name Order, xmlattributes(o.id, cus.customer_phone), xmlelement(name Milk_type, prod.typemilk)) from \"Order\" as o inner join \"Customer\" as cus on o.customer_id = cus.id inner join \"Product\" as prod on o.product_id = prod.id where current_date - o.order_date < 30"

	rows, err := CreateConnection(str)

	if err != nil {
		return err.Error()
	} else {
		for rows.Next() {
			var tempStr string
			err := rows.Scan(&tempStr)
			if err != nil {
				return err.Error()
			} else {
				returnStr = returnStr + tempStr + "\n"
			}
		}
		rows.Close()
		if len(returnStr) > 0 {
			return returnStr
		} else {
			return ""
		}
	}

}
