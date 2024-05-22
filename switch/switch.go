package main

import "fmt"

func main() {
	var point = 10
	switch point {
	case 8:
		fmt.Println("Perfect")
	case 7:
		{
			fmt.Println("Awesome")
		}
	default: // ga wajib, kalo ga pake default conditionalnya tetap jalan juga
		{
			fmt.Println("Not Bad")
		}
	}
}
