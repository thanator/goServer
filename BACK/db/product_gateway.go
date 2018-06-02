package db

import (
	"strconv"

	"../consts"
)

func CreateProduct(typeOfMilc string, fatMilk string, proizvMilk string) (int, error) {

	var str = "INSERT INTO public.\"Product\" (id, typemilk, fatness, creator, status) VALUES (nextval('\"Product_id_seq\"'::regclass), '" + typeOfMilc + "', " + fatMilk + ", '" + proizvMilk + "', 2)"

	_, err := CreateConnection(str)
	if err != nil {
		return -1, err
	}

	return ReadProductByParams(typeOfMilc, fatMilk, proizvMilk)
}

func ReadAllProducts() string {
	var returnString string
	str := "SELECT typemilk, fatness, creator, status FROM public.\"Product\" WHERE status <> 1"
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

func ReadProductById(productId int) string {
	var returnString string

	str := "SELECT typemilk, fatness, creator, status FROM public.\"Product\" WHERE id = " + strconv.Itoa(productId)
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
				returnString += "Тип " + col1 + ", жирность: " + col2 + ", произв. " + col3 + ", статус: " + consts.PRODUCT_STATUS[col4] + "\n"
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

	var str = "SELECT id FROM public.\"Product\" WHERE typemilk = '" + typeOfMilc + "' AND fatness = " + fatMilk + " AND creator = '" + proizvMilk + "'"
	println(str)
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
	str := "UPDATE public.\"Product\" SET status = " + strconv.Itoa(1) + " WHERE id = " + strconv.Itoa(id)
	CreateConnection(str)
}

func ReadAllProductsIds() []string {
	var returnMas []string

	str := "SELECT id, status FROM public.\"Product\" WHERE status <> 1"

	rows, err := CreateConnection(str)
	if err != nil {
		return []string{""}
	} else {
		for rows.Next() {
			var tempInt int
			var tempDead int
			err := rows.Scan(&tempInt, &tempDead)
			if err != nil {
				return []string{""}
			} else {
				if tempDead == 4 {
					returnMas = append(returnMas, strconv.Itoa(tempInt)+"-"+consts.PRODUCT_STATUS[4])
				} else {
					returnMas = append(returnMas, strconv.Itoa(tempInt))
				}
			}
		}
		rows.Close()
		if len(returnMas) > 0 {
			return returnMas
		} else {
			return []string{""}
		}
	}

}
