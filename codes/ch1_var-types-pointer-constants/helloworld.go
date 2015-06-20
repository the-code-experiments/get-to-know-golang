/**
 * Declare the file as main package.
 */
package main

/**
 * import format;
 */
import "fmt"

/**
 * Main function acts as a single entry point.
 */
func main() {

	/**
	 * Print a message
	 */
	fmt.Println("Hello World");

	/**
	 * Declare a variable with string type and assign a string value to it;
	 * Type is optional.
	 * 
	 * Note: Declaring variable and do not use is an error in Go
	 */
	var message string
	message = "Welcome to Go World!"

	fmt.Println(message);

	/**
	 * Declare multiple variables with integer type and assign a integer values to them;
	 * Type is optional.
	 */
	var a, b, c int = 1, 2, 3

	fmt.Println(a, b, c);

	/**
	 * Declare a variable and assign a string value to it with short way;
	 */
	name := "Ashwin Hegde"
	
	fmt.Println(name)


}