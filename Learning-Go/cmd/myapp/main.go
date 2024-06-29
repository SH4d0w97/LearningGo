package main

import (
	"Learning-Go/pak/myexercises"
	// "Learning-Go/pak/mypkg"
	// "bufio"
	"fmt"
	// "log"
	// "os"
	// "strings"
)

// Pl is an alias for fmt.Println for brevity
var pl = fmt.Println

func main() {

	/*

	NOTE: ADD MENU OPTION AFTER LEARNING SWITCH-CASE

	pl("--------------- Hello World ---------------")

	pl("Enter your name : ")             // Prompt the user to enter their name
	reader := bufio.NewReader(os.Stdin)  // Create a new reader for standard input
	name, err := reader.ReadString('\n') // Read the user's input until a newline character

	if err == nil {
		// Trim any leading/trailing whitespace and newline characters strings.TrimSpace()
		pl(strings.TrimSpace(mypkg.Hello(name))) // If there's no error, greet the user using the mypkg.Hello function
	} else {
		log.Fatal(err) // Otherwise, log the error and exit the program
	}

	pl("--------------- Values ---------------")
	mypkg.Values()

	pl("--------------- Constants ---------------")
	mypkg.Constants()

	*/

	pl("--------------- EXERCISE 1 ---------------")
	// pl("Result : ",	myexercises.StrConcat()) V1
	pl("RESULT : ", myexercises.StrConcatV2())

}
