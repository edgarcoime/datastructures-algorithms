package main

// TC: O(n) - Iterating through arr only once
// SC: O(1) - Just pointers no extra allocation
func maxProfit(prices []int) int {
	l, r := 0, 1 // left = buy | right = sell
	maxProf := 0

	for r < len(prices) {
		// Profitable?
		if prices[r] > prices[l] {
			prof := prices[r] - prices[l]
			if prof > maxProf {
				maxProf = prof
			}
		} else {
			l = r
		}
		r++
	}

	return maxProf
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	println(maxProfit(prices))
}
