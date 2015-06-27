package main

import "fmt";

func GetPrefix(name string, mustDel bool) (prefix string) {

	/**
	 * Short hand way to declare and initialize map
	 */
	prefixMap := map[string] string {
		"Ashwin": "Sr. Fullstack Engineer",
		"Kumar": "Sr. Engineering Manager",
		"Saju": "Sr. Solution Architect",
		"Ajay": "Sr. Solution Architect", // comma is needed here
	}

	if mustDel {
		delete(prefixMap, "Saju")
	}

	/**
	 * Check if the value exist into the map or not.
	 */
	if _, exists := prefixMap[name]; exists {
		return prefixMap[name]
	} else {
		return "dude"
	}

}

func main() {

	fmt.Println("What is Saju's role? He is " + GetPrefix("Saju", false))

	fmt.Println("What is Saju's role? He is " + GetPrefix("Saju", true))

}