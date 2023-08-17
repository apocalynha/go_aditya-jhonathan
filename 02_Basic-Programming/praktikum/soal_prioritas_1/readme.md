# Soal Prioritas 1
## Luas Trapesium

```
package main

import "fmt"

func main() {

	var AA, AB, tinggi float64
	fmt.Print("Masukkan alas atas trapesium : ")
	fmt.Scanln(&AA)
	fmt.Print("Masukkan alas bawah trapesium : ")
	fmt.Scanln(&AB)
	fmt.Print("Masukkan tinggi trapesium : ")
	fmt.Scanln(&tinggi)

	luas := 0.5 * tinggi * (AA + AB)

	fmt.Println("Luas trapesium adalah : ", luas)
}
```

Output

![p1_no1](/02_Basic-Programming/screenshots/p1_no1.jpg)

## Ganjil Genap
```
package main

import "fmt"

func main() {

	var nil int
	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&nil)

	if nil%2 == 0 {
		fmt.Println(nil, "adalah bilangan genap.")
	} else {
		fmt.Println(nil, "adalah bilangan ganjil.")
	}
}
```

Output

![p1_no2](/02_Basic-Programming/screenshots/p1_no2.jpg)

## Grade
```
package main

import "fmt"

func main() {

	var nilai int
	fmt.Print("Masukkan nilai : ")
	fmt.Scanln(&nilai)

	var grade string
	if nilai >= 80 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 65 && nilai <= 79 {
		grade = "B"
	} else if nilai >= 50 && nilai <= 64 {
		grade = "C"
	} else if nilai >= 35 && nilai <= 49 {
		grade = "D"
	} else if nilai >= 0 && nilai <= 34 {
		grade = "E"
	} else {
		grade = "Nilai Invalid"
	}

	fmt.Println("Grade : ", grade)
}
```

Output

![p1_no3](/02_Basic-Programming/screenshots/p1_no3.jpg)

## Fizz Buzz
```
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(i, " ")
		}
	}
}
```

Output

![p1_no4](/02_Basic-Programming/screenshots/p1_no4.jpg)