package main

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	buckets := make([][]int, len(nums)+1)

	// Fill freq
	for _, n := range nums {
		freq[n]++
	}

	// Fill bucket
	for k, v := range freq {
		buckets[v] = append(buckets[v], k)
	}

	// Generate ans(res)
	res := make([]int, 0)
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, n := range buckets[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}

func main() {
}
