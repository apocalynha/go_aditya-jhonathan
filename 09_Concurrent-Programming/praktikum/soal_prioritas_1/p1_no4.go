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
		time.Sleep(0 * time.Second)
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
