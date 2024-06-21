package main

import (
	"fmt"
	"time"
)

func TimeCal() {
	var n int = 100000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v \n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v \n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()

	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

func loopHandlerFunc() {
	var mapObject = map[string]uint8{"adam": 9, "sarah": 5}

	for i, v := range mapObject {
		fmt.Printf("the index %v, the value %v, \n", i, v)
	}

	// while loop

	var i int = 0

	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i += 1
	}

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	// operations
	// i-- decrements by 1,
	// i++ increments by 1,
	// i+=10 increments by 10
	// i-=10  decrements by 10
	// i*= 10 multiply by 10
	// i/=10 divide by 10

}