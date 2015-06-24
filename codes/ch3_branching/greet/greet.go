package greet

import "fmt"

/**
 * User defined type Profile act as struct type
 */
type Profile struct {
	Name string
	Username string
	Message string
}

type Printer func(string) ()
/**
 * Define a CreateMessage function;
 * 
 * username is variadic function, can only use ... as final argument in list
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

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func GetPrefix(name string) (prefix string) {
	/**
	 * By default no break statement required.
	 * Use fallthrought to continue with the following cases.
	 * Cases can be expression.
	 */
	switch name {
		case "Ashwin Hegde": 
			prefix = "Mr. "
			fallthrough

		/**
		 * Either `Issac Newton` or `Albert Einstein`.
		 */
		case "Isaac Newton", "Albert Einstein" : 
			prefix = "Sir. "

		default: 
			prefix = "Dude. "
	}

	return
}

/**
 * Define a Greet function;
 */
func Greet(github Profile, do Printer, isFormat bool) {
	
	wel, inf := CreateMessage(github.Name, github.Username, github.Message)

	/**
	 * Commenting exact below "Println(wel) line would throw an error "wel declared and not used"
	 * In case you want to ignore the wel declaration and use info => replace wel with _ as below syntax
	 *
	 * E.g. _, info := CreateMessage(github.name, github.username, github.message)
	 */

	/**
	 * if statement with embedded-statement
	 * if [var] condition { ... } else { ... }
	 */
	if prefix := GetPrefix(github.Name); isFormat {
		do(prefix + wel)
	} else {
		do(inf)	
	}	

	// fmt.Println(_)  // Cannot use _ as value
}