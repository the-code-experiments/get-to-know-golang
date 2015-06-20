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
 * func <function-name>([param1, param2, ...]) (<naming-return-val1><return-type1>, <naming-return-val1><return-type2>) { ... }
 */
func CreateMessage(name, username, message string) (welcome string, info string) {

	/**
	 * <naming-return-val1> = string
	 * <naming-return-val2> = string
	 */
	welcome = "\n" + message + " " + name
	info = "You are authorize to access the system: " + username + "\n"

	return
}

/**
 * Define a Greeting function;
 */
func Greeting(github Profile) {
	
	wel, inf := CreateMessage(github.name, github.username, github.message)

	/**
	 * Commenting exact below "Println(wel) line would throw an error "wel declared and not used"
	 * In case you want to ignore the wel declaration and use info => replace wel with _ as below syntax
	 *
	 * E.g. _, info := CreateMessage(github.name, github.username, github.message)
	 */
	fmt.Println(wel)
	fmt.Println(inf)

	// fmt.Println(_)  // Cannot use _ as value
}

func main() {

	var github = Profile{"Ashwin Hegde", "hegdeashwin", "Welcome to Go world!"}

	/**
	 * Call the function and pass the data to the function
	 */
	Greeting(github)
}