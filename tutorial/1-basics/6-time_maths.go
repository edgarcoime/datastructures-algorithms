package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	mInt := 1
	mInt = mInt + 1
	mInt += 1
	mInt++

	// Seed is fro 1/1/1970
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	fmt.Println("Random:", randNum)

	// There are many math functions
	fmt.Println("Abs(-10) =", math.Abs(-10))
	fmt.Println("Pow(4, 2) =", math.Pow(4, 2))
	fmt.Println("Sqrt(16) =", math.Sqrt(16))
	fmt.Println("Cbrt(8) =", math.Cbrt(8))
	fmt.Println("Ceil(4.4) =", math.Ceil(4.4))
	fmt.Println("Floor(4.4) =", math.Floor(4.4))
	fmt.Println("Round(4.4) =", math.Round(4.4))
	fmt.Println("Log2(8) =", math.Log2(8))
	fmt.Println("Log10(100) =", math.Log10(100))
	// Get the log of e to the power of 2
	fmt.Println("Log(7.389) =", math.Log(math.Exp(2)))
	fmt.Println("Max(5,4) =", math.Max(5, 4))
	fmt.Println("Min(5,4) =", math.Min(5, 4))

	// Convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	// Convert 1.5708 radians to degrees
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%f radians = %f degrees\n", r90, d90)
	fmt.Println("Sin(90) =", math.Sin(r90))

	// There are also functions for Cos, Tan, Acos, Asin
	// Atan, Asinh, Acosh, Atanh, Atan2, Cosh, Sinh, Sincos
	// Htpot

	// ----- FORMATTED PRINT -----
	// Go has its own version of C's printf
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value

	fmt.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A',
		3.14, true, 1, 1)

	// Float formatting
	fmt.Printf("%9f\n", 3.14)      // Width 9
	fmt.Printf("%.2f\n", 3.141592) // Decimal precision 2
	fmt.Printf("%9.f\n", 3.141592) // Width 9 no precision

	// Sprintf returns a formatted string instead of printing
	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	fmt.Println(sp1)
}
