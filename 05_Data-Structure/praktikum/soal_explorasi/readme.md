# Soal Eksplorasi
## Matriks 2x2

```
package main

import "fmt"

func main() {
	// Membuat matriks 2x2
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	// Mendeklarasikan variabel untuk masing-masing diagonal
	LeftToRight := 0
	RightToLeft := 0

	size := len(matrix)

	for i := 0; i < size; i++ {
		// Menghitung jumlah elemen diagonal dari kiri ke kanan.
		LeftToRight += matrix[i][i]       
		// Menghitung jumlah elemen diagonal dari kanan ke kiri.
		RightToLeft += matrix[i][size-i-1] 
	}

	result := LeftToRight - RightToLeft

	fmt.Println("Diagonal kiri ke kanan : ", LeftToRight)
	fmt.Println("Diagonal kanan ke kiri : ", RightToLeft)
	fmt.Printf("Perbedaan mutlak %d - %d : %d\n", LeftToRight, RightToLeft, result)
}
```

Output

![e_no1](/05_Data-Structure/screenshots/e_no1.JPG)