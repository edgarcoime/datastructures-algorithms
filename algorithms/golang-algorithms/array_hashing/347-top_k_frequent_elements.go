package main

// 2023-04-26
// TC: O(n)
//   - Iterate through the array only once,
//     even though result is nested for the only operation
//     we are operating on is when we hit an input element
//
// SC: O(n)
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	println(count, freq)

	for _, n := range nums {
		count[n]++
	}

	for k, v := range count {
		freq[v] = append(freq[v], k)
	}

	ans := []int{}
	for i := len(freq) - 1; i > 0; i-- {
		for _, n := range freq[i] {
			ans = append(ans, n)
			if len(ans) >= k {
				return ans
			}
		}
	}

	return ans
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 3
	ans := topKFrequent(nums, k)
	println(ans)

	// for i := len(nums) - 1; i >= 0; i-- {
	// 	println(nums[i])
	// }
	for _, n := range nums {
		println(n)
	}
	println("yes")
}
