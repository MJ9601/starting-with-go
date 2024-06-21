package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func DefVariables() {

	// defining variables
	var _int int = 324

	fmt.Println(_int)

	var _int8 int8 = 8
	fmt.Println(_int8)

	var _float32 float32 = 44.43
	fmt.Println(_float32)

	var _float64 float64 = 43.53
	fmt.Println(_float64)

	var _sum = _float32 + float32(_float64)

	fmt.Println(_sum)

	var text = "text"
	fmt.Println(text)

	_text := "_text"

	fmt.Println(_text)
	fmt.Println(len(_text)) // length of ascii char

	num1, numb2 := 14, 21

	fmt.Println(num1, numb2)

	fmt.Println(utf8.RuneCountInString(_text))

	fmt.Println("hello world")

	boolean := true

	fmt.Println(boolean)

	const myConst string = "const value"

	fmt.Println(myConst)
}

func PrintHandler(result int, remainder int, err error) {
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("result = %v", result)
	default:
		fmt.Printf("result = %v, remainder = %v", result, remainder)
	}

	//  // conditional switch case
	// switch remainder{
	// case 0:
	// 	fmt.Print("The division was exact!")
	// case 1,2:
	// 	fmt.Print("The Division was close")
	// default:
	// 	fmt.Print("The remainder = %v", remainder)
	// }

}

func handleDiv(numb1 int, numb2 int) (int, int, error) {
	var err error

	if numb2 == 0 {
		err = errors.New("cannot Divide by 0")
		return 0, 0, err
	}

	var result int = numb1 / numb2
	var remainder int = numb1 % numb2

	return result, remainder, err
}