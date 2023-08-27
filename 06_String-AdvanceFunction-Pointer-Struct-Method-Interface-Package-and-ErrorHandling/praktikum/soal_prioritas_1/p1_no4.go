package main

import "fmt"

func getMinMax(numbers ...*int) (min, max int) {
	// Inisialisasi min dan max dengan merefrensikan ke pointer pada indeks 0 dari slice numbers
    min = *numbers[0]   
    max = *numbers[0]   

    // Melakukan loop melalui slice numbers untuk membandingkan nilai
    for _, num := range numbers {
		
		// menyimpan nilai min jika ditemukan angka yang lebih kecil
        if *num < min {
            min = *num   
        }

		// menyimpan nilai max jika ditemukan angka yang lebih besar
        if *num > max {
            max = *num   
        }
    }

    return min, max   // Mengembalikan nilai min dan max
}

func main() {
    var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Println("Input : ")
    fmt.Scan(&a1)
    fmt.Scan(&a2)
    fmt.Scan(&a3)
    fmt.Scan(&a4)
    fmt.Scan(&a5)
    fmt.Scan(&a6)
    fmt.Println()

    // Memanggil fungsi getMinMax dengan parameter-pointer dari angka
    min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Println("Output : ")
	fmt.Println(max, " is the maximum number")
    fmt.Println(min, " is the minimum number")
}
