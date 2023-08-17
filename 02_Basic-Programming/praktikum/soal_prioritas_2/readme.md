# Soal Prioritas 2
## Segitiga Asterik
```
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
```

Output

![p2_no1](/02_Basic-Programming/screenshots/p2_no1.jpg)

## Faktor Bilangan
```
package main

import "fmt"

func main() {
	var angka int
	fmt.Print("angka: ")
	fmt.Scanln(&angka)

	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			fmt.Println(i)
		}
	}
}
```

Output

![p2_no2](/02_Basic-Programming/screenshots/p2_no2.jpg)