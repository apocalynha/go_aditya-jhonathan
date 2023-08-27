package main

import (
	"fmt"
)

func Compare(a, b string) string {
	// Mengambil panjang string terpendek dari a dan b
	minLength := len(a)
	if len(b) < minLength {
		minLength = len(b)
	}

	// Inisialisasi variabel untuk menyimpan substring yang sama
	var sameSubstring string

	// Melakukan looping untuk mencari substring yang sama dari string terpendek
	for i := 0; i < minLength; i++ {
		// Jika bagian string dari a dan b sama maka akan dimasukkan ke sameSubstring
		if a[i] == b[i] {
			sameSubstring += string(a[i])
		} else {
			break // Keluar dari loop jika karakter berbeda
		}
	}
	return sameSubstring
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
