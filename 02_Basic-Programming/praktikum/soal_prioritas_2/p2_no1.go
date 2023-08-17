package main

import "fmt"

func main() {
	var input int
	fmt.Print("Input : ")
	fmt.Scanln(&input)

	fmt.Println("Output : ")
	for i := 1; i <= input; i++ {
		// Mencetak spasi sebelum asterisk pada setiap baris
		for j := 1; j <= input-i; j++ {
			fmt.Print(" ")
		}
		// Mencetak asterisk pada setiap baris
		for j := 1; j <= 2*i-1; j++ {
			if j%2 == 0 {
				fmt.Print("")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}
