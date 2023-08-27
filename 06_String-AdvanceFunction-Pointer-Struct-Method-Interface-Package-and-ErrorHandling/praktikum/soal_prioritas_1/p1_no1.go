package main

import (
	"fmt"
)

// Definisikan struct Car
type Car struct {
	carType string
	fuelin  float64
}

// Method untuk menghitung perkiraan jarak
func (c Car) Estimate() float64 {
	// Menghitung jarak 
	distance := c.fuelin * 1.5 
	return distance
}

func main() {
	// Membuat instance Car
	car := Car{
		carType: "SUV",
		fuelin:  10.5,
	}

	// Mencetak informasi tipe mobil dan perkiraan jarak
	fmt.Printf("car type: %s , est. distance: %.2f\n", car.carType, car.Estimate())
}
