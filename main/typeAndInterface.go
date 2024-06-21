package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engineInterface interface {
	milesLeft() uint8
}

func canMakeIt(e engineInterface, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("you can make it there!")
	} else {
		fmt.Println("Need to fuel up first")
	}
}

type owner struct {
	name string
}

func TypeHandler() {
	var car gasEngine
	car.gallons = 20
	car.name = "srrin"
	fmt.Printf("the car = %v", car)
	var newCar gasEngine = gasEngine{20, 15, owner{"string"}}

	// struct single use

	var myEngine = struct {
		mgp     uint8
		gallons uint8
	}{20, 32}

	fmt.Printf("the single use struct, newCAr =%v", myEngine)
	fmt.Printf("the new car = %v", newCar)

	fmt.Printf("prototype miles left value is =%v", newCar.milesLeft())
}
