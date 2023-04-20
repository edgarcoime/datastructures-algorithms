package main

func containsDuplicate(nums []int) bool {
	memo := make(map[int]int)

	// Create a map thats stores nums inside
	for i := 0; i < len(nums); i++ {
		_, ok := memo[nums[i]]
		if ok == true {
			return true
		} else {
			memo[nums[i]] = nums[i]
		}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	println(containsDuplicate(nums))
}
