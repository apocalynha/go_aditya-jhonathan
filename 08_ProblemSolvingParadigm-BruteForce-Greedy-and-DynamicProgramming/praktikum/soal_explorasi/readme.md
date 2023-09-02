# Soal Explorasi
## Angka Romawi

```
package main

import "fmt"

// Membuat slice yang berisi pasangan angka Romawi dan nilai desimalnya
var romanNum = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	roman := ""
	for _, pair := range romanNum {
		for num >= pair.Value {
			roman += pair.Symbol
			num -= pair.Value
		}
	}
	return roman
}

func main() {
	fmt.Println(intToRoman(4))     // IV
	fmt.Println(intToRoman(9))     // IX
	fmt.Println(intToRoman(23))    // XXIII
	fmt.Println(intToRoman(2021))  // MMXXI
	fmt.Println(intToRoman(1646))  // MDCXLVI
}
```

Output

![e_no1](/08_ProblemSolvingParadigm-BruteForce-Greedy-and-DynamicProgramming/screenshots/e_no1.JPG)