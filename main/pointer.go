package main

import "fmt"


func SquareFuncCaller() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	var result [5]float64 = squareFunc(&thing1) // use & to pass the pointer value of variables.
	fmt.Printf("\n the result is: %v", result)
}

func squareFunc(thing2 *[5]float64) [5]float64 { // if you use [5]float64 it create new var which is different with things1 but if you use *[5]float64, it uses the pointer of things1 and change things1 value without creating new variables.
	fmt.Printf(("\n the memory location of the thing2 array is: %p"), &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]

	}
	return *thing2
}



func handlePointers (){
	var p *int32 = new (int32)  // need initial with with new, and points to memory location

	var i int32
	fmt.Printf("The value of p points to is: %v", *p) // to get the value use *
	fmt.Printf("\n the value of i is: %v", i) // getting value of normal var

	*p = 10 // assign new value to p
	p = &i // when we need the memory address of variable not its value.

}