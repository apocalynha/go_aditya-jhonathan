package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	// Membuat pemetaan untuk menghitung jumlah kemunculan barang
	countMap := make(map[string]int)

	for _, item := range items {
		countMap[item]++
	}

	// Mengonversi pemetaan ke dalam bentuk slice
	var pairs []pair
	for name, count := range countMap {
		pairs = append(pairs, pair{name, count})
	}

	// Mengurutkan slice berdasarkan jumlah kemunculan (count) secara menurun
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count < pairs[j].count
	})

	return pairs
}

func main() {
	// golang->1 ruby->2 js->4
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) 
	for _, list := range pairs {
		fmt.Print(list.name, "->", list.count, " ")
	}
	fmt.Println()

	// C->1 D->2 B->3 A->4
	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})
	for _, list := range pairs {
		fmt.Print(list.name, "->", list.count, " ")
	}
	fmt.Println()

	// football->1 basketball->1 tenis->1
	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) 
	for _, list := range pairs {
		fmt.Print(list.name, "->", list.count, " ")
	}
	fmt.Println()
}
