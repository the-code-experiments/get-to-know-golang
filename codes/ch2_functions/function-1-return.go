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
 * func <function-name>([param1, param2, ...]) <return-type> { ... }
 */
func CreateMessage(name, username, message string) string {
	return message + ", " + name + " (" + username + ")"
}

/**
 * Define a Greeting function;
 */
func Greeting(github Profile) {
	fmt.Println(CreateMessage(github.name, github.username, github.message))
}

func main() {

	var github = Profile{"Ashwin Hegde", "hegdeashwin", "Welcome to Go world!"}

	/**
	 * Call the function and pass the data to the function
	 */
	Greeting(github)
}