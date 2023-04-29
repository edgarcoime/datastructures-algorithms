package main

import "fmt"

// ----- MAPS ----
// Maps are collections of key/value pairs
// Keys can be any data type that can be compared
// using == (They can be a different type than
// the value)
// var myMap map [keyType]valueType

func main() {
	var heroes map[string]string
	heroes = make(map[string]string)
	villains := make(map[string]string)

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"
	villains["Lex Luther"] = "Lex Luther"

	superPets := map[int]string{
		1: "Krypto",
		2: "Bat Hound",
	}
	fmt.Printf("Batman is %v\n", heroes["Batman"])
	fmt.Println("Chip:", superPets[3])
	_, ok := superPets[3]
	fmt.Println("Is there a 3rd pet?:", ok)
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}
	delete(heroes, "The Flash")
}
