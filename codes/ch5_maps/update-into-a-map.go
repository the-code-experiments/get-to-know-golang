package main

import "fmt";

func GetPrefix(name string) (prefix string) {

	/**
	 * Short hand way to declare and initialize map
	 */
	prefixMap := map[string] string {
		"Ashwin": "Sr. Fullstack Engineer",
		"Kumar": "Sr. Engineering Manager",
		"Saju": "Sr. Solution Architect",
		"Ajay": "Sr. Solution Architect", // comma is needed here
	}

	/**
	 * Perform update operation on map.
	 */
	prefixMap["Kumar"] = "Delivery Manager"
	prefixMap["Ajay"] = "Sr. Project Manager"

	return prefixMap[name]
}

func main() {

	fmt.Println("What is Kumar's role? He is " + GetPrefix("Kumar"))
	fmt.Println("What is Ajay's role? He is " + GetPrefix("Ajay"))

}