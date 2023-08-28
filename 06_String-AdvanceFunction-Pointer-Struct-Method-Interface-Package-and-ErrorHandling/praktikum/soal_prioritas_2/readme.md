# Soal Prioritas 1
## Caesar Cipher

```
package main

import "fmt"

func caesar(offset int, input string) string {
	// Variabel untuk menyimpan hasil enkripsi
	var result string
	// Perulangan melalui setiap karakter dalam input
	for _, char := range input{
		// Periksa apakah karakter adalah huruf kecil
		if char >= 'a' && char <= 'z' {
			// Menghitung karakter baru setelah pergeseran dengan menggunakan operator modulo
			newChar := rune(((int(char)-'a'+offset)%26)+'a')
			// Menambahkan karakter baru ke hasil enkripsi
			result += string(newChar)
		// Periksa apakah karakter adalah huruf besar
		} else if char >= 'A' && char <= 'Z' {
			// Menghitung karakter baru untuk huruf besar
			newChar := rune(((int(char)-'A'+offset)%26)+'A')
			// Menambahkan karakter baru ke hasil enkripsi
			result += string(newChar)
		} else {
			// Menambahkan karakter yang bukan huruf ke hasil enkripsi tanpa perubahan
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc")) // def
	fmt.Println(caesar(2, "alta")) // cnvc
	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi 
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza 
  	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl 
}
```

Output

![p2_no1](/06_String-AdvanceFunction-Pointer-Struct-Method-Interface-Package-and-ErrorHandling/screenshots/p2_no1.JPG)