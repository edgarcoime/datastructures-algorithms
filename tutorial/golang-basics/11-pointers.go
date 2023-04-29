package main

import "fmt"

func changeVal(myPtr *int) {
	*myPtr = 12
}

func dblArrVals(arr *[4]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, val := range nums {
		sum += val
	}
	return (sum / numSize)
}

func main() {
	val := 10
	changeVal(&val) // Pass by reference
	fmt.Println("Value after func:", val)

	val2 := 10
	var val2Ptr *int = &val2
	fmt.Println("f2 address:", val2Ptr)
	fmt.Println("f2 value:", *val2Ptr)
	*val2Ptr = 11
	fmt.Println("f2 value:", *val2Ptr)

	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	fmt.Println(pArr)

	isSlice := [6]float64{11, 13, 17, 20, 22, 23}
	fmt.Printf("Average: %.3f\n", getAverage(isSlice[:3]...))
}
