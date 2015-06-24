package main

import "./greet";

func main() {

	var github = greet.Profile{"Ashwin Hegde", "hegdeashwin", "Welcome to Go world!"}

	/**
	 * Call the function and pass the data to the function
	 */
	greet.Greet(github, greet.CreatePrintFunction("?"), true)
}