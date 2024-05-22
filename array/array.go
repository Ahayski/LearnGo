package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "banana", "mango", "grape"}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)
	fmt.Println(fruits[0])
}
