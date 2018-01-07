package BACK

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"log"
	_"./db"
	"./model"
)

const (
	DB_USER     = "TanDS"
	DB_PASSWORD = "6364"
	DB_NAME     = "MilkDB"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {

	switch r.RequestURI {
	case "/hi":
		w.Write([]byte("OK"))
	case "/delete":

	//todo CASES
	model.Zakazat()

	case "/database":

		dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
		dbase, err := sql.Open("postgres", dbinfo)

		if err != nil {
			w.Write([]byte(err.Error()))
		}
		//checkErr(err)

		err = dbase.Ping()
		if err != nil {
			fmt.Println("Ping error, %s", err)
			w.Write([]byte(err.Error()))
		} else {

		}

		rows, err := dbase.Query("SELECT * FROM test")

		if err != nil {
			w.Write([]byte(err.Error()))
		} else {

			for rows.Next() {
				var ivan string
				var artem string
				err := rows.Scan(&ivan, &artem)
				if err != nil {
					log.Fatal(err)
				} else {
					w.Write([]byte("artem:" + artem))
					w.Write([]byte("ivan:" + ivan + "\n"))
				}
			}
			rows.Close()
		}

		defer dbase.Close()

	default:
		w.Write([]byte("DEF"))
	}

}
