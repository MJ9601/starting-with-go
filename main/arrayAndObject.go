package main

import "fmt"

func DefineArray() {
	var intArr [3]int32
	intArr[0] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])

	var _intArr [3]int32 = [3]int32{1, 2, 3}

	__intArr := [...]int32{1, 2, 3}

	fmt.Print(_intArr, __intArr)

	var initSlice []int32 = []int32{4, 5, 6}

	fmt.Printf("the length is %v with capacity %v", len(initSlice), cap(initSlice))

	initSlice = append(initSlice, 7)

	fmt.Printf("the length is %v with capacity %v", len(initSlice), cap(initSlice))

	var initSlice2 []int32 = []int32{8, 9}

	initSlice = append(initSlice, initSlice2...)

	fmt.Println(initSlice)

	var initSlice3 []int32 = make([]int32, 3, 10)

	initSlice = append(initSlice, initSlice3...)
	fmt.Println(initSlice)

	var mapObject map[string]uint8 = make(map[string]uint8)
	fmt.Println(mapObject)

	var mapObject2 = map[string]uint8{"adam": 23, "sara": 34}

	fmt.Print(mapObject2)

	// map always returns something even if key doesn't exits
	// so always use second bool parameter to check whatever it exists or not

	var age, ok = mapObject2["jason"]
	if ok {
		fmt.Printf("this is %v", age)
	} else {
		fmt.Printf("invalid name")
	}
}
