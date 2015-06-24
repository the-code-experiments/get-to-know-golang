package main

import "./greeting";

func main() {

	var github = greeting.Profile{"Ashwin Hegde", "hegdeashwin", "Welcome to Go world!"}

	/**
	 * Call the function and pass the data to the function
	 */
	greeting.Greeting(github, greeting.CreatePrintFunction("?"), true)
}