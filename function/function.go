package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"Aku", "Kece"}
	printMessage("Hai gaes", names)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
