package main

import "fmt"

var memoTopDown = map[int]int{}

func fibTopDown(i int) int {
	nilai, isCount := memoTopDown[i]
	if isCount {
		return nilai
	}

	if i <= 1 {
		memoTopDown[i] = i
	} else {
		memoTopDown[i] = fibTopDown(i-1) + fibTopDown(i-2)
	}
	return memoTopDown[i]
}

func main() {
	fmt.Println(fibTopDown(0))  // 0
	fmt.Println(fibTopDown(2))  // 1
	fmt.Println(fibTopDown(9))  // 34
	fmt.Println(fibTopDown(10)) // 55
	fmt.Println(fibTopDown(12)) // 144
}
