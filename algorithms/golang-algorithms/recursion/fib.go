package main

import (
	"fmt"
	"time"
)

func timeExec(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

// Recursive without optimization
func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

// recursive optimized with caching
// Max n for int32 is fib(46~)
func fib2(n int) uint64 {
	cache := make(map[int]uint64)

	return fib2_helper(n, cache)
}

func fib2_helper(n int, cache map[int]uint64) uint64 {
	// Returns if 0, 1
	if n < 2 {
		return uint64(n)
	}

	if val, ok := cache[n]; ok {
		return val
	}

	result := fib2_helper(n-1, cache) + fib2_helper(n-2, cache)
	cache[n] = result

	return result
}

// Iterative
func fib3(n int) uint64 {
	if n < 2 {
		return uint64(n)
	}

	a := uint64(0)
	b := uint64(1)
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

func main() {
	fibTarget := 35

	timeExec(func() {
		println(fib(fibTarget))
	})

	timeExec(func() {
		println(fib2(fibTarget))
	})

	timeExec(func() {
		println(fib3(fibTarget))
	})
}
