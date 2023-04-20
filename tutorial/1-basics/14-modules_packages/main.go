package main

import (
	"fmt"
	"reflect"

	mypackage "example.com/project/mypackage"
)

func main() {
	fmt.Println("Hello")
	intArr := []int{2, 3, 5, 7, 11, 13}
	strArr := mypackage.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))
}
