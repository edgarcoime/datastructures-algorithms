package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	res := nums[0]
	curMin, curMax := 1, 1

	for i := 0; i < len(nums); i++ {
		tmp := curMax * nums[i] // Use old curMax
		curMax = max(nums[i], max(nums[i]*curMax, nums[i]*curMin))
		curMin = min(nums[i], min(nums[i]*tmp, nums[i]*curMin))
		res = max(res, curMax)
	}

	return res
}

func main() {
	nums := []int{2, 3, 0, -2, 4, -1}
	// nums := []int{-2, 0, -1}
	fmt.Println(maxProduct(nums))
}
