package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MenuKasir() int {
	x := 1
	var in_kasir int
	for i := 0; i < x; i++ {
		scan := bufio.NewScanner(os.Stdin)
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println("=== Menu Kasir Laundry ===")
		fmt.Println("1. View Transaski Laundry")
		fmt.Println("2. Kasir Laundry")
		fmt.Println("9. Kembali")
		fmt.Println("0. Keluar")

		fmt.Println(strings.Repeat("-", 50))
		fmt.Print("Masukan Kode Menu Kasir :")
		scan.Scan()
		input := scan.Text()
		if len(input) > 0 {
			in, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(strings.Repeat("-", 50))
				fmt.Println("Input Harus Angka !")
				x++
			} else {
				in_kasir = in
			}
		} else {
			fmt.Println(strings.Repeat("-", 50))
			fmt.Println("Input Kode Menu Transaksi Kosong !")
			x++
		}
	}
	return in_kasir
}
