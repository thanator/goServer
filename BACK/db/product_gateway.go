package db

import (
	"strconv"
	"../consts"
)

func CreateProduct(typeOfMilc string, fatMilk string, proizvMilk string) (int, error) {

	var str = "INSERT INTO public.\"Product\" (id, milktype, fatness, creator, status) VALUES (nextval('\"Product_id_seq\"'::regclass), '" + typeOfMilc + "', " + fatMilk + ", '" + proizvMilk + "', 2)"

	_, err := CreateConnection(str)
	if err != nil {
		return -1, err
	}

	return ReadProductByParams(typeOfMilc, fatMilk, proizvMilk)
}

func ReadAllProducts() (string) {
	var returnString string
	str := "SELECT milktype, fatness, creator, status FROM public.\"Product\""
	rows, err := CreateConnection(str)
	if err != nil {
		return err.Error()
	} else {
		for rows.Next() {
			var col1 string
			var col2 string
			var col3 string
			var col4 int

			err := rows.Scan(&col1, &col2, &col3, &col4)
			if err != nil {
				return err.Error()
			} else {
				returnString += "Тип " + col1 + ", жирность: " + col2 + "\nПроизв. " + col3 + ", статус: " + consts.PRODUCT_STATUS[col4] + "\n"
			}
		}
		rows.Close()
		if len(returnString) > 0 {
			return returnString
		}
		return ""
	}
}

func ReadProductById(productId int) (string) {
	var returnString string

	str := "SELECT milktype, fatness, creator, status FROM public.\"Product\" WHERE id = " + strconv.Itoa(productId)
	rows, err := CreateConnection(str)

	if err != nil {
		return err.Error()
	} else {
		for rows.Next() {
			var col1 string
			var col2 string
			var col3 string
			var col4 int

			err := rows.Scan(&col1, &col2, &col3, &col4)
			if err != nil {
				return err.Error()
			} else {
				returnString += "Тип" + col1 + ", жирность: " + col2 + ", произв. " + col3 + ", статус: " + consts.PRODUCT_STATUS[col4] + "\n"
			}
		}
		rows.Close()
		if len(returnString) > 0 {
			return returnString
		}
		return ""
	}
}

func ReadProductByParams(typeOfMilc string, fatMilk string, proizvMilk string) (int, error) {

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
	//stringin := strconv.FormatFloat(fatMilk, 'f', 1, 64)

	var str = "SELECT id FROM public.\"Product\" WHERE milktype = '" + typeOfMilc + "' AND fatness = " + fatMilk + " AND creator = '" + proizvMilk + "'"

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

	return -1, nil
}

func UpdateProduct(id int) {

}
