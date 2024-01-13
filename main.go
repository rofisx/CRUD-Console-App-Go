package main

import (
	"bufio"
	"challenge-godb/kasirlaundry"
	"challenge-godb/mastersatuan"
	"challenge-godb/masterservice"
	"challenge-godb/view"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	var loop int = 1
	for i := 0; i < loop; i++ {
		input := MenuUtama()
		if len(input) > 0 {
			in, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Input Harus Angka !")
				loop++
				continue
			}
			switch in {
			case 1:
				if MenuSatuan() == 9 {
					loop++
				}
			case 2:
				if MenuService() == 9 {
					loop++
				}
			case 3:
				if MenuKasir() == 9 {
					loop++
				}
			case 0:
				os.Exit(0)
			default:
				fmt.Println("Invalid Input")
				loop++
			}
		} else {
			loop++
		}
	}
}

func MenuUtama() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(strings.Repeat("=", 14), "Laundry Enigma", strings.Repeat("=", 14))
	fmt.Println("1. Master Satuan")
	fmt.Println("2. Master Service")
	fmt.Println("3. Kasir Laundry")
	fmt.Println("0. Exit")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Print("Masukan No Menu : ")
	scanner.Scan()
	input := scanner.Text()
	return input
}

func MenuSatuan() int {
	var int_input int
	for {
		input := view.MenuSatuan()
		if len(input) > 0 {
			in, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Input Harus Angka !")
			} else {
				int_input = in
			}

			if int_input == 9 {
				return int_input
			}

			switch in {
			case 1:
				xsatuan := mastersatuan.ShowAllSatuan()
				for _, satuan := range xsatuan {
					fmt.Println(strings.Repeat("=", 30))
					fmt.Println("Satuan Id :", satuan.SatuanId)
					fmt.Println("Nama Satuan :", satuan.NamaSatuan)
					fmt.Println("In Date :", satuan.InDate)
					fmt.Println("In By :", satuan.InBy)
					fmt.Println("Updated Date :", satuan.UpdatedDate)
					fmt.Println("Updated By :", satuan.UpdatedBy)
					fmt.Println(strings.Repeat("=", 30))
				}
			case 2:
				if !mastersatuan.InsertSatuan() {
					fmt.Println("Error Insert Satuan")
				}
			case 3:
				if !mastersatuan.UpdateSatuan() {
					fmt.Println("Error Update Satuan")
				}
			case 4:
				if !mastersatuan.DeleteSatuan() {
					fmt.Println("Gagal Delete Satuan")
				}
			case 0:
				os.Exit(0)
			default:
				fmt.Println("Input Invalid")
			}
		}
	}
}

func MenuService() int {
	var int_input int
	for {
		input := view.MenuService()
		if len(input) > 0 {
			in, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Input Harus Angka !")
			} else {
				int_input = in
			}

			if int_input == 9 {
				return int_input
			}

			switch in {
			case 1:
				xservice := masterservice.ShowAllService()
				for _, serv := range xservice {
					fmt.Println(strings.Repeat("=", 30))
					fmt.Println("Service Id :", serv.ServiceId)
					fmt.Println("Nama Service :", serv.NamaService)
					fmt.Println("Harga Service :", serv.Harga)
					fmt.Println("Satuan  :", serv.Satuan)
					fmt.Println("In Date :", serv.InDate)
					fmt.Println("In By :", serv.InBy)
					fmt.Println("Updated Date :", serv.UpdatedDate)
					fmt.Println("Updated By :", serv.UpdatedBy)
					fmt.Println(strings.Repeat("=", 30))
				}
			case 2:
				if !masterservice.InsertService() {
					fmt.Println("Error Insert Service")
				}
			case 3:
				if !masterservice.UpdateService() {
					fmt.Println("Error Update Service")
				}
			case 4:
				if !masterservice.DeleteService() {
					fmt.Println("Gagal Delete Service")
				}
			case 0:
				os.Exit(0)
			default:
				fmt.Println("Input Invalid")
			}
		}
	}
}

func MenuKasir() int {
	var int_input int
	for {
		input := view.MenuKasir()
		if len(input) > 0 {
			in, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Input Harus Angka !")
			} else {
				int_input = in
			}

			if int_input == 9 {
				return int_input
			}
			switch in {
			case 1:
				xkasir := kasirlaundry.ShowAllTransaksi()
				for _, trs := range xkasir {
					fmt.Println(strings.Repeat("=", 30))
					fmt.Println("Trs Id :", trs.TrsID)
					fmt.Println("Nama Customer :", trs.NamaCustomer)
					fmt.Println("Contact :", trs.Contact)
					fmt.Println("Total Qty :", trs.TotalQty)
					fmt.Println("Total Tagihan :", trs.TotalTagihan)
					fmt.Println("In Date :", trs.InDate)
					fmt.Println("In By :", trs.Inby)
					fmt.Println(strings.Repeat("=", 30))
				}
			case 2:
				if !kasirlaundry.InsertTansaksi() {
					fmt.Println("Data Tidak Tersimpan")
				}
			case 0:
				os.Exit(0)
			default:
				fmt.Println("Invalid Input")
			}
		}
	}
}
