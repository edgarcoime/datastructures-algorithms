package main

// ----- CLOSURES -----
// Closures are functions that don't have to be
// associated with an identifier (Anonymous)

func useFunc(f func(int, int) int, x, y int) {
	println("Answer:", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func main() {
	intSum := func(x, y int) int { return x + y }
	println("5 + 4 =", intSum(5, 4))

	sample1 := 1
	changeVar := func() {
		sample1 += 1
	}
	changeVar()
	println("sample1 =", sample1)

	useFunc(sumVals, 5, 8)
}
