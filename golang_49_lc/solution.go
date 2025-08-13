package main

import (
	"fmt"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	elems := make(map[string][]string)
	for _, val := range strs {
		valRunes := []rune(val)
		slices.Sort(valRunes)
		key := string(valRunes)
		elems[key] = append(elems[key], val)
	}

	for _, val := range elems {
		res = append(res, val)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
