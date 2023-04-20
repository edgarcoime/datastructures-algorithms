package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// Strings
	sv1 := "a word"
	replacer := strings.NewReplacer("A", "Another")
	sv2 := replacer.Replace(sv1)
	fmt.Println(sv2)
	fmt.Println("Length: ", len(sv2))
	fmt.Println("Contains Another", strings.Contains(sv2, "Another"))
	fmt.Println("o index: ", strings.Index(sv2, "o"))
	fmt.Println("Replace: ", strings.Replace(sv2, "o", "0", -1))
	sv3 := "\nSome Words\n" // \t \" \\
	sv3 = strings.TrimSpace(sv3)
	fmt.Println("Split: ", strings.Split("a-b-c-d", "-"))
	fmt.Println("Lower: ", strings.ToLower(sv2))
	fmt.Println("Upper: ", strings.ToUpper(sv2))
	fmt.Println("Prefix: ", strings.HasPrefix("tacocat", "taco"))
	fmt.Println("Suffix: ", strings.HasPrefix("tacocat", "cat"))

	// Runes
	rStr := "abcdefg"
	fmt.Println("Rune Count: ", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}
