# Soal Prioritas 1
## Car

```
package main

import (
	"fmt"
)

// Definisikan struct Car
type Car struct {
	carType string
	fuelin  float64
}

// Method untuk menghitung perkiraan jarak
func (c Car) Estimate() float64 {
	// Menghitung jarak 
	distance := c.fuelin * 1.5 
	return distance
}

func main() {
	// Membuat instance Car
	car := Car{
		carType: "SUV",
		fuelin:  10.5,
	}

	// Mencetak informasi tipe mobil dan perkiraan jarak
	fmt.Printf("car type: %s , est. distance: %.2f\n", car.carType, car.Estimate())
}
```

Output

![p1_no1](/06_String-AdvanceFunction-Pointer-Struct-Method-Interface-Package-and-ErrorHandling/screenshots/p1_no1.JPG)

## Siswa dan Skor

```
package main

import (
	"fmt"
)

// Mendefinisikan struct Studen berisi array string name dan array integer score
type Student struct {
	name  []string
	score []int
}

// Mendefinisikana method average pada struct penerima s Student
func (s Student) Average() float64 {
	totalScore := 0
	// Menghitung total skor dari banyaknya score
	for _, score := range s.score {
		totalScore += score
	}
	// Menghitung rata-rata skor dari total jumlah score dan banyaknya score
	return float64(totalScore) / float64(len(s.score))
}

// Mendefinisikan method min pada struct penerima s Student
func (s Student) Min() (min int, name string) {
	// Tempat menyimpan nilai minimum dan nama siswanya
	min = s.score[0]
	name = s.name[0]
	// Mencari nilai minimum dan nama siswa dengan skor terendah
	for i := 1; i < len(s.score); i++ {
		// Menyimpan nilai minimum dan siswa ke variabel min dan name
		if s.score[i] < min {
			min = s.score[i]
			name = s.name[i]
		}
	}
	return min, name
}

// Mendefinisikana method min pada struct penerima s Student
func (s Student) Max() (max int, name string) {
	// Tempat menyimpan nilai maksimum dan nama siswanya
	max = s.score[0]
	name = s.name[0]
	// Mencari nilai maksimum dan nama siswa dengan skor tertinggi
	for i := 1; i < len(s.score); i++ {
		// Menyimpan nilai minimum dan siswa ke variabel min dan name
		if s.score[i] > max {
			max = s.score[i]
			name = s.name[i]
		}
	}
	return max, name
}

func main() {
	var a = Student{}

	for i := 1; i <= 5; i++ {
		var name string
		fmt.Print("Input ", i, " Student's Name ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input ", i, " Student's Score ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\nAverage Score Students is", a.Average())

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score of Students : ",nameMin,"(",scoreMin,")")

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score of Students : ",nameMax,"(",scoreMax,")")
}
```

Output

![p1_no2](/06_String-AdvanceFunction-Pointer-Struct-Method-Interface-Package-and-ErrorHandling/screenshots/p1_no2.JPG)

## Substring

```
package main

import (
	"fmt"
)

func Compare(a, b string) string {
	// Mengambil panjang string terpendek dari a dan b
	minLength := len(a)
	if len(b) < minLength {
		minLength = len(b)
	}

	// Inisialisasi variabel untuk menyimpan substring yang sama
	var sameSubstring string

	// Melakukan looping untuk mencari substring yang sama dari string terpendek
	for i := 0; i < minLength; i++ {
		// Jika bagian string dari a dan b sama maka akan dimasukkan ke sameSubstring
		if a[i] == b[i] {
			sameSubstring += string(a[i])
		} else {
			break // Keluar dari loop jika karakter berbeda
		}
	}
	return sameSubstring
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
```

Output

![p1_no3](/06_String-AdvanceFunction-Pointer-Struct-Method-Interface-Package-and-ErrorHandling/screenshots/p1_no3.JPG)

## Pointer Nilai Maksimum dan Minimum

```
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
```

Output

![p1_no4](/06_String-AdvanceFunction-Pointer-Struct-Method-Interface-Package-and-ErrorHandling/screenshots/p1_no4.JPG)