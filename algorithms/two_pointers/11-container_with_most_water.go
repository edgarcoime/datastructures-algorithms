package main

// 2023-04-27
// TC: O(n) - Iterating through height arr using two pointers
// SC: O(n) - Pointer logic no extra space needed
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	max := 0
	l := 0
	r := len(height) - 1
	for l < r {
		// Calculate area
		a := min(height[l], height[r]) * (r - l)
		if a > max {
			max = a
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return max
}

func main() {
}
