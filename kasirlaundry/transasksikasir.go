package kasirlaundry

import (
	"bufio"
	"challenge-godb/connection"
	"challenge-godb/getdate"
	"challenge-godb/masterservice"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TrsHeaderLaundry struct {
	TrsID        string
	NamaCustomer string
	Contact      int
	TotalQty     int
	TotalTagihan int
	InDate       string
	Inby         string
}

type TrsDetailLundry struct {
	TrsID     string
	ServiceId int
	Harga     int
	Qty       int
}

func ShowAllTransaksi() []TrsHeaderLaundry {
	db, err := sql.Open("postgres", connection.Psqlinfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	sqlStatement := "SELECT trsid,customer,contact,totalqty,totaltagihan,indate,inby FROM trx_laundry;"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	xtransaksi := scanTransaksi(rows)
	if len(xtransaksi) <= 0 {
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println("Transaksi Kosong, Silahkan Input Transaksi")
		fmt.Println(strings.Repeat("-", 50))
	}
	return xtransaksi
}

func InsertTansaksi() bool {
	var status bool = false
	var trs TrsHeaderLaundry
	var detailtrs TrsDetailLundry
	trs.TrsID = getTrs()
	trs.NamaCustomer = inputCustomer()
	trs.Contact = inputContact()
	x := 1
	servid := []int{}
	qty := []int{}
	for i := 0; i < x; i++ {
		servid = append(servid, inputServiceId())
		qty = append(qty, inputQty())
		if answerDetail() == "y" {
			x++
		} else {
			trs.Inby = inputUserby()
			trs.TotalQty = sumQty(qty)
			totalHarga := 0
			ix := 0
			for _, val := range servid {
				if masterservice.CheckServiceById(val) {
					totalHarga += getHarga(val) * qty[ix]
					ix++
				}
			}
			trs.TotalTagihan = totalHarga
			trs.InDate = getdate.TimeNow()

			db := connection.ConnectDB()
			defer db.Close()
			tx, err := db.Begin()
			if err != nil {
				panic(err)
			}

			insert(trs, tx)

			for i := 0; i < len(servid); i++ {
				detailtrs.TrsID = trs.TrsID
				detailtrs.ServiceId = servid[i]
				// fmt.Println(detailtrs.ServiceId)
				detailtrs.Harga = getHarga(servid[i])
				detailtrs.Qty = qty[i]
				insertDetail(detailtrs, tx)
			}

			err = tx.Commit()
			if err != nil {
				status = false
				// panic(err)
			} else {
				fmt.Println(strings.Repeat("-", 50))
				fmt.Println("Data Tersimpan")
				status = true
			}
		}
	}
	return status
}

func scanTransaksi(rows *sql.Rows) []TrsHeaderLaundry {
	xtrs := []TrsHeaderLaundry{}
	var err error

	for rows.Next() {
		trskasir := TrsHeaderLaundry{}
		err := rows.Scan(&trskasir.TrsID, &trskasir.NamaCustomer, &trskasir.Contact, &trskasir.TotalQty, &trskasir.TotalTagihan, &trskasir.InDate, &trskasir.Inby)

		if err != nil {
			panic(err)
		}
		xtrs = append(xtrs, trskasir)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return xtrs
}

func checkTrsExist(trs string) bool {
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
	selectService := "SELECT trsid FROM trx_laundry WHERE trsid = $1"
	var xtrs string
	errx = db.QueryRow(selectService, trs).Scan(&xtrs)
	if errx == nil {
		result = true
	}
	return result
}

func getTrs() string {
	var result string
	date := getdate.GetYMD()
	var trs string = "TRX" + date + "000001"
	if checkTrsExist(trs) {
		xtrs := trs[11:]
		trsint, _ := strconv.Atoi(xtrs)
		lenindex := len(xtrs) - len(strconv.Itoa(trsint))
		zero := ""
		for i := 0; i < lenindex; i++ {
			zero += "0"
		}
		trsint++
		result = trs[0 : len(trs)-len(xtrs)]
		result += zero + strconv.Itoa(trsint)
	} else {
		result = trs
	}
	return result
}

func inputCustomer() string {
	scanner := bufio.NewScanner(os.Stdin)
	var customer string
	loop := 1
	for i := 0; i < loop; i++ {
		fmt.Print("Masukan Nama Customer : ")
		scanner.Scan()
		if len(scanner.Text()) > 0 {
			customer = scanner.Text()
		} else {
			fmt.Println(strings.Repeat("-", 30))
			fmt.Println("Nama Customer Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return customer
}

func inputContact() int {
	scanner := bufio.NewScanner(os.Stdin)
	var contact int
	loop := 1
	for i := 0; i < loop; i++ {
		fmt.Print("Masukan Contact Customer [+62]:")
		scanner.Scan()
		str_con := scanner.Text()

		if len(str_con) > 0 {
			serv, err := strconv.Atoi(str_con)
			if err != nil {
				fmt.Println(strings.Repeat("-", 30))
				fmt.Println("Contact Harus Angka!, Silahkan Coba Lagi")
				fmt.Println(strings.Repeat("-", 30))
				loop++
			} else {
				if len(str_con) <= 10 {
					fmt.Println(strings.Repeat("-", 30))
					fmt.Println("Minimal 11 Angka!, Silahkan Coba Lagi")
					fmt.Println(strings.Repeat("-", 30))
					loop++
				} else {
					contact = serv
				}
			}
		} else {
			fmt.Println("Input Contact Customer Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return contact
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
			fmt.Println(strings.Repeat("-", 30))
			fmt.Println("Input Nama User Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return userby
}

func getHarga(servid int) int {
	db, err := sql.Open("postgres", connection.Psqlinfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	Harga := 0
	sqlStatement := "SELECT harga FROM mst_service WHERE serviceid = $1;"
	err = db.QueryRow(sqlStatement, servid).Scan(&Harga)
	if err != nil {
		panic(err)
	}
	return Harga
}

func sumQty(qty []int) int {
	res := 0
	for _, val := range qty {
		res += val
	}
	return res
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
				fmt.Println(strings.Repeat("-", 30))
				fmt.Println("Input Id Service Harus Angka!, Silahkan Coba Lagi")
				fmt.Println(strings.Repeat("-", 30))
				loop++
			} else {
				if !masterservice.CheckServiceById(serv) {
					fmt.Println(strings.Repeat("-", 30))
					fmt.Println("Service Id Tidak Ada!, Silahkan Coba Lagi")
					fmt.Println(strings.Repeat("-", 30))
					loop++
				} else {
					serviceid = serv
				}
			}
		} else {
			fmt.Println(strings.Repeat("-", 30))
			fmt.Println("Input Id Service Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return serviceid
}

func inputQty() int {
	scanner := bufio.NewScanner(os.Stdin)
	var qty int
	loop := 1
	for i := 0; i < loop; i++ {
		fmt.Print("Masukan Qty Service : ")
		scanner.Scan()
		str_qty := scanner.Text()

		if len(str_qty) > 0 {
			qtyint, err := strconv.Atoi(str_qty)
			if err != nil {
				fmt.Println(strings.Repeat("-", 30))
				fmt.Println("Input Qty Service Harus Angka Minimal 1!, Silahkan Coba Lagi")
				loop++
			} else {
				if qtyint <= 0 {
					fmt.Println(strings.Repeat("-", 30))
					fmt.Println("Qty Service Minimal 1!, Silahkan Coba Lagi")
					loop++
				} else {
					qty = qtyint
				}
			}
		} else {
			fmt.Println(strings.Repeat("-", 30))
			fmt.Println("Input Qty Service Kosong!, Silahkan Coba Lagi")
			loop++
		}
	}
	return qty
}

func answerDetail() string {
	var result string
	scanner := bufio.NewScanner(os.Stdin)
	x := 1
	for i := 0; i < x; i++ {
		fmt.Print("Tambah Service ? y [Ya] / n [Tidak] :")
		scanner.Scan()
		answer := scanner.Text()
		if answer == "y" {
			result = answer
		} else if answer == "n" {
			result = answer
		} else {
			fmt.Println(strings.Repeat("-", 30))
			fmt.Println("Harap Masukan Sesuai Petunjuk")
			fmt.Println(strings.Repeat("-", 30))
			x++
		}
	}
	return result
}

func insert(headerTrs TrsHeaderLaundry, tx *sql.Tx) {
	insert := "INSERT INTO trx_laundry (trsid,customer,contact,totalqty,totaltagihan,indate,inby) VALUES ($1,$2,$3,$4,$5,$6,$7);"
	_, err := tx.Exec(insert, headerTrs.TrsID, headerTrs.NamaCustomer, headerTrs.Contact,
		headerTrs.TotalQty, headerTrs.TotalTagihan, headerTrs.InDate, headerTrs.Inby)
	validate(err, "Insert Transaksi", tx)
}

func insertDetail(detailTrs TrsDetailLundry, tx *sql.Tx) {
	insert := "INSERT INTO trx_laundry_detail (trsid,serviceid,harga,qty) VALUES ($1,$2,$3,$4);"
	_, err := tx.Exec(insert, detailTrs.TrsID, detailTrs.ServiceId, detailTrs.Harga, detailTrs.Qty)
	validate(err, "Insert Detail Transaksi", tx)
}

func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println(err, "Transaction Rollback")
	} else {
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println("Successfully " + message + " data !")
	}

}
