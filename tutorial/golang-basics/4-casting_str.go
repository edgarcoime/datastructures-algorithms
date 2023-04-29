package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	cV3 := "5000"
	cV4, err := strconv.Atoi(cV3)
	fmt.Println(cV4, err, reflect.TypeOf(cV4))

	cv5 := 5000
	cv6 := strconv.Itoa(cv5)
	fmt.Println(cv6)

	cv7 := "3.14"
	if cv8, err := strconv.ParseFloat(cv7, 65); err == nil {
		fmt.Println(cv8)
	}

	cv9 := fmt.Sprintf("%f", 3.14)
	fmt.Println(cv9)
}
