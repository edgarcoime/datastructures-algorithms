package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TC: O(n) - Iterating through array once
// SC: O(1) - Pointers used only
func trap(height []int) int {
	// endpoints can't hold water
	l, r := 0, len(height)-1
	maxL, maxR := height[l], height[r]
	res := 0

	for l < r {
		// Evaluate, then update max
		if maxL < maxR {
			res += max(maxL-height[l], 0)

			l++
			maxL = max(maxL, height[l])
		} else {
			res += max(maxR-height[r], 0)

			r--
			maxR = max(maxR, height[r])
		}
	}

	return res
}

// TC: O(4n) - Iterating array only 4 times
// SC: O(3n) - Needs extra space to calculate
// func trap(height []int) int {
// 	maxLeft := make([]int, len(height))
// 	maxRight := make([]int, len(height))
// 	minHeight := make([]int, len(height))
//
// 	// Keep track of maxleft
// 	for i := 0; i < len(height)-1; i++ {
// 		if i == 0 {
// 			continue
// 		}
//
// 		maxLeft[i] = max(maxLeft[i-1], height[i-1])
// 	}
//
// 	// Keep track of maxRight
// 	for i := len(height) - 1; i >= 0; i-- {
// 		if i == len(height)-1 {
// 			continue
// 		}
//
// 		maxRight[i] = max(maxRight[i+1], height[i+1])
// 	}
//
// 	// find min height
// 	for i := 0; i < len(height)-1; i++ {
// 		minHeight[i] = min(maxLeft[i], maxRight[i])
// 	}
//
// 	res := 0
// 	// Sum up water
// 	for i := 0; i < len(height)-1; i++ {
// 		res += max(minHeight[i]-height[i], 0)
// 	}
//
// 	return res
// }

func main() {
}
