package main

import (
	"Learning-Go/pak/mypkg"
	"bufio"
	"fmt"
	"log"
	"os"
)

// Pl is an alias for fmt.Println for brevity
var Pl = fmt.Println

func main() {
	Pl("Enter your name : ")             // Prompt the user to enter their name
	reader := bufio.NewReader(os.Stdin)  // Create a new reader for standard input
	name, err := reader.ReadString('\n') // Read the user's input until a newline character

	if err == nil {
		Pl(mypkg.Hello(name)) // If there's no error, greet the user using the mypkg.Hello function
	} else {
		log.Fatal(err) // Otherwise, log the error and exit the program
	}
}
