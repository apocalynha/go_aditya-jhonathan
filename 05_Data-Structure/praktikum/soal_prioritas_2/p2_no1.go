package main

import "fmt"

func PairSum(arr []int, target int) []int {
	// menginisialisasi dua variabel indeks yaitu left, right, dan panjang indeks array-1
	left, right := 0, len(arr)-1 

	// melakukan perulangan selama indeks kiri kurang dari indeks kanan.
	for left < right { 
		// Hitung jumlah elemen di kedua indeks.
		sum := arr[left] + arr[right] 

		// Jika jumlah sesuai dengan target, kembalikan indeks kedua indeks.
		if sum == target {
			return []int{left, right} 
		// Jika jumlah kurang dari target, naikkan indeks kiri untuk menambahkan indeks yang lebih besar.
		} else if sum < target {
			left++ 
		// Jika jumlah lebih dari target, turunkan indeks kanan untuk mencari indeks yang lebih kecil.
		} else {
			right-- 
		}
	}

	return []int{} // Jika tidak ditemukan pasangan yang sesuai, kembalikan slice kosong.
}

func main() {
	// Test cases
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}
