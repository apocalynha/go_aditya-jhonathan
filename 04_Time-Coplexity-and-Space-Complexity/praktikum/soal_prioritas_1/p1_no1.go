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