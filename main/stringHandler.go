package main

import (
	"fmt"
	"strings"
)



func StringHandler(){
	var myString = []rune("resume")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i,v := range myString{
	fmt.Println(i, v)	
	}

	fmt.Printf("\n the length of 'myString' is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\n myRune = %v", myRune)

	var strSlice = []string{"s", "u", "b", "s"}
	var strBuilder strings.Builder;

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])	
	}

	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)


}