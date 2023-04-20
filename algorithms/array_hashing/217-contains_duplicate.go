package main

// Time complexity: O(n)
//   - Iterate through array only once
//
// Space complexity: O(n)
//   - Have to generate a map for array
func containsDuplicate(nums []int) bool {
	nums_map := map[int]int{}
	for _, n := range nums {
		if _, ok := nums_map[n]; !ok {
			nums_map[n] = 1
		} else {
			return true
		}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	println(containsDuplicate(nums))
}
