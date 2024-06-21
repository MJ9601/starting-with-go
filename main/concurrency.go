package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// without concurrency

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func DbCallerWithoutConcurrency() {
	t0 := time.Now()
	for i := range dbData {
		dbCall(i)
		
	}
	fmt.Printf("\n total execuation time: %v", time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Microsecond)
	fmt.Println("the result from db is:", dbData[i])
}



var wg = sync.WaitGroup{}
var mutex = sync.Mutex{}
var dbData_ = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func DbCallerWithConcurrency() {
	t0 := time.Now()
	for i := range dbData_ {
		wg.Add(1)
		go dbCallWith(i)

	}
	wg.Wait()
	fmt.Printf("\n total excucation time: %v", time.Since(t0))
	fmt.Printf("\n total results are: %v", results)
}

func dbCallWith(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Microsecond)
	fmt.Println("the result from db is:", dbData_[i])
mutex.Lock() // it's very important where you put Lock and Unlock

results = append(results, dbData_[i])
mutex.Unlock()
wg.Done()
}