package main

import (
	"net/http"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"./db"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	switch r.RequestURI {
	case "/hi":
		w.Write([]byte("OK"))
	case "/database":

		var database *sqlx.DB

		database, _ = db.InitDb()

		if (database != nil) {
			//todo именно тут крч надо сделать нормальный селект
			if rows, err := database.Query("SELECT ALL"); err != nil {
				w.Write([]byte(err.Error()))
			} else {
				var something string
				rows.Scan(&something)
				w.Write([]byte(something))
			}
		}

	default:
		w.Write([]byte("DEF"))
	}

}
