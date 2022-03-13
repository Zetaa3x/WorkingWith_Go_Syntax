package main

import "log"

func main() {
	//initialize an int array of 5 elements
	//values are defaulted to 0
	var numbers [5]int

	//assign a value to the 3rd index of the array
	numbers[2] = 3

	//declare a variable with the same value as the 3rd
	//index in the numbers array.
	i := numbers[2]

	//print out test message, the value of i and use the
	// '==' operator to compare variable i with an integer
	// the comparison will return a true or false value
	log.Println("test, the value of i is", i, "is i equal to 3?", i == 3)
	log.Printf("Array item = %v", numbers[2])
}
