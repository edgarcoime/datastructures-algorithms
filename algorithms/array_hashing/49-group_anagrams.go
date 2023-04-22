package main

import (
	"sort"
	"strings"
)

// TC: O(nlogn) - sorting algorithm uses quicksort BEST:nlogn WORST: n^2
// SC: O(logn) - call stack space complexity BEST: logn WORST: O(n)
func sortString(s string) string {
	strArr := strings.Split(s, "")
	sort.Strings(strArr)
	return strings.Join(strArr, "")
}

// TC: O(m * nlogn) - length of strings arr * the sorting algorithm needed
// SC: O(m) - Where m is the length of strings array that is stored in the map
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)

	for _, str := range strs {
		key := sortString(str)
		if _, ok := strMap[key]; !ok {
			strMap[key] = []string{}
		}
		strMap[key] = append(strMap[key], str)
	}

	out := make([][]string, len(strMap))
	i := 0
	for _, anagrams := range strMap {
		out[i] = anagrams
		i++
	}

	return out
}

// TODO: Create a better hashing function that is at at most O(n)
func generateID(s string) byte {
	// id := 0
	var id byte
	id = 0
	for i := 0; i < len(s); i++ {
		id += s[i] - 'a'
	}
	return id
}

func groupAnagrams2(strs []string) [][]string {
	strMap := make(map[byte][]string)

	for _, str := range strs {
		id := generateID(str)
		if _, ok := strMap[id]; !ok {
			strMap[id] = []string{}
		}
		strMap[id] = append(strMap[id], str)
	}

	out := make([][]string, len(strMap))
	i := 0
	for _, anagrams := range strMap {
		out[i] = anagrams
		i++
	}

	return out
}

func main() {
	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	strs := []string{"ac", "c"}
	println(groupAnagrams(strs))
	println(generateID("car"))
	println(generateID("rac"))
	println(generateID("rat"))
}
