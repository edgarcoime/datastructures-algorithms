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

	return nil
}

// Hash Map store complement
// TC: O(n)
// SC: O(n)
func twoSum(nums []int, target int) []int {
	// Create a map
	// key (num) = value (index)
	num_map := map[int]int{}

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if v, ok := num_map[diff]; ok {
			return []int{i, v}
		} else {
			num_map[nums[i]] = i
		}
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	println(twoSum(nums, target))
}
