package main

import (
	"Learning-Go/pak/mypkg"
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
		Pl(mypkg.Hello(name))
	} else {
		log.Fatal(err)
	}
	
}
