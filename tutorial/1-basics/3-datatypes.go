package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(25))
	fmt.Println(reflect.TypeOf(3.13))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello"))
	fmt.Println(reflect.TypeOf('😂'))
}
