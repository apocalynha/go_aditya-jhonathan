# Soal Explorasi
## Palindrome

```
package main

import "fmt"

func main() {
	var input string
	fmt.Print("Apakah Palindrome?\nmasukan kata: ")
	fmt.Scanln(&input)
	fmt.Printf("capture: %s\n", input)

	if Palindrome(input) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}
}

func Palindrome(str string) bool {
	lastIdx := len(str) - 1
	// using for loop
	for i := 0; i < lastIdx/2 && i < (lastIdx-i); i++ {
	   if str[i] != str[lastIdx-i] {
		  return false
	   }
	}
	return true
 }
```

Output

![e_no1](/02_Basic-Program/screenshots/e_no1.jpg)