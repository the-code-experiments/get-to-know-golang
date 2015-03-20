package main

import "fmt"

func main() {
	i := 1

	switch i {
	case 0:
		fmt.Println("Zero")

	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")

	default:
		fmt.Println("Unknown number")
	}
}
