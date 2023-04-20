package main

import "fmt"

// ----- DEFINED TYPES -----
// We used a defined type previously with structs
// You can use them also to enhance the quality
// of other data types
// We'll create them for different measurements

type (
	Tsp float64
	TBs float64
	ML  float64
)

func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func TBtoML(tsp TBs) ML {
	return ML(tsp * 14.79)
}

func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}

func (tbs TBs) ToMLs() ML {
	return ML(tbs * 14.79)
}

func main() {
	tsp1 := Tsp(3)
	fmt.Printf("%.2f tsp = %.2f mL\n", tsp1, tsp1.ToMLs())
}
