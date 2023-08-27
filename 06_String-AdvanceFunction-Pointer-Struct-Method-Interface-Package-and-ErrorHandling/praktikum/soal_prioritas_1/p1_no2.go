package main

import (
	"fmt"
)

// Mendefinisikan struct Studen berisi array string name dan array integer score
type Student struct {
	name  []string
	score []int
}

// Mendefinisikana method average pada struct penerima s Student
func (s Student) Average() float64 {
	totalScore := 0
	// Menghitung total skor dari banyaknya score
	for _, score := range s.score {
		totalScore += score
	}
	// Menghitung rata-rata skor dari total jumlah score dan banyaknya score
	return float64(totalScore) / float64(len(s.score))
}

// Mendefinisikana method min pada struct penerima s Student
func (s Student) Min() (min int, name string) {
	// Tempat menyimpan nilai minimum dan nama siswanya
	min = s.score[0]
	name = s.name[0]
	// Mencari nilai minimum dan nama siswa dengan skor terendah
	for i := 1; i < len(s.score); i++ {
		// Menyimpan nilai minimum dan siswa ke variabel min dan name
		if s.score[i] < min {
			min = s.score[i]
			name = s.name[i]
		}
	}
	return min, name
}

// Mendefinisikana method min pada struct penerima s Student
func (s Student) Max() (max int, name string) {
	// Tempat menyimpan nilai maksimum dan nama siswanya
	max = s.score[0]
	name = s.name[0]
	// Mencari nilai maksimum dan nama siswa dengan skor tertinggi
	for i := 1; i < len(s.score); i++ {
		// Menyimpan nilai minimum dan siswa ke variabel min dan name
		if s.score[i] > max {
			max = s.score[i]
			name = s.name[i]
		}
	}
	return max, name
}

func main() {
	var a = Student{}

	for i := 1; i <= 5; i++ {
		var name string
		fmt.Print("Input ", i, " Student's Name ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input ", i, " Student's Score ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\nAverage Score Students is", a.Average())

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score of Students : ",nameMin,"(",scoreMin,")")

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score of Students : ",nameMax,"(",scoreMax,")")
}
