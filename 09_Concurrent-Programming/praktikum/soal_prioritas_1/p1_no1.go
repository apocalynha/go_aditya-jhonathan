package main

import (
    "fmt"
    "sync"
    "time"
)

func multiples(x int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 1; i < 20; i++ {
        if i%x == 0 {
            fmt.Println(i)
        }
        time.Sleep(3 * time.Second)
    }
}

func main() {
    var x int
    fmt.Print("Masukkan nilai x : ")
    fmt.Scan(&x)

    var wg sync.WaitGroup
    wg.Add(1)

    go multiples(x, &wg)

    // Menunggu goroutine selesai
    wg.Wait()
}
