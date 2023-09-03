# Soal Prioritas 1
## Goroutine 3 Detik

```
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
```

Output

![p1_no1](/08_ProblemSolvingParadigm-BruteForce-Greedy-and-DynamicProgramming/screenshots/p1_no1.JPG)

## Kelipatan 3 dengan Goroutine dan Channel

```
package main

import (
    "fmt"
	"time"
)

func multiplesThree(x chan int) {
    for i := 1; i < 20; i++{
        if i%3 == 0{
        	x <- i
        }
		time.Sleep(3 * time.Second)
    } 
	close(x)
}

func main() {
	x := make(chan int)
    go multiplesThree(x)
	for num := range x{
		fmt.Println(num)
	}
}
```

Output

![p1_no2](/08_ProblemSolvingParadigm-BruteForce-Greedy-and-DynamicProgramming/screenshots/p1_no2.JPG)

## Kelipatan 3 dengan Goroutine dan Buffer Channel

```
package main

import (
    "fmt"
	"time"
)

func multiplesThree(x chan int) {
    for i := 1; i < 20; i++{
        if i%3 == 0{
        	x <- i
        }
		time.Sleep(3 * time.Second)
    } 
	close(x)
}

func main() {
	x := make(chan int, 1)
    go multiplesThree(x)
	for num := range x{
		fmt.Println(num)
	}
}
```

Output

![p1_no3](/08_ProblemSolvingParadigm-BruteForce-Greedy-and-DynamicProgramming/screenshots/p1_no3.JPG)

## Kelipatan 3 dengan Mutex

```
package main

import (
	"fmt"
	"sync"
	"time"
)

func multiplesThree(x chan int, mu *sync.Mutex) {
	for i := 1; i < 20; i++ {
		if i%3 == 0 {
			mu.Lock()         // Mengunci mutex sebelum menambahkan angka ke dalam channel
			x <- i
			mu.Unlock()       // Membuka kunci mutex setelah selesai
		}
		time.Sleep(3 * time.Second)
	}
	close(x)
}

func main() {
	x := make(chan int, 1)
	var mu sync.Mutex // Mutex untuk mengamankan akses ke channel

	go multiplesThree(x, &mu)

	for num := range x {
		fmt.Println(num)
	}
}
```

Output

![p1_no4](/08_ProblemSolvingParadigm-BruteForce-Greedy-and-DynamicProgramming/screenshots/p1_no4.JPG)