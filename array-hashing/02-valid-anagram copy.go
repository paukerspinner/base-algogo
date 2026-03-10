// https://leetcode.com/problems/valid-anagram/description/
package main

import "fmt"

func isAnagram(s string, t string) bool {
	frequence := map[rune]int{}

	for _, c := range s {
		frequence[c]++
	}

	for _, c := range t {
		frequence[c]--
	}

	for _, f := range frequence {
		if f != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("a", "a"))
	fmt.Println(isAnagram("ax", "xa"))
	fmt.Println(isAnagram("rrrr", "rrr"))
	fmt.Println(isAnagram("axx", "xaa"))
}
