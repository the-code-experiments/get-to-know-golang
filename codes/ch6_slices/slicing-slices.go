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
	employee = make([]string, 7)

	employee[0] = "Ashwin"
	employee[1] = "Kumar"
	employee[2] = "Saju"
	employee[3] = "Ajay"
	employee[4] = "Jerin"
	employee[5] = "Vinayak"
	employee[6] = "Pavan"

	/**
	 * Slices and updating slices
	 */
	employee = employee[0:]
	fmt.Println("[0:]", employee)

	employee = employee[0:7]
	fmt.Println("[0:7]", employee)

	employee = employee[0:len(employee)]
	fmt.Println("[0:len(employee)]", employee)

	employee = employee[:4]
	fmt.Println("[0:4]", employee)

	employee = employee[0:4]
	fmt.Println("[:4]", employee)

	employee = employee[4:7]
	fmt.Println("[4:7]", employee)
}