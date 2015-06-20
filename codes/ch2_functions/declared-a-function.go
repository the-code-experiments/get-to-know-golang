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

type Printer func(string) ()
/**
 * Define a CreateMessage function;
 * 
 * username is variadic function, can only use ... as final argument in list
 */
func CreateMessage(name string, username string, message ...string) (welcome string, info string) {

	/**
	 * <naming-return-val1> = string
	 * <naming-return-val2> = string
	 */
	welcome = "\n" + message[0] + " " + name
	info = "You are authorize to access the system: " + username + "\n"

	fmt.Println(message[1])
	fmt.Println(message[2])
	fmt.Println("Number of parameters: ", len(message))

	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

/**
 * Define a Greeting function;
 */
func Greeting(github Profile, do Printer) {
	
	wel, inf := CreateMessage(github.name, github.username, github.message, "Go is concurrent", "Go is awesome")

	/**
	 * Commenting exact below "Println(wel) line would throw an error "wel declared and not used"
	 * In case you want to ignore the wel declaration and use info => replace wel with _ as below syntax
	 *
	 * E.g. _, info := CreateMessage(github.name, github.username, github.message)
	 */
	do(wel)
	do(inf)

	// fmt.Println(_)  // Cannot use _ as value
}

func main() {

	var github = Profile{"Ashwin Hegde", "hegdeashwin", "Welcome to Go world!"}

	/**
	 * Call the function and pass the data to the function
	 */
	Greeting(github, Print) 
	Greeting(github, PrintLine)
}