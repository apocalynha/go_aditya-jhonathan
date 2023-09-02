package main

import "fmt"

func Frog(jumps []int) int {
	n := len(jumps)
	if n < 2 {
		return 0 // Tidak ada cost jika hanya ada satu atau kurang Batu.
	}

	// Buat sebuah slice untuk menyimpan cost minimum pada setiap Batu.
	minCost := make([]int, n)

	// Cost di Batu pertama adalah nol.
	minCost[0] = 0

	// Cost di Batu kedua adalah biaya loncatan ke Batu kedua.
	minCost[1] = abs(jumps[1] - jumps[0])

	// Iterasi dari Batu ketiga hingga Batu terakhir.
	for i := 2; i < n; i++ {
		// Pilihan cost minimum antara loncatan ke Batu ke-(i-1) dan Batu ke-(i-2).
		minCost[i] = min(minCost[i-1]+abs(jumps[i]-jumps[i-1]), minCost[i-2]+abs(jumps[i]-jumps[i-2]))
	}

	// Hasil akhir adalah cost minimum di Batu terakhir.
	return minCost[n-1]
}

// nilai absolut atau non-negatif
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// menemukan cost terkecil
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
