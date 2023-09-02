# Soal Prioritas 1
## Bilangan Binner

```
package main

import "fmt"

func countBits(n int) []string {
    result := make([]string, n+1)
    for i := 0; i <= n; i++ {
        result[i] = fmt.Sprintf("%b", i)
    }
    return result
}

func main() {
    n1 := 2
    fmt.Println(countBits(n1)) // [0 1 10]

    n2 := 3
    fmt.Println(countBits(n2)) // [0 1 10 11]
}
```

Output

![p1_no1](/08_ProblemSolvingParadigm-BruteForce-Greedy-and-DynamicProgramming/screenshots/p1_no1.JPG)

## Segitiga Pascal

```
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
```

Output

![p1_no2](/08_ProblemSolvingParadigm-BruteForce-Greedy-and-DynamicProgramming/screenshots/p1_no2.JPG)

## Fibonacci Top Down

```
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
```

Output

![p1_no3](/08_ProblemSolvingParadigm-BruteForce-Greedy-and-DynamicProgramming/screenshots/p1_no3.JPG)

## Persamaan

```
package main

import "fmt"

func SimpleEquations(a, b, c int) {
	for x := 1; x <= a; x++ {
		for y := 1; y <= a; y++ {
			for z := 1; z <= a; z++ {
				if x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
					fmt.Println(x, y, z)
					return
				}
			}
		}
	}
	fmt.Println("No solution")
}

func main() {
	SimpleEquations(1, 2, 3)  // No solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
```

Output

![p1_no4](/08_ProblemSolvingParadigm-BruteForce-Greedy-and-DynamicProgramming/screenshots/p1_no4.JPG)