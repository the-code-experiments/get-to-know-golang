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
		 * loop will continue; once i is greater than 5
		 */
		if i > 5 {
			continue;
		}

		/**
		 * this will get skipped once i is greate then 5;
		 */
		fmt.Println("i =", i)
	}

}