package main

import "fmt"

func main() {

	/**
	 * Blank [] will declare a slices.
	 * If you provide size like [5] will declare an array.
	 */
	var employee []string

	/**
	 * Initialize a slices with type, length and capacity.
	 * length is required field.
	 * 
	 * make([]string, length, [capacity])
	 */
	employee = make([]string, 3)

	employee[0] = "Ashwin"
	employee[1] = "Kumar"
	employee[2] = "Saju"

	fmt.Println("Slices: ", employee)

	/**
	 * Appending into slices.
	 */
	employee = append(employee, "Ajay")
	employee = append(employee, "Vinayak")

	fmt.Println("After appending: ", employee)

}