package main

import (
    "fmt"
    "strings"
    "sync"
)

func calculateFrequency(text string, ch chan rune, wg *sync.WaitGroup) {
    defer wg.Done()

    for _, char := range text {
        ch <- char
    }
    close(ch)
}

func countFrequency(ch chan rune, freq map[rune]int, wg *sync.WaitGroup) {
    defer wg.Done()

    for char := range ch {
        freq[char]++
    }
}

func main() {
    text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

    text = strings.ReplaceAll(strings.ToLower(text), " ", "")

    ch := make(chan rune)
    var wg sync.WaitGroup
    wg.Add(2)

    freq := make(map[rune]int)

    go calculateFrequency(text, ch, &wg)
    go countFrequency(ch, freq, &wg)

    wg.Wait()

    for char, count := range freq {
        fmt.Printf("%c : %d\n", char, count)
    }
}