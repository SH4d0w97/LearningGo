package myexercises

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StrConcat() string{

	/*
	NOTE: When using fmt.Scanf("%s\n", &inputStr), 
	it reads only the first word or sequence of characters until the first whitespace (space, tab, newline).
	*/

	var outputString string
	var inputStr string
	
	fmt.Printf("Enter the First string : ")
	// fmt.Scanln(&inputStr) V1.0
	// fmt.Scanf("%s",&inputStr) V2.0
	fmt.Scanf("%s\n",&inputStr) //V3.0
	outputString += inputStr + " "


	fmt.Printf("Enter the Second string : ")
	// fmt.Scanf("%s",&inputStr)
	fmt.Scanf("%s\n",&inputStr) //V3.0
	outputString += inputStr + " "

	fmt.Printf("Enter the Third string : ")
	// fmt.Scanf("%s\n",&inputStr)
	fmt.Scanf("%s\n",&inputStr) //V3.0
	outputString += inputStr + " "

	return outputString
}

func StrConcatV2() string{

	var outputString string
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the first string : ")
	str1, err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("ERROR!! : ",err)
		return ""
	}
	outputString += strings.TrimSpace(str1)
	outputString += " "

	fmt.Printf("Enter the second string : ")
	str2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ERROR!! : ",err)
		return ""
	}
	outputString += strings.TrimSpace(str2)
	outputString += " "

	fmt.Printf("Enter the third string : ")
	str3, err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("ERROR!! : ",err)
	}
	outputString += strings.TrimSpace(str3)

	return outputString
}