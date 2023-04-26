package main

import (
	"sort"
	"strings"
)

// 2023-04-26
// TC: O(m*n*26)
// 		- m: length of strs array
// 		- n: average length of string in arr
// 		- 26: Array size for key
// SC: O(n*26) - Create a cache and create an integer array size 26 to store as key

func generateAnagramID(s string) [26]int {
	var key [26]int
	for i := 0; i < len(s); i++ {
		key[s[i]-'a']++
	}
	return key
}

func groupAnagrams2(strs []string) [][]string {
	// Cache anagrams together
	cache := make(map[[26]int][]string)
	for _, s := range strs {
		k := generateAnagramID(s)
		cache[k] = append(cache[k], s)
	}

	// Collect anagrams
	res := [][]string{}
	for _, anagrams := range cache {
		res = append(res, anagrams)
	}

	return res
}

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

func main() {
	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	strs := []string{"ac", "c"}
	println(groupAnagrams(strs))
}
