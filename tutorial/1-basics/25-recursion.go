package main

// ----- RECURSION -----
// Recursion occurs when a function calls itself
// There must be a condition that ends this
// Finding a factorial is commonly used

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	println("Factorial 4 =", factorial(4))
	// 1st : result = 4 * factorial(3) = 4 * 6 = 24
	// 2nd : result = 3 * factorial(2) = 3 * 2 = 6
	// 3rd : result = 2 * factorial(1) = 2 * 1 = 2
}
