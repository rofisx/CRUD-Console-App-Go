package mastersatuan

import (
	"bufio"
	"challenge-godb/connection"
	"challenge-godb/getdate"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

type Satuan struct {
	SatuanId    string
	NamaSatuan  string
	InDate      string
	InBy        string
	UpdatedDate string
	UpdatedBy   string
}

func ShowAllSatuan() []Satuan {
	db, err := sql.Open("postgres", connection.Psqlinfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	sqlStatement := "SELECT satuanid,nama,indate,inby,COALESCE(updateddate,'3000/01/01'),COALESCE(updatedby,'NULL') FROM mst_satuan;"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	xsatuan := scanSatuan(rows)
	if len(xsatuan) <= 0 {
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println("Satuan Kosong, Silahkan Input Satuan")
		fmt.Println(strings.Repeat("-", 50))
	}
	return xsatuan
}

func InsertSatuan() bool {
	var satuan Satuan
	var res bool = false
	db := connection.ConnectDB()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	satuan.SatuanId = inputSatuanId()
	satuan.NamaSatuan = inputNamaSatuan()
	satuan.InBy = inputUseby()
	satuan.InDate = getdate.TimeNow()
	insert(satuan, tx)
	err = tx.Commit()
	if err != nil {
		res = false
		// panic(err)
	} else {
		fmt.Println("Data Tersimpan")
		res = true
	}
	return res
}

func UpdateSatuan() bool {
	var satuan Satuan
	var res bool = false
	xsatuan := inputSatuanIdNotExist()
	if len(xsatuan) > 0 {
		if CheckSatuanById(xsatuan) {
			db := connection.ConnectDB()
			defer db.Close()
			tx, err := db.Begin()
			if err != nil {
				panic(err)
			}
			satuan.SatuanId = xsatuan
			satuan.NamaSatuan = inputNamaSatuan()
			satuan.UpdatedBy = updateUseby()
			satuan.UpdatedDate = getdate.TimeNow()
			update(satuan, tx)
			err = tx.Commit()
			if err != nil {
				res = false
				// panic(err)
			} else {
				fmt.Println("Update Tersimpan")
				res = true
			}
		} else {
			fmt.Println(strings.Repeat("-", 50))
			fmt.Println("Satuan Tidak Ada")
			res = false

		}
	}
	return res
}

func DeleteSatuan() bool {
	var res bool = false
	satuanid := inputSatuanIdNotExist()
	if CheckSatuanById(satuanid) {
		db := connection.ConnectDB()
		defer db.Close()
		tx, err := db.Begin()
		if err != nil {
			panic(err)
		}
		delete(satuanid, tx)
		err = tx.Commit()
		if err != nil {
			panic(err)
		} else {
			res = true
		}
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println("Satuan Terhapus")
	} else {
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println("Satuan Tidak Ada")
		res = false
	}
	return res
}

func inputSatuanId() string {
	scanner := bufio.NewScanner(os.Stdin)
	var satuanid string
	loop := 1
	for i := 0; i < loop; i++ {
		fmt.Print("Masukan Id Satuan : ")
		scanner.Scan()
		if len(scanner.Text()) > 0 {
			satuanid = scanner.Text()
			if CheckSatuanById(satuanid) {
				fmt.Println("Satuan Id Sudah Ada!, Silahkan Coba Lagi")
				loop++
			}
		} else {
			fmt.Println("Input Id Satuan Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return satuanid
}
func inputSatuanIdNotExist() string {
	scanner := bufio.NewScanner(os.Stdin)
	var satuanid string
	loop := 1
	for i := 0; i < loop; i++ {
		fmt.Print("Masukan Id Satuan : ")
		scanner.Scan()
		if len(scanner.Text()) > 0 {
			satuanid = scanner.Text()
		} else {
			fmt.Println("Input Id Satuan Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return satuanid
}
func inputUseby() string {
	scanner := bufio.NewScanner(os.Stdin)
	var userby string
	loop := 1
	for i := 0; i < loop; i++ {
		fmt.Print("Dibuat Oleh User : ")
		scanner.Scan()
		if len(scanner.Text()) > 0 {
			userby = scanner.Text()
		} else {
			fmt.Println("Input Nama User Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return userby
}
func updateUseby() string {
	scanner := bufio.NewScanner(os.Stdin)
	var userby string
	loop := 1
	for i := 0; i < loop; i++ {
		fmt.Print("Diupdate Oleh User : ")
		scanner.Scan()
		if len(scanner.Text()) > 0 {
			userby = scanner.Text()
		} else {
			fmt.Println("Input Nama User Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return userby
}
func inputNamaSatuan() string {
	scanner := bufio.NewScanner(os.Stdin)
	var namaSatuan string
	loop := 1
	for i := 0; i < loop; i++ {
		fmt.Print("Masukan Nama Satuan : ")
		scanner.Scan()
		if len(scanner.Text()) > 0 {
			namaSatuan = scanner.Text()
		} else {
			fmt.Println("Input Nama Satuan Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return namaSatuan
}

func CheckSatuanById(id string) bool {
	db, err := sql.Open("postgres", connection.Psqlinfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var result bool = false
	var errx error
	selectSatuan := "SELECT satuanid,nama,indate,inby,COALESCE(updateddate,'3000/01/01'),COALESCE(updatedby,'NULL') FROM mst_satuan WHERE satuanid = $1"
	satuan := Satuan{}
	errx = db.QueryRow(selectSatuan, id).Scan(&satuan.SatuanId, &satuan.NamaSatuan, &satuan.InDate, &satuan.InBy, &satuan.UpdatedBy, &satuan.UpdatedDate)
	if errx == nil {
		result = true
	}
	return result
}

func insert(satuan Satuan, tx *sql.Tx) {
	insert := "INSERT INTO mst_satuan (satuanid,nama,indate,inby) VALUES ($1,$2,$3,$4)"
	_, err := tx.Exec(insert, satuan.SatuanId, satuan.NamaSatuan, satuan.InDate, satuan.InBy)
	validate(err, "Insert Satuan", tx)
}

func update(satuan Satuan, tx *sql.Tx) {
	updateStatement := "UPDATE mst_satuan SET nama = $2, UpdatedDate = $3, UpdatedBy = $4 WHERE satuanid = $1;"
	_, err := tx.Exec(updateStatement, satuan.SatuanId, satuan.NamaSatuan, satuan.UpdatedDate, satuan.UpdatedBy)
	validate(err, "Update Satuan", tx)
}

func delete(satuanid string, tx *sql.Tx) {
	deleteStatement := "DELETE FROM mst_satuan WHERE satuanid = $1;"
	_, err := tx.Exec(deleteStatement, satuanid)
	validate(err, "Delete Satuan", tx)
}

func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println(err, "Transaction Rollback")
	} else {
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println("Successfully " + message + " data !")
	}

}

func scanSatuan(rows *sql.Rows) []Satuan {
	xsatuan := []Satuan{}
	var err error

	for rows.Next() {
		satuan := Satuan{}
		err := rows.Scan(&satuan.SatuanId, &satuan.NamaSatuan, &satuan.InDate, &satuan.InBy, &satuan.UpdatedDate, &satuan.UpdatedBy)

		if err != nil {
			panic(err)
		}

		xsatuan = append(xsatuan, satuan)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return xsatuan
}
