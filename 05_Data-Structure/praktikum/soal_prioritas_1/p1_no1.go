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
