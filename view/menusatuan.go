package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MenuSatuan() int {
	x := 1
	var in_satuan int
	for i := 0; i < x; i++ {
		scan := bufio.NewScanner(os.Stdin)
		fmt.Println(strings.Repeat("-", 20))
		fmt.Println("=== Menu Satuan ===")
		fmt.Println("1. View Satuan")
		fmt.Println("2. Add Satuan")
		fmt.Println("3. Update Satuan")
		fmt.Println("4. Delete Satuan")
		fmt.Println("9. Kembali")
		fmt.Println("0. Keluar")
		fmt.Println(strings.Repeat("-", 20))
		fmt.Print("Masukan Kode Menu Satuan :")
		scan.Scan()
		input := scan.Text()
		if len(input) > 0 {
			in, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(strings.Repeat("-", 20))
				fmt.Println("Input Harus Angka !")
				x++
			} else {
				in_satuan = in
			}
		} else {
			x++
		}
	}
	return in_satuan
}
