package main

// Works for Unicode
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sr := []rune(s)
	tr := []rune(t)

	s_map := map[rune]int{}
	t_map := map[rune]int{}

	// Populate s_map
	// Populate t_map
	for i := 0; i < len(sr); i++ {
		s_map[sr[i]] += 1
		t_map[tr[i]] += 1
	}

	for k := range s_map {
		if s_map[k] != t_map[k] {
			return false
		}
	}

	return true
}

func main() {
	// t := "anagram"
	// s := "nagaram"
	t := "rat"
	s := "car"
	println(isAnagram(s, t))
}
