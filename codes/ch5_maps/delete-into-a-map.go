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
		/**
		 * Old way to perform delete operation on map.
		 * Will no more work for new version of Go compiler
		 */
		// prefixMap["Saju"] = "", false

		/**
		 * New way to perform delete operation on map.
		 */
		delete(prefixMap, "Saju")
	}
	
	return prefixMap[name]
}

func main() {

	fmt.Println("What is Saju's role? He is " + GetPrefix("Saju", false))

	fmt.Println("What is Saju's role? He is " + GetPrefix("Saju", true))

}