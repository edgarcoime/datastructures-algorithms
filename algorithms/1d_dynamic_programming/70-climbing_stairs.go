package main

import "fmt"

func climbStairsIt(n int) int {
	one, two := 1, 1

	for i := 0; i < n; i++ {
		tmp := one
		one = one + two
		two = tmp
	}

	return one
}

func climbStairs(n int) int {
	memo := make(map[int]int)
	return _helper(n, memo)
}

func _helper(n int, memo map[int]int) int {
	if n <= 2 {
		return n
	}

	// Check if already calculated
	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = _helper(n-1, memo) + _helper(n-2, memo)
	return memo[n]
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	for _, n := range nums {
		fmt.Println(climbStairs(n))
	}
}
