# Soal Prioritas 1
## Merge String

```
package main

import "fmt"

// Deklarasikan fungsi ArrayMerge dengan parameter arrayA dan arrayB berupa slice string 
func ArrayMerge(arrayA, arrayB []string) []string {
	
	// membuat variabel result untuk membuat slice kosong yang berisi hasil gabungan dengan kapasitas 
	result := make([]string, 0, len(arrayA)+len(arrayB))
	
	// membuat map dengan key string dan value boolean untuk menyimpan nama-nama
	listNames := make(map[string]bool)

	// melakukan iterasi melalui arrayA dan arrayB yang telah digabungkan
	for _, name := range append(arrayA, arrayB...) {
		
		// memeriksa setiap isi dari listNames apakah "name" sudah ada atau tidak
		if !listNames[name] { // menandakan bahwa "name" ini belum ada
			listNames[name] = true // kemudian mengubah nilai "name" ke true yang berarti sudah ada didalam listNames
			result = append(result, name) // Menambahkan nama ke slice "result"
		}
	}

	return result
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}

```

Output

![p1_no1](/05_Data-Structure/screenshots/p1_no1.JPG)

## Menghitung String

```
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
```

Output

![p1_no2](/05_Data-Structure/screenshots/p1_no2.JPG)

## Muncul Sekali

```
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
```

Output

![p1_no3](/05_Data-Structure/screenshots/p1_no3.JPG)