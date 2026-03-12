// https://leetcode.com/problems/group-anagrams/
package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}

	for _, str := range strs {
		f := frequency(str)

		if _, found := m[f]; found {
			m[f] = append(m[f], str)
		} else {
			m[f] = []string{str}
		}
	}

	results := [][]string{}

	for _, groups := range m {
		results = append(results, groups)
	}

	return results
}

func frequency(str string) [26]int {
	f := [26]int{}

	for _, c := range str {
		idx := int(c - 'a')
		f[idx]++
	}

	return f
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
	fmt.Println(groupAnagrams([]string{"a", "b", "cd", "de"}))
	fmt.Println(groupAnagrams([]string{"abc", "abc", "abc", "bac", "cba", "cab"}))
	fmt.Println(groupAnagrams([]string{"z", "az", "za", "a"}))
}
