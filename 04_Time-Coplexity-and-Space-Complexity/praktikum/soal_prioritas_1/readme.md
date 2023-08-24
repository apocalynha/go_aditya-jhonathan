# Soal Prioritas
## Bilangan Prima

```
package main

import "fmt"

func primeNumber(n int) bool{	
	// 2,3,5,7,11,13

	if n == 2 {
		return true
	}
	if n % 2 == 0 || n <= 1{
		return false
	}
	for i := 3; i*i <= n; i+=2 { 
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
 }
```

Output

![p1_no1](/02_Basic-Programming/screenshots/p1_no1.JPG)

## Perpangkatan

```
package main

import "fmt"

func pow(x, n int) int {
	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		temp := pow(x, n/2)
		return temp * temp
	}

	temp := pow(x, (n-1)/2)
	return x * temp * temp
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
 }
```

Output

![p1_no2](/02_Basic-Programming/screenshots/p1_no2.JPG)