package mypkg // Package declaration

import "fmt" // Importing the fmt package for formatted I/O

// Values function demonstrates basic printing and arithmetic operations
func Values() {
	fmt.Println("Hello," + "Go") // Concatenating and printing strings
	fmt.Println("1 + 2 =", 1+2)   // Performing addition and printing the result
	fmt.Println("1.20 + 3.80 = ", 1.20+3.80) // Performing floating-point addition and printing

	// Prints the result of "true AND false"
	fmt.Println("True and False = ", true && false) 
	// Prints the result of "true OR false"
	fmt.Println("True || False = ", true || false) 
	// Prints the negation of "true"
	fmt.Println("!True = ", !true)

}
