package mypkg

import (
	"fmt"
)

func Constants(){
	const message string = "Hello, World"
	fmt.Println(message)
	message := "Hi, World"
	const pi = 3.14
	fmt.Println("PI = ",pi)

}