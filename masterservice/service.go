package masterservice

import (
	"bufio"
	"challenge-godb/connection"
	"challenge-godb/getdate"
	"challenge-godb/mastersatuan"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Service struct {
	ServiceId   int
	NamaService string
	Satuan      string
	Harga       int
	InDate      string
	InBy        string
	UpdatedDate string
	UpdatedBy   string
}

func ShowAllService() []Service {
	db, err := sql.Open("postgres", connection.Psqlinfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	sqlStatement := "SELECT serviceid,nama,harga,satuan,indate,inby,COALESCE(updateddate,'3000/01/01 00:00:00') AS updateddate,COALESCE(updatedby,'NULL') AS updatedby FROM mst_service;"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	xservice := scanService(rows)
	return xservice
}

func InsertService() bool {
	var service Service
	var res bool = false
	db := connection.ConnectDB()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	service.ServiceId = inputServiceId()
	service.NamaService = inputNamaService()
	service.Harga = inputHarga()
	service.Satuan = inputSatuan()
	service.InDate = getdate.TimeNow()
	service.InBy = inputUserby()

	insert(service, tx)

	err = tx.Commit()
	if err != nil {
		res = false
		// panic(err)
	} else {
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println("Data Tersimpan")
		res = true
	}
	return res
}

func UpdateService() bool {
	var service Service
	var res bool = false
	xservice := inputServiceIdNotExist()
	if len(strconv.Itoa(xservice)) > 0 {
		if CheckServiceById(xservice) {
			db := connection.ConnectDB()
			defer db.Close()
			tx, err := db.Begin()
			if err != nil {
				panic(err)
			}

			service.ServiceId = xservice
			service.NamaService = inputNamaService()
			service.Harga = inputHarga()
			service.Satuan = inputSatuan()
			service.UpdatedDate = getdate.TimeNow()
			service.UpdatedBy = updateUseby()

			update(service, tx)
			err = tx.Commit()
			if err != nil {
				res = false
				// panic(err)
			} else {
				fmt.Println(strings.Repeat("-", 30))
				fmt.Println("Update Tersimpan")
				res = true
			}
		} else {
			fmt.Println("Satuan Tidak Ada")
			res = false

		}
	}
	return res
}

func DeleteService() bool {
	var res bool = false
	serviceid := inputServiceIdNotExist()
	if CheckServiceById(serviceid) {
		db := connection.ConnectDB()
		defer db.Close()
		tx, err := db.Begin()
		if err != nil {
			panic(err)
		}
		delete(serviceid, tx)
		err = tx.Commit()
		if err != nil {
			panic(err)
		} else {
			res = true
		}
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println("Service Terhapus")
	} else {
		fmt.Println("Service Tidak Ada")
		res = false
	}
	return res
}

func CheckServiceById(id int) bool {
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
	selectService := "SELECT serviceid,nama,harga,satuan,indate,inby,COALESCE(updateddate,'3000/01/01 00:00:00') AS updateddate,COALESCE(updatedby,'NULL') AS updatedby FROM mst_service WHERE serviceid = $1"
	service := Service{}
	errx = db.QueryRow(selectService, id).Scan(&service.ServiceId, &service.NamaService, &service.Harga, &service.Satuan, &service.InDate, &service.InBy, &service.UpdatedBy, &service.UpdatedDate)
	if errx == nil {
		result = true
	}
	return result
}

func inputServiceId() int {
	scanner := bufio.NewScanner(os.Stdin)
	var serviceid int
	loop := 1
	for i := 0; i < loop; i++ {
		fmt.Print("Masukan Id Service : ")
		scanner.Scan()
		str_service := scanner.Text()

		if len(str_service) > 0 {
			serv, err := strconv.Atoi(str_service)
			if err != nil {
				fmt.Println("Input Id Service Harus Angka!, Silahkan Coba Lagi")
				loop++
			} else {
				if CheckServiceById(serv) {
					fmt.Println("Service Id Sudah Ada!, Silahkan Coba Lagi")
					loop++
				} else {
					serviceid = serv
				}
			}
		} else {
			fmt.Println("Input Id Service Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return serviceid
}

func inputServiceIdNotExist() int {
	scanner := bufio.NewScanner(os.Stdin)
	var serviceid int
	loop := 1
	for i := 0; i < loop; i++ {
		fmt.Print("Masukan Id Service : ")
		scanner.Scan()
		str_service := scanner.Text()

		if len(str_service) > 0 {
			serv, err := strconv.Atoi(str_service)
			if err != nil {
				fmt.Println("Input Id Service Harus Angka!, Silahkan Coba Lagi")
				loop++
			} else {
				serviceid = serv
			}
		} else {
			fmt.Println("Input Id Satuan Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return serviceid
}

func inputNamaService() string {
	scanner := bufio.NewScanner(os.Stdin)
	var namaService string
	loop := 1
	for i := 0; i < loop; i++ {
		fmt.Print("Masukan Nama Service : ")
		scanner.Scan()
		if len(scanner.Text()) > 0 {
			namaService = scanner.Text()
		} else {
			fmt.Println("Input Nama Service Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return namaService
}

func inputSatuan() string {
	scanner := bufio.NewScanner(os.Stdin)
	var satuanid string
	loop := 1
	for i := 0; i < loop; i++ {
		fmt.Print("Masukan Satuan : ")
		scanner.Scan()
		if len(scanner.Text()) > 0 {
			satuanid = scanner.Text()
			if !mastersatuan.CheckSatuanById(satuanid) {
				fmt.Println("Satuan Id Tidak Ada!, Silahkan Coba Lagi")
				loop++
			}
		} else {
			fmt.Println("Input Satuan Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return satuanid
}

func inputHarga() int {
	scanner := bufio.NewScanner(os.Stdin)
	var hargaService int
	loop := 1
	for i := 0; i < loop; i++ {
		fmt.Print("Masukan Harga Service : ")
		scanner.Scan()
		str_service := scanner.Text()

		if len(str_service) > 0 {
			serv, err := strconv.Atoi(str_service)
			if err != nil {
				fmt.Println("Input Harga Service Harus Angka!, Silahkan Coba Lagi")
				loop++
			} else {
				hargaService = serv
			}
		} else {
			fmt.Println("Input Harga Service Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return hargaService
}

func inputUserby() string {
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

func insert(service Service, tx *sql.Tx) {
	insert := "INSERT INTO mst_service (serviceid,nama,satuan,harga,indate,inby) VALUES ($1,$2,$3,$4,$5,$6)"
	_, err := tx.Exec(insert, service.ServiceId, service.NamaService, service.Satuan, service.Harga, service.InDate, service.InBy)
	validate(err, "Insert Service", tx)
}

func update(service Service, tx *sql.Tx) {
	updateStatement := "UPDATE mst_service SET nama = $2,satuan = $3, harga = $4, updateddate = $5, updatedby = $6 WHERE serviceid = $1;"
	_, err := tx.Exec(updateStatement, service.ServiceId, service.NamaService, service.Satuan, service.Harga, service.UpdatedDate, service.UpdatedBy)
	validate(err, "Update Service", tx)
}

func delete(serviceid int, tx *sql.Tx) {
	deleteStatement := "DELETE FROM mst_service WHERE serviceid = $1;"
	_, err := tx.Exec(deleteStatement, serviceid)
	validate(err, "Delete Service", tx)
}

func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "Transaction Rollback")
	} else {
		fmt.Println("Successfully " + message + " data !")
	}

}

func scanService(rows *sql.Rows) []Service {
	xservice := []Service{}
	var err error

	for rows.Next() {
		service := Service{}
		err := rows.Scan(&service.ServiceId, &service.NamaService, &service.Harga, &service.Satuan, &service.InDate, &service.InBy, &service.UpdatedDate, &service.UpdatedBy)

		if err != nil {
			panic(err)
		}

		xservice = append(xservice, service)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return xservice
}
