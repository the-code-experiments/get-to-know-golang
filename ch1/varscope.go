package main

import "fmt"

var name string = "Ashwin"
func main() {
	fmt.Println("My name is ", name)
	fullname()
}

func fullname() {
	fmt.Println("My fullname is ", name + "Hegde") 
}
