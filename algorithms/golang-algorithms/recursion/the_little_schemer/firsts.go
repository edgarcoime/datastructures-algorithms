package main

func firsts(l [][]string) []string {
	if len(l) <= 0 {
		return []string{}
	}
	return nil
}

func firsts2(l [][]string) []string {
	acc := []string{}
	return _helper(l, acc)
}

func _helper(l [][]string, acc []string) []string {
	return nil
}

func main() {
}
