package main

import (
	"fmt"
	"os"
	"strconv"
)

func findMax() {
	fmt.Println(os.Args)
	args := os.Args[1:]
	iArgs := []int{}
	for _, arg := range args {
		val, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}
	max := 0
	for _, arg := range iArgs {
		if arg > max {
			max = arg
		}
	}
	fmt.Println("Max value:", max)
}

func main() {
	findMax()
}
