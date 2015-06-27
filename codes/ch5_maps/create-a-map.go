package main

/**
 * Note: The code will not run and throw an error: prefixMap declared and not used
 * As we have declare map and not used in and this is error in go-lang
 */
func main() {

	/**
	 * Declare a map of map type with string.
	 * var <variable-name> map[<key-type>] <value-type>
	 */
	var prefixMap map[string] string

	/**
	 * make will initialize the map; if we don't use make 
	 * declaration will set to null;
	 */
	prefixMap = make(map[string] string)

}