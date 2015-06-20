package main

import "fmt"

/**
 * Define constants
 * Float, String and iota (incremental)
 */
const (
	PI = 3.14
	LANGUAGE = "Go"

	A = iota
	B = iota
	C = iota
)

/**
 * Define constants iota (incremental)
 */
const (
	P = iota
	Q = iota
	R = iota
)

/**
 * Define constants iota (incremental)
 * Here, Y and Z does has type, so Go will assume Y and Z are of type iota from X
 */
const (
	X = iota
	Y
	Z
)

func main() {

	/**
	 * OUTPUT: 3.14
	 */
	fmt.Println(PI)

	/**
	 * OUTPUT: Go
	 */
	fmt.Println(LANGUAGE)

	/**
	 * OUTPUT: 2 3 4
	 */
	fmt.Println(A, B, C)

	/**
	 * OUTPUT: 0 1 2
	 */
	fmt.Println(P, Q, R)

	/**
	 * OUTPUT: 0 1 2
	 */
	fmt.Println(X, Y, Z)
}