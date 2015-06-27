package main

import "fmt";

func main() {

	/**
	 * `for` loop
	 *
	 * for [assignment]; [conditions]; [increment/decrement] { ... }
	 */
	for i := 0; i < 10; i++ {

		/**
		 * loop will break; once i is greater than 4
		 */
		if i > 4 {
			break;
		}

		fmt.Println("i =", i)
	}

}