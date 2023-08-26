package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	// Membuat map dengan key 'rune' dan value 'integer' untuk melacak jumlah kemunculan setiap karakter.
	countMap := make(map[rune]int)
	// Membuat slice 'result' untuk menyimpan karakter yang muncul hanya sekali.
	var result []int
	
	// Melakukan iterasi pada setiap karakter dalam string 'angka'.
	for _, char := range angka {
		// Menambah jumlah kemunculan karakter dalam "countMap".
		countMap[char]++ 
	}

	// Melakukan iterasi pada setiap karakter dalam string 'angka' lagi.
	for _, char := range angka {
		// Jika jumlah kemunculan karakter hanya 1, tambahkan nilai numerik karakter ke dalam slice 'result'.
		if countMap[char] == 1 {
			// Menambah nilai numerik karakter ke dalam slice.
			result = append(result, int(char-'0'))
			// Mengubah jumlah kemunculan karakter menjadi 0 untuk menghindari penghitungan ganda.
			countMap[char] = 0 
		}
	}
	return result // Mengembalikan slice yang berisi nilai numerik dari angka yang muncul hanya sekali.
}

func main() {
	// Test cases
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}