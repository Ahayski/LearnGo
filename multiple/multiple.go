package main

import (
	"fmt"
	"math"
)

func main() {
	area, circum := calculate(28)
	fmt.Println("Luas lingkaran", area)
	fmt.Println("Keliling lingkaran", circum)
}

func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d

	return area, circumference
}
