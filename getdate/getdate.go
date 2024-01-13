package getdate

import (
	"challenge-godb/connection"
	"database/sql"
	"fmt"
)

func TimeNow() string {
	db, err := sql.Open("postgres", connection.Psqlinfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Koneksi Berhasil")
	}
	defer db.Close()
	var errx error
	sqlStatement := "SELECT TO_CHAR(CURRENT_TIMESTAMP, 'YYYY-MM-DD HH24:MI:SS');"
	var getdatetime string
	errx = db.QueryRow(sqlStatement).Scan(&getdatetime)
	if errx != nil {
		panic(errx)
	}
	return getdatetime
}

func GetYMD() string {
	db, err := sql.Open("postgres", connection.Psqlinfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Koneksi Berhasil")
	}
	defer db.Close()
	var errx error
	sqlStatement := "SELECT TO_CHAR(CURRENT_TIMESTAMP, 'YYYYMMDD');"
	var getdatetime string
	errx = db.QueryRow(sqlStatement).Scan(&getdatetime)
	if errx != nil {
		panic(errx)
	}
	return getdatetime
}
