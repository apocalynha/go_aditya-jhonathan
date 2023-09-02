# Soal Prioritas 1
## Fibonacci

```
package main

import "fmt"

func fibonacci(number int) int {
	if number <= 1 {
		return number
	}
	return fibonacci(number-1) + fibonacci(number-2)
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}
```

Output

![p1_no1](/07_Recursive-Number-Theory-Sorting-Searching/screenshots/p1_no1.JPG)

## Kemunculan Barang

```
package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	// Membuat pemetaan untuk menghitung jumlah kemunculan barang
	countMap := make(map[string]int)

	for _, item := range items {
		countMap[item]++
	}

	// Mengonversi pemetaan ke dalam bentuk slice
	var pairs []pair
	for name, count := range countMap {
		pairs = append(pairs, pair{name, count})
	}

	// Mengurutkan slice berdasarkan jumlah kemunculan (count) secara menurun
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count < pairs[j].count
	})

	return pairs
}

func main() {
	// golang->1 ruby->2 js->4
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) 
	for _, list := range pairs {
		fmt.Print(list.name, "->", list.count, " ")
	}
	fmt.Println()

	// C->1 D->2 B->3 A->4
	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})
	for _, list := range pairs {
		fmt.Print(list.name, "->", list.count, " ")
	}
	fmt.Println()

	// football->1 basketball->1 tenis->1
	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) 
	for _, list := range pairs {
		fmt.Print(list.name, "->", list.count, " ")
	}
	fmt.Println()
}
```

Output

![p1_no2](/07_Recursive-Number-Theory-Sorting-Searching/screenshots/p1_no2.JPG)

## Bilangan Prima

```
package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	i := 5
	for i*i <= num {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func primeX(number int) int {
	count := 0
	i := 2
	for {
		if isPrime(i) {
			count++
		}
		if count == number {
			return i
		}
		i++
	}
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
```

Output

![p1_no3](/07_Recursive-Number-Theory-Sorting-Searching/screenshots/p1_no3.JPG)