package main

import "fmt"

/**
 * User defined type customString act as string type
 */
type customString string

/**
 * User defined type userInfo act as struct type
 */
type userInfo struct {
	name string
	greeting string
}

func main() {

	/**
	 * Use customString as type
	 */
	var message customString = "Hello World!"
	fmt.Println(message)

	/**
	 * 1 way: Assign values to userInfo struct
	 */
	var defaults = userInfo{};
	defaults.name = "Unknown"
	defaults.greeting = "Welcome"

	fmt.Println(defaults.name)
	fmt.Println(defaults.greeting)

	/**
	 * 2 way: Assign values to userInfo struct; Go by order
	 */
	var facebookProfile = userInfo{"Ashwin", "Hello!"}

	fmt.Println(facebookProfile.name)
	fmt.Println(facebookProfile.greeting)

	/**
	 * 3 way: Assign values to userInfo struct; Go by member name
	 */
	var linkedinProfile = userInfo{greeting: "Hello Go!", name: "Ashwin Hegde"}

	fmt.Println(linkedinProfile.name)
	fmt.Println(linkedinProfile.greeting)
}