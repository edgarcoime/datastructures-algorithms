package main

import "fmt"

func sliceBasics() {
	// var name []datatype
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "Of"
	sl1[2] = "The"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"
	fmt.Println("Slice Size:", len(sl1))
	for i := 0; i < len(sl1); i++ {
		fmt.Println(sl1[i])
	}
	for _, x := range sl1 {
		fmt.Println(x)
	}
}

func sliceArr() {
	sArr := [5]int{1, 2, 3, 4, 5}
	sl := sArr[0:2] // exclusive of 2
	fmt.Println(sl)
	fmt.Println("First 3:", sArr[:3])
	fmt.Println("Last 3:", sArr[2:])

	// Slices will change if array is changed
	sArr[0] = 10
	fmt.Println("sl:", sl)

	// Changing the slice will also change the arr
	sl[0] = 1000
	fmt.Println("arr:", sArr)

	// Adding values in slice will overwrite arr
	sl = append(sl, 12)
	fmt.Println("sl", sl)
	fmt.Println("arr", sArr)

	// Creating a slice will just print nils
	sl2 := make([]string, 6)
	fmt.Println("sl2:", sl2)
	fmt.Println("sl2[0]:", sl2[0])
}

func main() {
	// sliceBasics()
	sliceArr()
}
