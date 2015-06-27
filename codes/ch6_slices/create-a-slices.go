package main

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
	// employee[3] = "Ajay" // This will throw an error `out of range`
}