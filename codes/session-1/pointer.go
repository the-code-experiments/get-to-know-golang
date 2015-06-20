package main 

import "fmt"

func main() {

	/**
	 * Created message variable and assigned string to it.
	 */
	message := "Hello World"

	/**
	 * Created greeting string pointer to hold the memory address of message variable
	 */
	var greeting *string = &message

	/**
	 * Print message
	 */
	fmt.Println(message)

	/**
	 * Print greeting memory address
	 */
	fmt.Println(greeting)

	/**
	 * Print value of greeting
	 */

	fmt.Println(*greeting)
}