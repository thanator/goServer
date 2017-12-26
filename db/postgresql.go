package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"flag"
	"log"
)

type Config struct {
	ConnectString string
}

func InitDb() (*sqlx.DB, error) {
	var someString string
	flag.StringVar(&someString, "db-connect", "host=/Users/TanDS/Library/Application\\ Support/Postgres/var-10 dbname=temp sslmode=disable", "DB Connect String")
	flag.Parse()

	if dbConn, err := sqlx.Connect("postgres", someString); err != nil {
		log.Printf(err.Error());
		return nil, err
	} else {
		//p := &pgDb{dbConn: dbConn}
		return dbConn, nil
	}

}

type pgDb struct {
	dbConn *sqlx.DB

	sqlSelectPeople *sqlx.Stmt
	sqlInsertPerson *sqlx.NamedStmt
	sqlSelectPerson *sql.Stmt
}
