package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func basicForLoops() {
	// for initialization; condition; postStatement { BODY }
	for x := 5; x >= 1; x-- {
		// x only in scope within loop
		fmt.Println(x)
	}

	// While like syntax
	fX := 0
	for fX < 5 {
		fmt.Println(fX)
		fX++
	}

	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		fmt.Println(num)
	}
}

func guessingGame() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1

	gameContinue := true
	for gameContinue {
		fmt.Println("Guess a number between 0 and 50:")
		fmt.Println("Random number is:", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}

		if iGuess == randNum {
			fmt.Printf("Yes! %d is the right number\n", iGuess)
			break
		} else if iGuess > randNum {
			fmt.Println("Pick a lower value")
		} else if iGuess < randNum {
			fmt.Println("Pick a higher value")
		} else {
			fmt.Println("%d, seems invalid, try again.", guess)
		}

		fmt.Println()
	}
}

func main() {
	basicForLoops()
	guessingGame()
}
