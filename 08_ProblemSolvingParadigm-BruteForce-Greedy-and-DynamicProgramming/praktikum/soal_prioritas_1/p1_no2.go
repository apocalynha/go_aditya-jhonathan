package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}

	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

func main() {
	var numRows int
	fmt.Print("Input numRows: ")
	fmt.Scan(&numRows)

	triangle := generate(numRows)

	fmt.Print("Output: ")
	for _, row := range triangle {
		fmt.Print(row)
	}
}
