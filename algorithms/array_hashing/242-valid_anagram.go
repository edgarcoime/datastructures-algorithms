package main

// Time Complexity: O(s+t)
// Memory Complexity: O(s+t)
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

// Assuming all lower
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	comparator := make([]int, 26) // 26 chars assyming all lower case
	// evaluate chars
	for i := 0; i < len(s); i++ {
		comparator[s[i]-'a']++
		comparator[t[i]-'a']--
	}

	// Check if 0s
	for _, n := range comparator {
		if n != 0 {
			return false
		}
	}

	return true
}

// Follow up question
// How to solve with O(1) space?
//   - Sort the string first
//   - use two pointers to compare the two strings

func main() {
	// t := "anagram"
	// s := "nagaram"
	// t := "rat"
	// s := "car"
	t := "rac"
	s := "car"
	// println(isAnagram(s, t))
	println(isAnagram2(s, t))
	// println('z' - 'a')
}
