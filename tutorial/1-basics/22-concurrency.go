package main

import (
	"errors"
	"fmt"
	"time"
)

// ----- CONCURRENCY -----
// Concurrency allows us to have multiple
// blocks of code share execution time by
// pausing their execution. We can also
// run blocks of codes in parallel at the same
// time. (Concurrent tasks in Go are called
// goroutines)

// To execute multiple functions in new
// goroutines add the word go in front of
// the function calls (Those functions can't
// have return values)

// We can't control when functions execute
// so we may get different results

func printUpToN(n int) error {
	if n < 1 {
		return errors.New("n has to be greater or equal to 1")
	}
	for i := 1; i <= n; i++ {
		fmt.Printf("Fun (n=%d): %d\n", n, i)
	}
	return nil
}

func basicConcurrency() {
	go printUpToN(15)
	go printUpToN(10)

	// We have to pause the main function because
	// if main ends so will the goroutines
	time.Sleep(2 * time.Second)
}

// You can have goroutines communicate by
// using channels. The sending goroutine
// also makes sure the receiving goroutine
// receives the value before it attempts
// to use it

// Create a channel : Only carries values of
// 1 type
func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

func channels() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	go nums1(channel1)
	go nums2(channel2)
	println(<-channel1)
	println(<-channel1)
	println(<-channel1)
	println(<-channel2)
	println(<-channel2)
	println(<-channel2)
}

func main() {
	basicConcurrency()
	channels()
}
