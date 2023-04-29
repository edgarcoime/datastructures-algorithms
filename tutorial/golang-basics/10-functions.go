package main

import "fmt"

func getSum(x int, y int) int {
	return x + y
}

func returnTwo(x int) (int, int) {
	return x + 1, x + 2
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You cannot divide by zero")
	}

	return x / y, nil
}

// Varadic functions
func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArrSum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(getSum(5, 4))
	fmt.Println(returnTwo(5))
	fmt.Println(getQuotient(5, 4))
	fmt.Println(getSum2(1, 2, 3, 4, 5))

	vArr := []int{1, 2, 3, 4, 5}
	fmt.Println("Array sum:", getArrSum(vArr)) // Pass by value

	val := 10
	changeVal(&val) // Pass by reference
	fmt.Println("Value after func:", val)
}
