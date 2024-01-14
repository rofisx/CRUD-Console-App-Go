package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MenuService() int {
	x := 1
	var in_service int
	for i := 0; i < x; i++ {
		scan := bufio.NewScanner(os.Stdin)
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println("=== Menu Service ===")
		fmt.Println("1. View Service")
		fmt.Println("2. Add Service")
		fmt.Println("3. Update Service")
		fmt.Println("4. Delete Service")
		fmt.Println("9. Kembali")
		fmt.Println("0. Keluar")
		fmt.Println(strings.Repeat("-", 50))
		fmt.Print("Masukan Kode Menu Service :")
		scan.Scan()
		input := scan.Text()
		if len(input) > 0 {
			in, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(strings.Repeat("-", 50))
				fmt.Println("Input Harus Angka !")
				x++
			} else {
				in_service = in
			}
		} else {
			x++
		}
	}
	return in_service
}
