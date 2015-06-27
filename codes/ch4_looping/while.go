package main

import "fmt";

func main() {

	/**
	 * There is no while loop in Go-lang;
	 * but we can use `for` loop to behave like `while` loop.
	 */
	i := 0;
	for i < 10 {
		fmt.Println("i =", i)

		i++
	}

}