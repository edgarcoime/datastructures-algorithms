package main

import (
	"fmt"
	"log"
	"reflect"

	mypackage "example.com/project/mypackage"
)

func main() {
	fmt.Println("Hello")
	intArr := []int{2, 3, 5, 7, 11, 13}
	strArr := mypackage.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))

	date := mypackage.Date{}
	err := date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetYear(1974)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("1st Day: %d/%d/%d\n", date.Year(), date.Month(), date.Day())
}
