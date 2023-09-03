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
		time.Sleep(0 * time.Second)
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