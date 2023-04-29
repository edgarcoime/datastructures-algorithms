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

// Normal recursion
func fact(n uint64) uint64 {
	if n <= 0 {
		return 1
	}

	return n * fact(n-1)
}

// tail Recursion
func tailFact(n uint64) uint64 {
	acc := uint64(1)
	return _helper(n, acc)
}

func _helper(n uint64, acc uint64) uint64 {
	if n <= 0 {
		return acc
	}

	return _helper(n-1, acc*n)
}

func main() {
	target := uint64(100)
	timeExec(func() {
		println(fact(target))
	})
	timeExec(func() {
		println(tailFact(target))
	})
}
