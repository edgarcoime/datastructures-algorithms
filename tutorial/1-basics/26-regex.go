package main

import "regexp"

// TODO: Not working properly
// ----- REGULAR EXPRESSIONS -----
// You can use regular expressions to test
// if a string matches a pattern

func main() {
	// Search for ape followed by not a space
	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?)", reStr)
	println(match)

	// You can compile them
	// Find multiple words ending with at
	reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("([crmfp]at)")

	// Did you find any matches?
	println("MatchString :", r.MatchString(reStr2))

	// Return first match
	println("FindString :", r.FindString(reStr2))

	// Starting and ending index for 1st match
	println("Index :", r.FindStringIndex(reStr2))

	// Return all matches
	println("All String :", r.FindAllString(reStr2, -1))

	// Get 1st 2 matches
	println("All String :", r.FindAllString(reStr2, 2))

	// Get indexes for all matches
	println("All Submatch Index :", r.FindAllStringSubmatchIndex(reStr2, -1))

	// Replace all matches with Dog
	println(r.ReplaceAllString(reStr2, "Dog"))
}
