package main

import "fmt"

func main() {

	var AA, AB, tinggi float64
	fmt.Print("Masukkan alas atas trapesium : ")
	fmt.Scanln(&AA)
	fmt.Print("Masukkan alas bawah trapesium : ")
	fmt.Scanln(&AB)
	fmt.Print("Masukkan tinggi trapesium : ")
	fmt.Scanln(&tinggi)

	luas := 0.5 * tinggi * (AA + AB)

	fmt.Println("Luas trapesium adalah : ", luas)
}