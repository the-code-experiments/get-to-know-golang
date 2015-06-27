package main

import "fmt";

func GetPrefix(name string) (prefix string) {

	var prefixMap map[string] string
	prefixMap = make(map[string] string)

	/**
	 * Insert key and value into the map.
	 *
	 * <map-variable>[<key>] = <value>
	 */
	prefixMap["Ashwin"] = "Sr. Fullstack Engineer"
	prefixMap["Kumar"] = "Sr. Engineering Manager"
	prefixMap["Saju"] = "Sr. Solution Architect"
	prefixMap["Ajay"] = "Sr. solution Architect"

	return prefixMap[name]
}

func main() {

	fmt.Println("What is Ashwin's role? He is " + GetPrefix("Ashwin"))

}