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
	//http.HandleFunc("/", foo)
	http.HandleFunc("/", createOrder)
	http.ListenAndServe(":3000", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	var err string

	//todo CASES
	switch r.RequestURI {
	case "/hi":
		w.Write([]byte("OK"))
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
	switch r.URL.Path {
	case "/making_order.html":
		fmt.Printf("/ making order\n")
	case "/making_order":
		fmt.Printf("making_order\n")
	}
	//if r.URL.Path != "/" {
	//	fmt.Printf(r.URL.Path)
	//	fmt.Fprintf(w, "kekeke\n")
    //    http.Error(w, "404 not found.", http.StatusNotFound)
    //    return
    //}
 	fmt.Printf("not error\n")
    switch r.Method {
    case "GET":
    	var str string
    	if r.URL.Path == "/" {
    		str = "../FRONT/index.html"
    	} else {
    		fmt.Printf("Getted\n")
    		str = "../FRONT/" + r.URL.Path
    	}
    	fmt.Printf(str)     
        http.ServeFile(w, r, str)
    	
    case "POST":
    	fmt.Printf("posted\n")
        // Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
        if err := r.ParseForm(); err != nil {
            fmt.Printf("ParseForm() err: %v", err)
            return
        }
        fmt.Printf("Post from website! r.PostFrom = %v\n", r.PostForm)
        volume := r.FormValue("volume")
        //address := r.FormValue("address")
        fmt.Printf("Volume = %s\n", volume)
        http.Redirect(w, r, "/", http.StatusSeeOther)
        //fmt.Fprintf(w, "Address = %s\n", address)
    default:
        fmt.Printf("Sorry, only GET and POST methods are supported.")
    }
}