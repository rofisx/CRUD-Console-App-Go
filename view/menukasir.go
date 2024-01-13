package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MenuKasir() string {
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
	return input
}
