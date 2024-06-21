package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func sumSliceCaller() {
	var intSlice = []int{1, 2, 3}

	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice(float32Slice))
}

// T also can be type of any
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func callIsEmpty() {
	var intSlice = []int{}
	var float32Slice = []float32{1, 2, 3}
	fmt.Println(isEmpty(float32Slice))
	fmt.Println(isEmpty(intSlice))

}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name    string
	Price   float32
	Account int
}

func loadJSONcaller() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("\n%v", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")

	fmt.Printf("\n%v", purchases)
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	var data, _ = ioutil.ReadFile(filePath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}

// generic with struct

type gasEngineStruct struct {
	gallons float32
	mpg     float32
}

type electricEngineStruct struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngineStruct | electricEngineStruct] struct {
	carMake  string
	carModel string
	engine   T
}

func callCarStructFunc() {
	var gasCar = car[gasEngineStruct]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngineStruct{
			gallons: 21.3,
			mpg:     22,
		},
	}
	var elecCar = car[electricEngineStruct]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: electricEngineStruct{
			kwh:   30,
			mpkwh: 43,
		},
	}
}