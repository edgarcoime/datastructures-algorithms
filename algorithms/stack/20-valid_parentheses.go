package main

// TC: O(n) - Iterating through string once
// SC: O(n) - Creating a stack with potentially the same size or half size of array
func isValid(s string) bool {
	pairs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	// LIFO structure
	stack := []byte{}

	// [({})]
	for _, char := range []byte(s) {
		pair, ok := pairs[char]

		// If not found append
		if !ok {
			stack = append(stack, char)
			continue
		}

		// if found and length 0 means invalid closing
		if len(stack) == 0 {
			return false
		}

		// Peek stack and check if matching with compliment pair
		if stack[len(stack)-1] != pair {
			return false
		}

		// Pop last
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func main() {
	s := "({[]})"
	println(isValid(s))
}
