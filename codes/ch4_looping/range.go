package main

import "fmt";

type Profile struct {
	name string
	role string
}

func main() {

	slice := []Profile {
		{"Ashwin", "Sr. Fullstack Engineer"},
		{"Kumar", "Sr. Engineering Manager"},
		{"Saju", "Sr. Solution Architect"},
		{"Ajay", "Sr. Solution Architect"}, // comma is needed here
	}

	/**
	 * _ is index; we are not using it; thus replacing with _
	 */
	for _, employee := range slice {
		fmt.Println("Profile: " + employee.name + " (" + employee.role + ")" )
	}

}