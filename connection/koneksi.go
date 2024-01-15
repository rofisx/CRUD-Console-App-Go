package connection

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "sw"
	dbname   = "laundryChallenge"
)

var Psqlinfo = fmt.Sprintf("host =%s port=%d user=%s password = %s dbname=%s sslmode=disable", host, port, user, password, dbname)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", Psqlinfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Koneksi Berhasil")
	}
	return db
}
