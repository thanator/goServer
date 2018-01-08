package db

import (
	"database/sql"
	"strconv"
)

func ReadProductById(productId int) (*sql.Rows, error) {
	return nil, nil
}

func ReadProductByParams(typeOfMilc string, fatMilk float64, proizvMilk string) (int, error) {

	/*dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", consts.DB_USER, consts.DB_PASSWORD, consts.DB_NAME)
	dbase, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return -1, err
	}

	err = dbase.Ping()
	if err != nil {
		fmt.Println("Ping error, %s", err)
		return -1, err
	} else {

	}
*/
	stringin := strconv.FormatFloat(fatMilk, 'f', 1, 64)

	var str = "SELECT id FROM public.\"Product\" WHERE milktype = '" + typeOfMilc + "' AND fatness = " + stringin + " AND creator = '" + proizvMilk + "'"

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
	}

	return 0, nil
}

func UpdateProduct(id int) {

}
