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

	return prefixMap[name]
}

func main() {

	fmt.Println("What is Ashwin's role? He is " + GetPrefix("Ashwin"))

}