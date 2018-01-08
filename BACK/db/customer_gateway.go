package db

import (
	"database/sql"
	"strconv"
	"log"
)

func CreateCustomer(phone string) (int, error) {

	str := "INSERT INTO public.\"Customer\" VALUES (nextval('\"Customer_id_seq\"'::regclass), '" + phone + "')"

	_, err := CreateConnection(str)
	if err != nil {
		return -1, err
	}

	return ReadCustomerByPhone(phone)
}

func ReadCustomer(id int) (*sql.Rows, error) {

	return nil, nil
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
				return -1, err
				log.Fatal(err)
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
