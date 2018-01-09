package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"log"
	_ "./db"
	"./model"
	"./consts"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":63342", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	var err string

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//todo CASES
	switch r.RequestURI {
	case "/hi":
		w.Write([]byte("OK"))
	case "/goProjectSecond/FRONT/making_order.html":
		switch r.Method {
		case "POST":
			if err := r.ParseForm(); err != nil {

			}
			volume := r.FormValue("volume")
			fmt.Fprintf(w, "Volume = %s\n", volume)
			log.Println(volume)
		}
	case "/manager":
		model.GetWaitingOrder()
	case "/manager/accept":
		//model.AcceptOrder()
	case "/making_order/create":
		 err = model.MakeOrder("Молоко", 1, 2.5, "20171201", "Домик в деревне", "+780053535")
		 if err != "succ" {
			w.Write([]byte(err))
		}
	case "/database":

		dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", consts.DB_USER, consts.DB_PASSWORD, consts.DB_NAME)
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

func createOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		//t, _ := template.ParseFiles
	} else {
		r.ParseForm()
		fmt.Println("Request", r.Form["volume"])
	}
}