# Soal Explorasi
## Maksimum Penjumlahan

```
package main

import (
	"fmt"
)

func MaxSequence(arr []int) int {
	maxSum := 0
	currentSum := 0

	for _, num := range arr {
		currentSum += num
		if currentSum < 0 {
			currentSum = 0
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}
```

Output

![e_no1](/07_Recursive-Number-Theory-Sorting-Searching/screenshots/e_no1.JPG)

## Membeli Barang

```
package main

import (
	"fmt"
	"sort"
)

func MaximumBuyProduct(money int, productPrice []int) int {
	// Sort harga produk dari yang termurah
	sort.Ints(productPrice)

	buyCount := 0
	for _, price := range productPrice {
		if money >= price {
			money -= price
			buyCount++
		} else {
			break
		}
	}

	return buyCount
}

func main() {
	fmt.Println(MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000}))      // 3
	fmt.Println(MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})) // 4
	fmt.Println(MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000}))   // 4
	fmt.Println(MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000}))           // 1
	fmt.Println(MaximumBuyProduct(0, []int{10000, 30000}))                        // 0
}
```

Output

![e_no2](/07_Recursive-Number-Theory-Sorting-Searching/screenshots/e_no2.JPG)