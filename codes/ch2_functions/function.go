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
 * Define a Greeting function;
 */
func Greeting(github Profile) {
	fmt.Println("Name: ", github.name)
	fmt.Println("Username: ", github.username)
	fmt.Println("Message: ", github.message)
}

func main() {

	var github = Profile{"Ashwin Hegde", "hegdeashwin", "Welcome to Go world!"}

	/**
	 * Call the function and pass the data to the function
	 */
	Greeting(github)
}