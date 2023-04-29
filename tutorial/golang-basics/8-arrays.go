package main

import "fmt"

func arrBasics() {
	// Array Init
	var arr1 [5]int
	arr1[0] = 1

	// Array iteration
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Index 0:", arr2[0])
	fmt.Println("Arr Length:", len(arr2))
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}

	// Matrix
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arr3[i][j])
		}
	}
}

func stringArr() {
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune Array: %d\n", v)
	}

	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	fmt.Println("I'm a string:", bStr)
}

func main() {
	// arrBasics()
	stringArr()
}
