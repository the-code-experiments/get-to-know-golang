package main

import "fmt"

/**
 * User defined type Profile act as struct type
 */
type Profile struct {
	name string
	username string
	message string
}

/**
 * Define a CreateMessage function;
 * Either specify string for each or just make string at last parameter and first 2 will be asssumed as string
 * Also specify the return type which in case here is string
 *
 * func <function-name>([param1, param2, ...]) (<return-type1>, <return-type2>) { ... }
 */
func CreateMessage(name, username, message string) (string, string) {

	/**
	 * return <first-return>, <second-return>, ...
	 */
	return "\n" + message + " " + name, "You are authorize to access the system: " + username + "\n"
}

/**
 * Define a Greeting function;
 */
func Greeting(github Profile) {
	

	welcome, info := CreateMessage(github.name, github.username, github.message)

	/**
	 * Commenting exact below "Println(welcome) line would throw an error "welcome declared and not used"
	 * In case you want to ignore the welcome declaration and use info => replace welcome with _ as below syntax
	 *
	 * E.g. _, info := CreateMessage(github.name, github.username, github.message)
	 */
	fmt.Println(welcome)
	fmt.Println(info)

	// fmt.Println(_)  // Cannot use _ as value
}

func main() {

	var github = Profile{"Ashwin Hegde", "hegdeashwin", "Welcome to Go world!"}

	/**
	 * Call the function and pass the data to the function
	 */
	Greeting(github)
}