package main

import (
	"net/http"
	"github.com/jmoiron/sqlx"
	"database/sql"
	_ "github.com/lib/pq"

	"./model"
	"log"
	"flag"
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
		db := workWithDB()


		//todo именно тут крч надо сделать нормальный селект
		if rows, err := db.dbConn.Query("SELECT ALL"); err != nil {
			w.Write([]byte(err.Error()))
		} else {
			var something string
			rows.Scan(&something)
			w.Write([]byte(something))
		}


	default:
		w.Write([]byte("DEF"))
	}

}

func workWithDB() (*pgDb) {
	var someString string
	flag.StringVar(&someString, "db-connect", "host=/Users/TanDS/Library/Application\\ Support/Postgres/var-10 dbname=temp sslmode=disable", "DB Connect String")
	flag.Parse()

	dbConn, err := sqlx.Connect("postgres", someString)

	if err != nil {
		log.Printf(err.Error());
	}

	p := &pgDb{dbConn: dbConn}
	return p
}

func (p *pgDb) selectPeople() ([]*model.Boss, error) {
	people := make([]*model.Boss, 0)
	if err := p.sqlSelectPeople.Select(&people); err != nil {
		return nil, err
	}
	return people, nil
}

type pgDb struct {
	dbConn *sqlx.DB

	sqlSelectPeople *sqlx.Stmt
	sqlInsertPerson *sqlx.NamedStmt
	sqlSelectPerson *sql.Stmt
}
