package db

import (
	"database/sql"
	"log"
	"strconv"
)

func CreateCustomer(phone string) (int, error) {

	str := "INSERT INTO public.\"Customer\" VALUES (nextval('\"Customer_id_seq\"'::regclass), '" + phone + "')"
	println(str)
	_, err := CreateConnection(str)
	if err != nil {
		return -1, err
	}

	return ReadCustomerByPhone(phone)
}

func ReadCustomer(id int) (*sql.Rows, error) {
	return nil, nil
}

func ReadCustomersPhoneById(id int) (string, error) {

	str := "SELECT customer_phone FROM public.\"Customer\" WHERE id = " + strconv.Itoa(id)

	rows, err := CreateConnection(str)
	var col1 string
	if err != nil {
		return "", err
	} else {
		for rows.Next() {

			err := rows.Scan(&col1)
			if err != nil {
				return "", err
			}
		}
		if len(col1) > 0 {
			return col1, nil
		}
		rows.Close()
	}
	return "", nil
}

func ReadCustomerByPhone(phone string) (int, error) {

	var str = "SELECT id FROM public.\"Customer\" WHERE customer_phone = '" + phone + "'"

	rows, err := CreateConnection(str)

	if err != nil {
		return -1, err
	} else {

		for rows.Next() {
			var i string
			err := rows.Scan(&i)
			if err != nil {
				rows.Close()
				log.Fatal(err)
				return -1, err
			} else {
				int2, err := strconv.Atoi(i)
				if err != nil {
					rows.Close()
					return -1, err
				} else {
					rows.Close()
					return int2, nil
				}
			}
		}
		rows.Close()
		return -1, nil
	}
}
