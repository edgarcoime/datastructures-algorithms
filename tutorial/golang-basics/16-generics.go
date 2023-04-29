package main

// ----- GENERICS -----
// We can specify the data type to be used at a
// later time with generics
// It is mainly used when we want to create
// functions that can work with
// multiple data types

// https://pkg.go.dev/golang.org/x/exp/constraints
type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	println("5 + 4 =", getSumGen(5, 4))
	println("5.6 + 4.7 =", getSumGen(5.6, 4.7))
}
