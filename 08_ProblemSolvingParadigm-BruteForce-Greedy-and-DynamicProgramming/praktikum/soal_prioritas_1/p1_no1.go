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
