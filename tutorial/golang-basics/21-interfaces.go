package main

import "fmt"

// ----- INTERFACES -----
// Interfaces allow you to create contracts
// that say if anything inherits it that
// they will implement defined methods

// If we had animals and wanted to define that
// they all perform certain actions, but in their
// specific way we could use an interface

// With Go you don't have to say a type uses
// an interface. When your type implements
// the required methods it is automatic

type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	fmt.Println("Cat Attacks its Prey")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	fmt.Println("Cat says hisss")
}

func (c Cat) HappySound() {
	fmt.Println("Cat says Purr")
}

func main() {
	var kitty1 Animal
	kitty1 = Cat("kitty")
	kitty1.AngrySound()

	var kitty2 Cat = kitty1.(Cat)
	kitty2.Attack()
	println("Cats name:", kitty2.Name())
}
