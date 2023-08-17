package main

import "fmt"

func main() {

	var nil int
	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&nil)

	if nil%2 == 0 {
		fmt.Println(nil, "adalah bilangan genap.")
	} else {
		fmt.Println(nil, "adalah bilangan ganjil.")
	}
}