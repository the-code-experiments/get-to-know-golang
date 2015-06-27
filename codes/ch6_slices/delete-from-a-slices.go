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
	employee = make([]string, 6)

	employee[0] = "Ashwin"
	employee[1] = "Kumar"
	employee[2] = "Saju"
	employee[3] = "Ajay"
	employee[4] = "Vinayak"
	employee[5] = "Jerin"

	fmt.Println("Slices: ", employee)

	// 0:2 - include 0, 1 but excludes 2
	// 4:  - include last
	// ... - expanding slices
	employee = append(employee[0:2], employee[4:]...)

	fmt.Println("After Delete: ", employee)
}