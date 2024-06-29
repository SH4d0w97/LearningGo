package mypkg

import (
	"fmt"
)

func Constants(){
	// Declare a constant string variable "message"
	const message string = "Hello, World" 
	// Print the value of "message"
	fmt.Println(message)
	// This will throw an error as you are trying to redeclare a variable that was declared as a constant.
	// message := "Hi, World" 
	// Declare a constant variable pi
	const pi = 3.14 
	// Print the value of pi
	fmt.Println("PI = ",pi)
}
