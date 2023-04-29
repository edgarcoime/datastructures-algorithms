package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

// Printing and reader
func main() {
	pl("What is your name? ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %s\n", name)
}
