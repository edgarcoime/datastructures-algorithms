package main

// Brute Force
// TC: O(n^2)
// SC: O(1)
func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for k := i + 1; k < len(nums); k++ {
			if nums[i]+nums[k] == target {
				return []int{i, k}
			}
		}
	}

	return []int{}
}

func twoSum(nums []int, target int) []int {
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	println(twoSum2(nums, target))
}
