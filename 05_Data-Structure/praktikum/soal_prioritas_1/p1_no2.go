package main

import "fmt"

func Mapping(slice []string) map[string]int {
	// Membuat map dengan key 'string' dan value 'integer' untuk menyimpan jumlah kemunculan setiap string.
	count := make(map[string]int) 

	// melakukan perulangan untuk setiap string dalam slice.
	for _, str := range slice {
		count[str]++ // Menaikkan jumlah kemunculan string dalam 'count'.
	}

	return count
}

func main() {
	// Test cases
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) // map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      // map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                         // map[]
}