package main

import "fmt"

func main() {
	var lastName uint8 = 89
	var firstName string = "Febryan"

	var numberA int = 4
	var numberB *int = &numberA
	fmt.Print("Halo sayangku ", firstName, lastName, numberB, "!\n")

}
