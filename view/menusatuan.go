package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MenuSatuan() string {
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
	return input
}
