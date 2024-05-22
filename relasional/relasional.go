package main

import "fmt"

func main() {
	var value = (((2+6)%3)*4 - 2) / 2
	var isEqual = (value == 2)
	fmt.Println("nilai %d (%t) \n", value, isEqual)
}
