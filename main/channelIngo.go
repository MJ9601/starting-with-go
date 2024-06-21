package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_TOFU_PRICE float32 = 20
var MAX_CHICKEN_PRICE float32 = 10

func sampleChannel() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "contco.com", "wholefoods.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)

}

func checkTofuPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofu_price = rand.Float32() * 20
		if tofu_price < MAX_TOFU_PRICE {
			c <- website
			break
		}
	}
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}

	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\n Text sent: Found deal on chicken at %v", website)
	case website := <-tofuChannel:
		fmt.Printf("\n Email sent: Found deal on tofu at %v", website)
	}
}

func callingChannelProcess() {
	var c = make(chan int, 5)
	go processChannel(c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second + 1)

	}
}

func processChannel(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting process")
}