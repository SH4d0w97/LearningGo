package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//NOTE Alais
var Pl = fmt.Println

func main() {
	Pl("Enter your name : ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		fmt.Println("Hello,", name)
	} else {
		log.Fatal(err)
	}
	
}
